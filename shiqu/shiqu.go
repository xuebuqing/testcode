package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Cmd(commandName string, para string) {
	cmd := exec.Command("cp", commandName, para)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	time.Sleep(time.Second)
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func main() {
	ts := DetectPort("Asia/Shanghai")
	fmt.Println(ts)
}
func DetectPort(aim string) string {

	req := "/usr/share/zoneinfo/" + aim

	t, _ := PathExists(req)
	if t {
		Cmd(req, "/etc/localtime")
		return "yes"
	} else {
		return "no"
	}

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
