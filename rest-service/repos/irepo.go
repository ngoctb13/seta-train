package repos

import (
	teamRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	userRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/user/repos"
)

type IRepo interface {
	Users() userRepo.IUserRepo
	Teams() teamRepo.ITeamRepo
}
