package main_test

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	var buffer bytes.Buffer
	cmd := exec.Command("go", "run", ".", "--version")
	cmd.Stdout = &buffer
	t.Logf("$ %s", cmd)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("cmd.Run() %s: %v", cmd, err)
	}

	match := regexp.MustCompile(`^curl (\d+\.\d+\.\d+)`).FindStringSubmatch(buffer.String())
	version := match[1]
	t.Logf("embedded curl version=%s", version)
	if version != "8.4.0" {
		t.Fatalf("actual=%v expected=%v", version, "8.4.0")
	}
}

func TestGet(t *testing.T) {
	var buffer bytes.Buffer
	cmd := exec.Command("go", "run", ".", "-fsSL", "https://example.org/")
	cmd.Stdout = &buffer
	t.Logf("$ %s", cmd)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("cmd.Run() %s: %v", cmd, err)
	}

	t.Logf("received %d bytes", buffer.Len())
	if buffer.Len() < 100 {
		t.Fatalf("actual=%v expected=%v", buffer.Len(), ">=100")
	}
}
