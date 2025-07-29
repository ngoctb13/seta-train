package repos

import (
	folderRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/repos"
	teamRepo "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
)

type IRepo interface {
	Teams() teamRepo.ITeamRepo
	Folders() folderRepo.IFolderRepo
}
