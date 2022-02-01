package xclip

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	res, err := http.Get("https://raw.githubusercontent.com/gtoxlili/SelectedTranslator/main/lib/xclip")
	if err != nil {
		panic("xclip download failed")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile(filepath.Join(os.Getenv("HOME"), ".xclip", "xclip"), body, 0755)
	fmt.Println("xclip is set to local")
}

func init() {
	if !xclipIsExistPath() && !xclipIsExistLocal() {
		fmt.Println("xclip is not installed, try to install it")
		setXclipToLocal()
	}
}

func GetSelection() string {
	return RunShell(filepath.Join(os.Getenv("HOME"), ".xclip", "xclip"), "-selection primary", "-out")
}
