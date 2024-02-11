package main

import (
  "log"
  "database/sql"

  "github.com/leonhsi/music-player/api"
  db "github.com/leonhsi/music-player/db/sqlc"
  "github.com/leonhsi/music-player/utils"
  _ "github.com/lib/pq"
)

func main() {
  config, err := utils.LoadConfig(".")
  if err != nil {
    log.Fatal("cannot load config:", err)
  }

  conn, err := sql.Open(config.DBDriver, config.DBSource)
  if err != nil {
    log.Fatal("cannot connect to db:", err)
  }

  store := db.NewStore(conn)
  server := api.NewServer(store)

  utils.InitDB(store);

  err = server.Start(config.ServerAddress)
}
