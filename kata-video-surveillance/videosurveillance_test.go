package main

import "testing"

type FakeSensor struct {
	motionDetected bool
}

func (d FakeSensor) isDetectingMotion() bool {
	return d.motionDetected
}

func (d *FakeSensor) changeMotionDetected(motionDetected bool) {
	d.motionDetected = motionDetected
}

type FakeRecorder struct {
	stopCalls  int
	startCalls int
}

func (r *FakeRecorder) startRecording() {
	r.startCalls++
	println("start recording...")
}

func (r *FakeRecorder) stopRecording() {
	r.stopCalls++
	println("stop recording...")
}

func TestAsksStopWhenNoMotion(t *testing.T) {
	sensor := &FakeSensor{}
	recorder := &FakeRecorder{startCalls: 0, stopCalls: 0}

	controller := SurveillanceController{sensor, recorder}
	controller.recordMotion()

	if recorder.stopCalls != 1 {
		t.Fatalf(`expected %d, got %d`, 1, recorder.stopCalls)
	}
}

func TestAsksStartWhenDetectsMotion(t *testing.T) {
	sensor := &FakeSensor{true}
	recorder := &FakeRecorder{startCalls: 0, stopCalls: 0}

	controller := SurveillanceController{sensor, recorder}
	controller.recordMotion()

	if recorder.startCalls != 1 {
		t.Fatalf(`expected %d, got %d`, 1, recorder.startCalls)
	}
}
