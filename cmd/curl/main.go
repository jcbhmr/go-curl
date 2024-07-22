package main

import (
	_ "embed"
	"errors"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/jcbhmr/go-curl/v8/internal"
	"github.com/jcbhmr/go-curl/v8/internal/curl"
)

func main() {
	log.SetFlags(0)

	var exeExt string
	if runtime.GOOS == "windows" {
		exeExt = ".exe"
	} else {
		exeExt = ""
	}

	curlPath, err := xdg.DataFile("go-curl/v8/curl-" + internal.Version + exeExt)
	if err != nil {
		log.Fatal(err)
	}

	stats, err := os.Stat(curlPath)
	if errors.Is(err, fs.ErrNotExist) {
		err := os.WriteFile(curlPath, curl.Curl, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		if stats.Mode().Perm() != 0755 {
			err = os.Chmod(curlPath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if exe, err := os.Executable(); err == nil {
		err := os.Rename(exe, exe + ".DELETEME")
		if err == nil {
			err := os.Symlink(curlPath, exe)
			if err == nil {
				_ = os.Remove(exe + ".DELETEME")
			} else {
				err := os.Rename(exe + ".DELETEME", exe)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		args := []string{"-c", `exec "$0" "$@"`, curlPath}
		args = append(args, os.Args[1:]...)
		cmd = exec.Command("sh", args...)
	} else {
		cmd = exec.Command(curlPath, os.Args[1:]...)
		cmd.Args[0] = os.Args[0]
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	} else if err != nil {
		log.Fatal(err)
	}
}
