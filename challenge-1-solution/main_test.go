package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	updateMessage("Hello World", &wg)

	wg.Wait()

	if msg != "Hello World" {
		t.Errorf("Expected to find 'Hello World', but got unexpected result: %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello World"
	printMessage()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello World") {
		t.Errorf("Expected to find 'Hello World', but got unexpected result: %s", msg)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find 'Hello, universe!', but got unexpected result: %s", msg)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find 'Hello, universe!', but got unexpected result: %s", msg)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find 'Hello, universe!', but got unexpected result: %s", msg)
	}
}
