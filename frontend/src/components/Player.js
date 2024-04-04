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
  
  const [currentTime, setCurrentTime] = useState(0); 
  
  useEffect(() => {
    const audioElement = audioRef.current;
    if (isPlaying) {
      audioElement.play();
    } else {
      audioElement.pause();
    }
    // get audio currentTime change event and update the value
    const updateTime = () => {
      if (audioElement.readyState === 4) {
        setCurrentTime(audioElement.currentTime);
      } 
    };
    audioElement.addEventListener("timeupdate", updateTime);
    return () => {
      audioElement.removeEventListener("timeupdate", updateTime);
    };
    
  }, [isPlaying, currentIndex]);
  
  const duration = audioRef.current?.duration ?? 0;

  if (audioRef.current) console.log("duration", duration, "current", currentTime);

  return (
    <div>
      <audio ref={audioRef} src={currentSong.music}></audio>
      <div className="player-card">
        <div className="image-container">
          <img className="music-image" src={currentSong.cover} alt="Music" />
        </div>
        {currentSong ? (
          <div>
            <h1 className="activeSong-name">{currentSong.song_name}</h1>
            <h2 className="activeArtist-name">{currentSong.artist_name}</h2>
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
        <div>
            <progress className="progress-bar" 
              value={currentTime} 
              max={duration}>
            </progress>
          </div>
      </div>
    </div>
  );
}
