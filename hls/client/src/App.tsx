import { useEffect, useRef, useState } from 'react'
import './App.css'
import Hls from 'hls.js'

function App() {

  const [song, setSong] = useState("candyland")

  const audioRef = useRef<HTMLAudioElement>(null)

  useEffect(() => {
    if (!audioRef || !audioRef.current)
      return

    if(Hls.isSupported()){
      const hls = new Hls
      hls.loadSource(`http://localhost:8080/${song}/outputlist.m3u8`)
      hls.attachMedia(audioRef.current)
    }else if (audioRef.current.canPlayType('application/vnd.apple.mpegurl')) {
      audioRef.current.src = `http://localhost:8080/${song}/outputlist.m3u8`
    }
  }, [audioRef, song])

  return (
    <div style={{
      display: "flex",
      flexDirection: "column",
      gap: "16px"
    }}>
      <div style={{
        display: 'flex',
        flexDirection: "column"
      }}>
      <label htmlFor='song'>Select a song to play:</label>
      <select id="song" name='Song' value={song} onChange={(e) => setSong(e.target.value)}>
        <option value={"candyland"}>Candyland</option>
        <option value={"dawn"}>Dawn</option>
      </select>
      </div>
      <audio ref={audioRef} controls></audio>
    </div>
  )
}

export default App
