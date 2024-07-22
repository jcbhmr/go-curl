//go:build ignore

package main

import (
	"bytes"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

func main() {
	log.SetFlags(0)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	curlName := "curl/curl"
	curlPath, err := filepath.Abs(curlName)
	if err != nil {
		log.Fatalf("filepath.Abs() %s: %v", curlName, err)
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("sh", "-c", `exec "$0" "$@"`, curlPath, "--version")
	} else {
		cmd = exec.Command(curlPath, "--version")
	}
	var stdoutBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	slog.Debug("cmd.Run()", "cmd", cmd.String())
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() %s: %v", cmd, err)
	}

	re := regexp.MustCompile(`curl (\d+\.\d+\.\d+)`)
	match := re.FindStringSubmatch(stdoutBuffer.String())
	if match == nil {
		log.Fatalf("%s output does not match %s", cmd, re)
	}

	versionName := "VERSION"
	slog.Debug("os.WriteFile()", "name", versionName, "data", match[1], "perm", 0644)
	err = os.WriteFile(versionName, []byte(match[1]), 0644)
	if err != nil {
		log.Fatalf("os.WriteFile() %s: %v", versionName, err)
	}
}
