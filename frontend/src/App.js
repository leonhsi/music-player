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
  };

  const prevSong = () => {
    var prevIndex = (currentIndex + audioLength - 1) % audioLength;
    setCurrentIndex(prevIndex);
    setCurrentSong(audios[prevIndex]);
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
      </div>
    </>
  );
}

export default App;
