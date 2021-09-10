package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/safchain/ethtool"
	"github.com/shirou/gopsutil/net"
)

func Cmd(commandName string, params string) (string, error) {
	cmd := exec.Command(commandName, params)
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

func main() {
	baseNicPath := "/sys/class/net/"
	cmd := exec.Command("ls", baseNicPath)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}
	output := string(buf)
	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			ethHandle, err := ethtool.NewEthtool()
			if err != nil {
				panic(err.Error())
			}
			defer ethHandle.Close()
			buf, _ := Cmd("ethtool", device)
			re := regexp.MustCompile("Speed:(.*)Mb/s")
			match := re.FindString(buf)
			if match != "" {
				re = regexp.MustCompile(`[0-2]+[0-2]`)
				match := re.FindString(match)
				speed, _ := strconv.Atoi(match)
				fmt.Println("speed is :", speed)
			}
		}
	}

	info, _ := net.IOCounters(true)
	for i := 0; i < len(info); i++ {
		if info[i].Name == "enp2s0" {
			fmt.Println(info[i])
		}
	}

}
