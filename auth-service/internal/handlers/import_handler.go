package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ngoctb13/seta-train/auth-service/internal/domains/user/usecases"
	"github.com/ngoctb13/seta-train/shared-modules/logger"
)

type ImportHandler struct {
	importUsecase *usecases.ImportUsecase
	logger        *logger.Logger
}

type ImportResponse struct {
	Success bool          `json:"success"`
	Summary ImportSummary `json:"summary"`
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

func NewImportHandler(importUsecase *usecases.ImportUsecase, logger *logger.Logger) *ImportHandler {
	return &ImportHandler{
		importUsecase: importUsecase,
		logger:        logger,
	}
}

func (h *ImportHandler) ImportUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (max 10MB)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Error("Failed to parse multipart form: %v", err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		h.logger.Error("Failed to get uploaded file: %v", err)
		http.Error(w, "No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	if header.Header.Get("Content-Type") != "text/csv" &&
		header.Filename[len(header.Filename)-4:] != ".csv" {
		h.logger.Error("Invalid file type: %s", header.Header.Get("Content-Type"))
		http.Error(w, "Only CSV files are allowed", http.StatusBadRequest)
		return
	}

	// Read file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("Failed to read file: %v", err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Process import
	summary, err := h.importUsecase.ImportUsersFromCSV(r.Context(), fileBytes)
	if err != nil {
		h.logger.Error("Import failed: %v", err)
		http.Error(w, "Import failed", http.StatusInternalServerError)
		return
	}

	// Return response
	response := ImportResponse{
		Success: true,
		Summary: ImportSummary{
			Total:     summary.Total,
			Succeeded: summary.Succeeded,
			Failed:    summary.Failed,
			Errors:    make([]ImportError, len(summary.Errors)),
		},
	}

	for i, err := range summary.Errors {
		response.Summary.Errors[i] = ImportError{
			Line:  err.Line,
			Error: err.Error,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
