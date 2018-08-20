package go_demo_lib

import "testing"

func TestHelloJoost(t *testing.T) {
	name := "Joost"
	expectedGreeting := "hello, " + name +"!"
	actualGreeting := Hello(name)
	if actualGreeting != expectedGreeting {
		t.Errorf("Greeting was incorrect, expected: %s, actual: %s", expectedGreeting, actualGreeting)
	}
}

func TestHelloWorld(t *testing.T) {
	name := "world"
	expectedGreeting := "hello, " + name +"!"
	actualGreeting := Hello(name)
	if actualGreeting != expectedGreeting {
		t.Errorf("Greeting was incorrect, expected: %s, actual: %s", expectedGreeting, actualGreeting)
	}
}

func TestHelloEmpty(t *testing.T) {
	var name string
	expectedGreeting := "hello, !"
	actualGreeting := Hello(name)
	if actualGreeting != expectedGreeting {
		t.Errorf("Greeting was incorrect, expected: %s, actual: %s", expectedGreeting, actualGreeting)
	}
}