package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"k8s.io/klog/v2"
)

func Cmd(commandName string, para string) (string, error) {
	cmd := exec.Command(commandName, para)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

func Execute1(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		klog.Errorf("output: %s, err: %v", string(output), err)
		return string(output), err
	}

	return string(output), nil
}

func main() {
	detectPort("  8880")
}
func detectPort(value string) string {
	buf, _ := Execute1("netstat", "-n")
	re := regexp.MustCompile(`[0-9]+[0-9]`)
	pname := re.FindString(value)

	fmt.Println("aim: ", pname)

	klog.Infof("port: %v ", pname)
	mes := regexp.MustCompile(pname)
	match := mes.MatchString(buf)
	klog.Infof("match: %v", match)
	if match {
		return "true"
	} else {
		return "false"
	}
}
