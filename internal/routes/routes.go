package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jt00721/tv-show-tracker/internal/handler"
)

func SetupRoutes(r *gin.Engine, showHandler *handler.ShowHandler, watchlistHandler *handler.WatchlistHandler) {
	r.GET("/search", showHandler.SearchShowsApi)

	r.POST("/watchlist", watchlistHandler.CreateWatchlistApi)
	r.GET("/watchlist/:id", watchlistHandler.GetWatchlistApi)
	r.PUT("/watchlist/:id", watchlistHandler.UpdateWatchlistApi)
	r.DELETE("/watchlist/:id", watchlistHandler.DeleteWatchlistApi)
	r.PATCH("/watchlist/:id/mark_watched", watchlistHandler.MarkWatchedApi)
}
