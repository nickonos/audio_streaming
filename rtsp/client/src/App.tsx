import './App.css'

function App() {
  return (
    <>
    <audio controls>
      {/* 
        On 11-1-2023, this does not work yet natively,
        https://developer.mozilla.org/en-US/docs/Web/Media/Audio_and_video_delivery/Live_streaming_web_audio_and_video#rtsp
       */}
      <source src='rtsp://localhost:8554/mystream' type='audio/mp3' />
    </audio>
    </>
  )
}

export default App
