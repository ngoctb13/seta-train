package repos

import "github.com/ngoctb13/seta-train/rest-service/internal/domains/user/repos"

type IRepo interface {
	Users() repos.IUserRepo
}
