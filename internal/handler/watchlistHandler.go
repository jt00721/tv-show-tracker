package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/tv-show-tracker/internal/domain"
	"github.com/jt00721/tv-show-tracker/internal/usecase"
)

type WatchlistHandler struct {
	Usecase usecase.WatchlistUsecase
}

func NewWatchlistHandler(u usecase.WatchlistUsecase) *WatchlistHandler {
	return &WatchlistHandler{Usecase: u}
}

func (handler *WatchlistHandler) CreateWatchlistApi(c *gin.Context) {
	var watchlist domain.Watchlist
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		log.Printf("Error binding json request body to create watchlist: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input to create watchlist"})
		return
	}

	err := handler.Usecase.CreateWatchlist(&watchlist)
	if err != nil {
		if err.Error() == "ShowID or UserID cannot be 0" {
			log.Println("Error: Cannot create watchlist without ShowID and UserID")
			c.JSON(http.StatusBadRequest, gin.H{"error": "ShowID or UserID cannot be 0"})
			return
		}

		log.Printf("Error creating watchlist: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create watchlist. Please try again later."})
		return
	}

	log.Println("Successfully created watchlist")
	c.JSON(http.StatusCreated, watchlist)
}

func (handler *WatchlistHandler) GetWatchlistApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting userID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	watchlists, err := handler.Usecase.GetWatchlist(uint(id))
	if err != nil {
		log.Printf("Error retrieving watchlist for user with ID(%d): %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "watchlist not found",
		})
		return
	}

	log.Printf("Successfully fetched watchlist for userwith ID(%d)", id)
	c.JSON(http.StatusOK, watchlists)
}

func (handler *WatchlistHandler) UpdateWatchlistApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting watchlist ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}

	var watchlist domain.Watchlist
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		log.Printf("Error binding json request body to update watchlist: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input to update watchlist",
		})
		return
	}

	watchlist.ID = uint(id)
	err = handler.Usecase.UpdateWatchlist(&watchlist)
	if err != nil {
		log.Printf("Error updating watchlist with ID(%d): %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update watchlist. Please try again later."})
		return
	}

	log.Printf("Successfully updated watchlist with ID: %v", watchlist.ID)
	c.JSON(http.StatusOK, watchlist)
}

func (handler *WatchlistHandler) DeleteWatchlistApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting watchlist ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}

	err = handler.Usecase.DeleteWatchlist(uint(id))
	if err != nil {
		log.Printf("Error deleting watchlist with ID(%d): %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete watchlist. Please try again later."})
		return
	}

	log.Printf("Successfully deleted watchlist with ID(%d)", id)
	c.JSON(http.StatusOK, gin.H{"message": "Watchlist deleted"})
}

func (h *WatchlistHandler) MarkWatchedApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting watchlist ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}

	var req struct {
		LastWatchedEpisode string `json:"last_watched_episode"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.LastWatchedEpisode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid last_watched_episode"})
		return
	}

	if err := h.Usecase.MarkWatched(uint(id), req.LastWatchedEpisode); err != nil {
		log.Printf("Error marking watchlist as watched: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark as watched"})
		return
	}

	log.Printf("Successfully marked watchlist with ID(%d) as watched", id)
	c.JSON(http.StatusOK, gin.H{"message": "Watchlist updated"})
}
