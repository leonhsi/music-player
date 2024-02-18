package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/leonhsi/music-player/db/sqlc"
)

type createSongRequest struct {
  SongName string `json:"song_name" binding="required"`
  ArtistName string `json:"artist_name" binding="required"`
  ArtistID int64 `json:"artist_id" binding="required"`
  ThumbnailS3Path string `json:"thumbnail_s3_path"`
  Mp3S3Path string `json:"mp3_s3_path"`
}

func (server *Server) createSong(ctx *gin.Context) {
	var req createSongRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  // check if artist exist
  var artist db.Artist
  artist, err := server.store.GetArtist(ctx, req.ArtistName)
  if err != nil {
    artist, err = server.store.CreateArtist(ctx, req.ArtistName)
    if err != nil {
	    ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	    return
    }
  }

  arg := db.CreateSongParams{
    SongName: req.SongName,  
    ArtistID: artist.ArtistID,
    ArtistName: req.ArtistName,
    ThumbnailS3Path: req.ThumbnailS3Path,
    Mp3S3Path: req.Mp3S3Path,
  }

	song, err := server.store.CreateSong(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

type getSongByNameRequest struct {
  SongName string `uri:"name" binding:"required"`
}

func (server *Server) getSongByName(ctx *gin.Context) {
	var req getSongByNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.SongName

	song, err := server.store.GetSongByName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

type getSongByIDRequest struct {
  SongID int64 `uri:"id" binding:"required"`
}

func (server *Server) getSongByID(ctx *gin.Context) {
	var req getSongByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.SongID

	song, err := server.store.GetSongByID(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}

type listSongsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (server *Server) listSongs(ctx *gin.Context) {
	var req listSongsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

  arg := req.PageID 

	song, err := server.store.ListSongs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, song)
}
