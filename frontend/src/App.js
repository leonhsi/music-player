import { useState, useEffect } from "react";
import "./App.css";
import Player from "./components/Player";

const songMetaDataService = "http://localhost:8080/songs/";

function getSongCount() {
  return new Promise((resolve, reject) => {
    var url = songMetaDataService + "count/";
    fetch(url, { method: "GET" })
      .then((response) => response.json())
      .then((json) => {
        resolve(json);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

function getSongMetaDataByID(id) {
  return new Promise((resolve, reject) => {
    var song = {};
    var url = songMetaDataService + "id/" + id;

    fetch(url, { method: "GET" })
      .then((response) => response.json())
      .then((json) => {
        song.song_name = json.song_name;
        song.artist_name = json.artist_name;
        song.cover = require(`${json.thumbnail_s3_path}`);
        song.music = require(`${json.mp3_s3_path}`);
        resolve(song);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

function App() {
  var audioLength = getSongCount().then((songCount) => {
    audioLength = songCount;
  });
  const [currentIndex, setCurrentIndex] = useState(0);
  const [currentSong, setCurrentSong] = useState({});

  const nextSong = () => {
    var nextIndex = (currentIndex + 1) % audioLength;
    setCurrentIndex(nextIndex);

    getSongMetaDataByID(nextIndex + 1)
      .then((song) => {
        setCurrentSong(song);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const prevSong = () => {
    var prevIndex = (currentIndex + audioLength - 1) % audioLength;
    setCurrentIndex(prevIndex);

    getSongMetaDataByID(prevIndex + 1)
      .then((song) => {
        setCurrentSong(song);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  useEffect(() => {
    getSongMetaDataByID(1)
      .then((song) => {
        setCurrentSong(song);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  return (
    <>
      <div className="player-main">
        <Player
          currentSong={currentSong}
          currentIndex={currentIndex}
          nextSong={nextSong}
          prevSong={prevSong}
        />
      </div>
    </>
  );
}

export default App;
