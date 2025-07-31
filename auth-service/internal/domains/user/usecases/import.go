package usecases

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"

	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/repos"
	"github.com/ngoctb13/seta-train/shared-modules/model"
	"github.com/ngoctb13/seta-train/shared-modules/utils"
)

type ImportUsecase struct {
	userRepo repos.IUserRepo
}

type ImportSummary struct {
	Total     int           `json:"total"`
	Succeeded int           `json:"succeeded"`
	Failed    int           `json:"failed"`
	Errors    []ImportError `json:"errors"`
}

type ImportError struct {
	Line  int    `json:"line"`
	Error string `json:"error"`
}

type CSVUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

const (
	WorkerCount = 10
	BatchSize   = 50
)

func NewImportUsecase(userRepo repos.IUserRepo) *ImportUsecase {
	return &ImportUsecase{
		userRepo: userRepo,
	}
}

func (i *ImportUsecase) ImportUsersFromCSV(ctx context.Context, fileBytes []byte) (*ImportSummary, error) {
	// Parse CSV
	users, errors := i.parseCSV(fileBytes)
	if len(errors) > 0 {
		return &ImportSummary{
			Total:     0,
			Succeeded: 0,
			Failed:    len(errors),
			Errors:    errors,
		}, nil
	}

	if len(users) == 0 {
		return &ImportSummary{
			Total:     0,
			Succeeded: 0,
			Failed:    0,
			Errors:    []ImportError{},
		}, nil
	}

	// Process users with worker pool
	return i.processUsersWithWorkerPool(ctx, users)
}

func (i *ImportUsecase) parseCSV(fileBytes []byte) ([]CSVUser, []ImportError) {
	reader := csv.NewReader(strings.NewReader(string(fileBytes)))
	reader.FieldsPerRecord = 4 // username, email, password, role

	var users []CSVUser
	var errors []ImportError
	lineNumber := 0

	// Skip header
	_, err := reader.Read()
	if err != nil {
		errors = append(errors, ImportError{
			Line:  1,
			Error: "Failed to read CSV header: " + err.Error(),
		})
		return users, errors
	}
	lineNumber++

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lineNumber++

		if err != nil {
			errors = append(errors, ImportError{
				Line:  lineNumber,
				Error: "Failed to read CSV line: " + err.Error(),
			})
			continue
		}

		if len(record) != 4 {
			errors = append(errors, ImportError{
				Line:  lineNumber,
				Error: "Invalid number of fields, expected 4",
			})
			continue
		}

		user := CSVUser{
			Username: strings.TrimSpace(record[0]),
			Email:    strings.TrimSpace(record[1]),
			Password: strings.TrimSpace(record[2]),
			Role:     strings.TrimSpace(record[3]),
		}

		// Validate user data
		if validationError := i.validateUser(user, lineNumber); validationError != nil {
			errors = append(errors, *validationError)
			continue
		}

		users = append(users, user)
	}

	return users, errors
}

func (i *ImportUsecase) validateUser(user CSVUser, lineNumber int) *ImportError {
	// Validate username
	if user.Username == "" {
		return &ImportError{
			Line:  lineNumber,
			Error: ErrUsernameRequired.Error(),
		}
	}

	// Validate email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return &ImportError{
			Line:  lineNumber,
			Error: ErrEmailInvalid.Error(),
		}
	}

	// Validate role
	if user.Role != ManagerRole && user.Role != MemberRole {
		return &ImportError{
			Line:  lineNumber,
			Error: ErrInvalidRole.Error(),
		}
	}

	return nil
}

func (i *ImportUsecase) processUsersWithWorkerPool(ctx context.Context, users []CSVUser) (*ImportSummary, error) {
	// Create channels
	userChan := make(chan CSVUser, 100)
	resultChan := make(chan ImportResult, 100)
	errorChan := make(chan ImportError, 100)

	// Start workers
	var wg sync.WaitGroup
	for w := 0; w < WorkerCount; w++ {
		wg.Add(1)
		go i.worker(ctx, userChan, resultChan, errorChan, &wg)
	}

	// Send users to workers
	go func() {
		defer close(userChan)
		for _, user := range users {
			userChan <- user
		}
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Process results
	summary := &ImportSummary{
		Total:     len(users),
		Succeeded: 0,
		Failed:    0,
		Errors:    []ImportError{},
	}

	// Collect successful results
	for range resultChan {
		summary.Succeeded++
	}

	// Collect errors
	for err := range errorChan {
		summary.Errors = append(summary.Errors, err)
		summary.Failed++
	}

	return summary, nil
}

type ImportResult struct {
	Success bool
	Error   string
}

func (i *ImportUsecase) worker(ctx context.Context, userChan <-chan CSVUser, resultChan chan<- ImportResult, errorChan chan<- ImportError, wg *sync.WaitGroup) {
	defer wg.Done()

	for user := range userChan {
		// Convert CSVUser to model.User
		modelUser := &model.User{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		}

		// Hash password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			errorChan <- ImportError{
				Line:  0, // We don't have line number in worker
				Error: fmt.Sprintf("Failed to hash password for user %s: %v", user.Username, err),
			}
			continue
		}
		modelUser.PasswordHash = hashedPassword

		// Check if user already exists
		existingUser, err := i.userRepo.GetUserByEmail(ctx, user.Email)
		if err == nil && existingUser != nil {
			errorChan <- ImportError{
				Line:  0,
				Error: fmt.Sprintf("Email already exists: %s", user.Email),
			}
			continue
		}

		// Create user with retry logic
		var createErr error
		for retry := 0; retry < 3; retry++ {
			createErr = i.userRepo.CreateUser(ctx, modelUser)
			if createErr == nil {
				break
			}
		}

		if createErr != nil {
			errorChan <- ImportError{
				Line:  0,
				Error: fmt.Sprintf("Failed to create user %s after 3 retries: %v", user.Username, createErr),
			}
			continue
		}

		resultChan <- ImportResult{Success: true}
	}
}
