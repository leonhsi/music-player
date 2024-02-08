import React, { useState, useRef, useEffect } from "react";
import {
  BiPlayCircle,
  BiPauseCircle,
  BiSkipPreviousCircle,
  BiSkipNextCircle,
} from "react-icons/bi";

export default function Player({
  currentSong,
  currentIndex,
  nextSong,
  prevSong,
}) {
  const [isPlaying, setIsPlaying] = useState(false);
  const audioRef = useRef(null);
  const togglePlay = () => {
    setIsPlaying(!isPlaying);
  };

  useEffect(() => {
    if (isPlaying) {
      audioRef.current.play();
    } else {
      audioRef.current.pause();
    }
  }, [isPlaying, currentIndex]);

  return (
    <div>
      <audio ref={audioRef} src={currentSong.music}></audio>
      <div className="player-card">
        <div className="image-container">
          <img className="music-image" src={currentSong.cover} alt="Music" />
        </div>
        {currentSong ? (
          <div>
            <h1 className="activeSong-name">{currentSong.name}</h1>
            <h2 className="activeArtist-name">{currentSong.creator}</h2>
          </div>
        ) : (
          ""
        )}
        <div className="control-icon">
          <BiSkipPreviousCircle
            className="icons"
            color="#2196f3"
            size={50}
            onClick={prevSong}
          />
          {isPlaying ? (
            <BiPauseCircle
              className="icons"
              color="#2196f3"
              size={70}
              onClick={togglePlay}
            />
          ) : (
            <BiPlayCircle
              className="icons"
              color="#2196f3"
              size={70}
              onClick={togglePlay}
            />
          )}
          <BiSkipNextCircle
            className="icons"
            color="#2196f3"
            size={50}
            onClick={nextSong}
          />
        </div>
      </div>
    </div>
  );
}
