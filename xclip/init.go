package xclip

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var xclipPath string

func xclipIsExistPath() bool {
	if path, err := exec.LookPath("xclip"); err != nil {
		return false
	} else {
		xclipPath = path
		return true
	}
}

func xclipIsExistLocal() bool {
	path := filepath.Join(os.Getenv("HOME"), ".selectedtranslator", "xclip")
	if _, err := os.Stat(path); err != nil {
		return false
	}
	xclipPath = path
	return true
}

func setXclipToLocal() {
	err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".selectedtranslator"), 0755)
	if err != nil {
		panic(err.Error())
	}
	res, err := http.Get("https://raw.githubusercontent.com/gtoxlili/SelectedTranslator/main/lib/xclip")
	if err != nil {
		panic("xclip download failed")
	}
	defer res.Body.Close()
	path := filepath.Join(os.Getenv("HOME"), ".selectedtranslator", "xclip")
	body, _ := ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile(path, body, 0755)
	xclipPath = path
	fmt.Println("xclip is set to local")
}

func init() {
	if !xclipIsExistPath() && !xclipIsExistLocal() {
		fmt.Println("xclip is not installed, try to install it")
		setXclipToLocal()
	}
}
