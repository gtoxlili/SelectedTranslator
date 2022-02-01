package xclip

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunShell(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// 阻塞
	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		fmt.Println(err.Error())
		return ""
	}
	return out.String()
}
