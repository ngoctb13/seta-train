package handler

import (
	"log"
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
			log.Printf("binding json error: %v", err)
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
			log.Printf("CreateNotesUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Notes created successfully"})
	}
}

func (h *Handler) ViewNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		note, err := h.note.ViewNote(c, &useCaseModel.ViewNoteInput{
			UserID: userID.(string),
			NoteID: noteID,
		})

		if err != nil {
			log.Printf("ViewNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, note)
	}
}

func (h *Handler) UpdateNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		var req models.UpdateNoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("binding json error: %v", err)
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
			log.Printf("UpdateNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
	}
}

func (h *Handler) DeleteNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		noteID := c.Param("noteId")
		if noteID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "note ID is required"})
		}

		userID, _ := c.Get(userIDKey)

		err := h.note.DeleteNote(c, &useCaseModel.DeleteNoteInput{
			UserID: userID.(string),
			NoteID: noteID,
		})

		if err != nil {
			log.Printf("DeleteNoteUsecase fail with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
	}
}
