package handler

import (
	"github.com/gin-gonic/gin"
	folder_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/folder/usecases"
	team_usecases "github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
)

type Handler struct {
	team   *team_usecases.Team
	folder *folder_usecases.Folder
	note   *folder_usecases.Note
}

func NewHandler(team *team_usecases.Team, folder *folder_usecases.Folder, note *folder_usecases.Note) *Handler {
	return &Handler{
		team:   team,
		folder: folder,
		note:   note,
	}
}

func (h *Handler) ConfigAuthRouteAPI(router *gin.RouterGroup) {
	router.GET("/hello", h.HelloHandler())

	// team routes
	router.POST("/teams", h.CreateTeamHandler())
	router.POST("/teams/:teamId/members", h.AddTeamMembersHandler())
	router.POST("/teams/:teamId/managers", h.AddTeamManagersHandler())
	router.DELETE("/teams/:teamId/members/:memberId", h.RemoveTeamMemberHandler())
	router.DELETE("/teams/:teamId/managers/:managerId", h.RemoveTeamManagerHandler())

	// folder routes
	router.POST("/folders", h.CreateFolderHandler())
	router.POST("/folders/:folderId", h.GetFolderDetailsHandler())
	router.PUT("/folders/:folderId", h.UpdateFolderHandler())
	router.DELETE("/folders/:folderId", h.DeleteFolderHandler())

	// note routes
	router.POST("/folders/:folderId/notes", h.CreateNotesHandler())
	router.GET("/notes/:noteId", h.ViewNoteHandler())
	router.PUT("/notes/:noteId", h.UpdateNoteHandler())
	router.DELETE("/notes/:noteId", h.DeleteNoteHandler())
}
