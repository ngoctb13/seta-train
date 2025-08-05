package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/seta-train/rest-service/handler/models"
	useCaseModel "github.com/ngoctb13/seta-train/rest-service/internal/domains/models"
)

func (h *Handler) CreateNotesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		folderID := c.Param("folderId")
		if folderID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "folder ID is required"})
		}

		var req models.CreateNoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			h.logger.Error("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.CreateNotes(c, &useCaseModel.CreateNotesInput{
			FolderID: folderID,
			Notes:    models.ToNoteUseCaseModel(req.Notes),
			UserID:   userID.(string),
		})

		if err != nil {
			h.logger.Error("CreateNotesUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("Created notes successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Notes created successfully"})
	}
}

func (h *Handler) ViewNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			h.logger.Error("note ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		note, err := h.note.ViewNote(c, &useCaseModel.ViewNoteInput{
			UserID: userID.(string),
			NoteID: noteID,
		})

		if err != nil {
			h.logger.Error("ViewNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("ViewNoteHandler successfully")
		c.JSON(http.StatusOK, note)
	}
}

func (h *Handler) UpdateNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			h.logger.Error("note ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		var req models.UpdateNoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			h.logger.Error("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.UpdateNote(c, &useCaseModel.UpdateNoteInput{
			UserID: userID.(string),
			NoteID: noteID,
			Note: useCaseModel.Note{
				Title: req.Title,
				Body:  req.Body,
			},
		})

		if err != nil {
			h.logger.Error("UpdateNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("UpdateNoteHandler successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
	}
}

func (h *Handler) DeleteNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			h.logger.Error("note ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.DeleteNote(c, &useCaseModel.DeleteNoteInput{
			UserID: userID.(string),
			NoteID: noteID,
		})

		if err != nil {
			h.logger.Error("DeleteNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		h.logger.Info("DeleteNoteHandler successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
	}
}

func (h *Handler) ShareNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			h.logger.Error("note ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		var req models.ShareNoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			h.logger.Error("binding json error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.ShareNote(c, &useCaseModel.ShareNoteInput{
			NoteID:        noteID,
			CurUserID:     userID.(string),
			SharedUserIDs: req.SharedUserIDs,
			AccessType:    req.AccessType,
		})

		if err != nil {
			h.logger.Error("ShareNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("ShareNoteHandler successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Note shared successfully"})
	}
}

func (h *Handler) RevokeSharingNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			h.logger.Error("note ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		sharedUserID := c.Param("userId")
		if sharedUserID == "" {
			h.logger.Error("shared user ID is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "shared user ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.RevokeSharingNote(c, &useCaseModel.RevokeSharingNoteInput{
			CurUserID:    userID.(string),
			NoteID:       noteID,
			SharedUserID: sharedUserID,
		})

		if err != nil {
			h.logger.Error("RevokeNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		h.logger.Info("RevokeSharingNoteHandler successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Note sharing revoked successfully"})
	}
}
