package repos

import (
	teamRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
)

type IRepo interface {
	Teams() teamRepo.ITeamRepo
}
