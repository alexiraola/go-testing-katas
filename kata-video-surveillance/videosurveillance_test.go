package main

import "testing"

type StubSensorDetectingMotion struct {
	motionDetected bool
}

func (d StubSensorDetectingMotion) isDetectingMotion() bool {
	return d.motionDetected
}

func (d *StubSensorDetectingMotion) changeMotionDetected(motionDetected bool) {
	d.motionDetected = motionDetected
}

type SpyRecorder struct {
	stopCalls  int
	startCalls int
}

func (r *SpyRecorder) startRecording() {
	r.startCalls++
	println("start recording...")
}

func (r *SpyRecorder) stopRecording() {
	r.stopCalls++
	println("stop recording...")
}

func TestAsksStopWhenNoMotion(t *testing.T) {
	sensor := &StubSensorDetectingMotion{}
	recorder := &SpyRecorder{startCalls: 0, stopCalls: 0}

	controller := SurveillanceController{sensor, recorder}
	controller.recordMotion()

	if recorder.stopCalls != 1 {
		t.Fatalf(`expected %d, got %d`, 1, recorder.stopCalls)
	}
}

func TestAsksStartWhenDetectsMotion(t *testing.T) {
	sensor := &StubSensorDetectingMotion{true}
	recorder := &SpyRecorder{startCalls: 0, stopCalls: 0}

	controller := SurveillanceController{sensor, recorder}
	controller.recordMotion()

	if recorder.startCalls != 1 {
		t.Fatalf(`expected %d, got %d`, 1, recorder.startCalls)
	}
}
