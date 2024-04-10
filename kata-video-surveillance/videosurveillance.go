package main

interface MotionSensor {
  func isDetectingMotion() boolean
}

interface VideoRecorder {
  func startRecording()
  func stopRecording()
}
