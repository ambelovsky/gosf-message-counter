package gosfmessagecounter

import (
	"os"
	"os/exec"
	"runtime"
)

// ConsoleClear erases all text in the server console
var ConsoleClear func()

func clearLinux() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearWindows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = clearLinux
	clear["darwin"] = clearLinux
	clear["windows"] = clearWindows

	ConsoleClear = func() {
		clear[runtime.GOOS]()
	}
}
