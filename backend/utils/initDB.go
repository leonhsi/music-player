package utils

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "os"
  "context"

	db "github.com/leonhsi/music-player/db/sqlc"
)

type SongMetadata struct {
  Name string `json:"name"`
  Creator string `json:"creator"`
}

func InitDB(store db.Store) {
  fmt.Println("[DEBUG] Start initialize song metadata DB...")

  jsonFile, err := os.Open("./metadata.json")
  if err != nil {
    fmt.Println("cannot open json file:", err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var songs []SongMetadata
  err = json.Unmarshal(byteValue, &songs)
  if err != nil {
    fmt.Println("cannot unmarshal json file:", err)
  }

  for i := 0; i < len(songs); i++ {
    // check if artist exist var artist db.Artist
    artist, err := store.GetArtist(context.Background(), songs[i].Creator)
    if err != nil {
      // create artist
      artist, err = store.CreateArtist(context.Background(), songs[i].Creator)
      if err != nil {
        fmt.Println("cannot create artist:", err)
      }
    }
    
    // check if song exist
    _, err = store.GetSong(context.Background(), songs[i].Name) 
    if err != nil {
      // create songs
      arg := db.CreateSongParams{
        SongName: songs[i].Name,
        ArtistID: artist.ArtistID,      
        ArtistName: artist.ArtistName,
        ThumbnailS3Path: "",
        Mp3S3Path: "",
      }
      _, err = store.CreateSong(context.Background(), arg)
      if err != nil {
        fmt.Println("cannot create song:", err)
      }
    }
  }
  
  fmt.Println("[DEBUG] Finish initialization of", len(songs), "songs")
}
