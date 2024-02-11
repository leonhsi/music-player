package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/leonhsi/music-player/db/sqlc"
)

// Server serves HTTP requests fro our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

  router.POST("/songs", server.createSong)
  router.GET("/songs/:name", server.getSong)
  router.GET("/songs/", server.listSongs)

  router.POST("/artists", server.createArtist)
  router.GET("/artists/:name", server.getArtist)
  router.GET("/artists/", server.listArtists)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
