package xclip

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func xclipIsExistPath() bool {
	if _, err := exec.LookPath("xclip"); err != nil {
		return false
	}
	return true
}

func xclipIsExistLocal() bool {
	if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".xclip", "xclip")); err != nil {
		return false
	}
	return true
}

func setXclipToLocal() {
	err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".xclip"), 0755)
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	if !xclipIsExistPath() && !xclipIsExistLocal() {
		fmt.Println("xclip is not installed")
		setXclipToLocal()
	}
}

func GetSelection() string {
	return RunShell("./lib/xclip -selection primary -out")
}
