package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createArtistRequest struct {
  ArtistName string `json:"artist_name" binding:"required"`
}

func (server *Server) createArtist(ctx *gin.Context) {
	var req createArtistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.ArtistName

	artist, err := server.store.CreateArtist(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, artist)
}

type getArtistRequest struct {
  ArtistName string `uri:"name" binding:"require"`
}

func (server *Server) getArtist(ctx *gin.Context) {
	var req getArtistRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.ArtistName

	artist, err := server.store.GetArtist(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, artist)
}

type listArtistsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (server *Server) listArtists(ctx *gin.Context) {
	var req listArtistsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.PageID 

	artists, err := server.store.ListArtists(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, artists)
}
