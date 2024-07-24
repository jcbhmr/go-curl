package main

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"os/exec"
	"runtime"
)

//go:generate curl -LO https://cosmo.zip/pub/cosmos/v/3.3.1/bin/curl
//go:embed curl
var curl []byte

func main() {
	log.SetFlags(0)

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("os.Executable(): %v", err)
	}
	err = os.Rename(exe, exe+".DELETEME")
	if err != nil {
		log.Fatalf("os.Rename() %s => %s: %v", exe, exe+".DELETEME", err)
	}
	err = os.WriteFile(exe, curl, 0755)
	if err != nil {
		log.Fatalf("os.WriteFile() %s %d bytes %d", exe, len(curl), 0755)
	}
	_ = os.Remove(exe + ".DELETEME")

	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		args := []string{"-c", `exec "$0" "$@"`, exe}
		args = append(args, os.Args[1:]...)
		cmd = exec.Command("sh", args...)
	} else {
		cmd = exec.Command(exe, os.Args[1:]...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	} else if err != nil {
		log.Fatalf("cmd.Run() %s: %v", cmd, err)
	}
}
