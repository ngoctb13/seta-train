package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/ngoctb13/seta-train/internal/domains/user/usecases"

type Resolver struct {
	UserUsecase *usecases.User
}
