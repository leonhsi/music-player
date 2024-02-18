import { useState } from "react";
import "./App.css";
// import AudioFiles from "./components/AudioFiles";
import Player from "./components/Player";
import metadataJson from "./metadata.json";

const audios = JSON.parse(JSON.stringify(metadataJson));

for (var audio of audios) {
  audio.name = decodeURIComponent(audio.name);
  audio.music = require(`${audio.music_path}`);
  audio.cover = require(`${audio.cover_path}`);
}

function App() {
  const songs = audios;
  const audioLength = audios.length;
  const [currentIndex, setCurrentIndex] = useState(0);
  const [currentSong, setCurrentSong] = useState(songs[0]);

  const nextSong = () => {
    var nextIndex = (currentIndex + 1) % audioLength;
    setCurrentIndex(nextIndex);
    setCurrentSong(audios[nextIndex]);
    const resp = fetch("http://localhost:8080/songs/id/${nextIndex}", {
      method: "GET",
    }).catch((err) => {
      console.log(err);
    });
    const data = resp.json();
    console.log(data);
  };

  const prevSong = () => {
    var prevIndex = (currentIndex + audioLength - 1) % audioLength;
    setCurrentIndex(prevIndex);
    setCurrentSong(audios[prevIndex]);
  };

  const handleClick = () => {
    fetch("http://localhost:1111/goodbye", { method: "GET" })
      .then((res) => res.json())
      .catch((e) => {
        console.log(e);
      });
  };

  return (
    <>
      <div className="player-main">
        <Player
          currentSong={currentSong}
          currentIndex={currentIndex}
          nextSong={nextSong}
          prevSong={prevSong}
        />
        <button onClick={handleClick}>Button</button>
      </div>
    </>
  );
}

export default App;
