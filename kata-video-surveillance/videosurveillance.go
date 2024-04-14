package main

type MotionSensor interface {
	isDetectingMotion() bool
}

type VideoRecorder interface {
	startRecording()
	stopRecording()
}

type SurveillanceController struct {
	sensor   MotionSensor
	recorder VideoRecorder
}

func (c *SurveillanceController) recordMotion() {
	c.recorder.stopRecording()
}
