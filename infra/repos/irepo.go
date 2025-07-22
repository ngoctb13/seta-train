package repos

import "github.com/ngoctb13/seta-train/internal/domains/user/repos"

type IRepo interface {
	Users() repos.IUserRepo
}
