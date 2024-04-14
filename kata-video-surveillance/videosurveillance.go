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
	if c.sensor.isDetectingMotion() {
		c.recorder.startRecording()
	} else {
		c.recorder.stopRecording()
	}
}
