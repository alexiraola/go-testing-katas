package library

import "testing"

func TestSayHello(t *testing.T) {
	text := sayHello()

	if text != "hi" {
		t.Fatalf(`sayHello() = %q`, text)
	}
}
