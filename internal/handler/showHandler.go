package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/tv-show-tracker/internal/usecase"
)

type ShowHandler struct {
	Usecase usecase.ShowUsecase
}

func NewShowHandler(u usecase.ShowUsecase) *ShowHandler {
	return &ShowHandler{Usecase: u}
}

func (h *ShowHandler) SearchShowsApi(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	shows, err := h.Usecase.SearchShows(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shows)
}
