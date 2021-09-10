package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"github.com/safchain/ethtool"
	"github.com/shirou/gopsutil/net"
)

type NetWordMessage struct {
	Name  string  `json:"name"`
	Speed float32 `json:"speed"`
	Usage float32 `json:"usage"`
}

// Run the cli command to obtain the nic speed
func LinuxSpeed(name string) float32 {
	file, _ := ioutil.ReadFile("/sys/class/net/" + name + "/speed")
	re := regexp.MustCompile(`[0-2]+[0-2]`)
	match := re.FindString(string(file))
	speed, _ := strconv.Atoi(match)
	return float32(speed)
}

// Calculate network card speed
func SpeedComputation(name string, speed float32) (float32, float32) {

	var times int
	var BytesSent1 float32
	var BytesRecv1 float32
	//to caculate percentage
	speedLen := speed / 100
	info, _ := net.IOCounters(true)

	for i := 0; i < len(info); i++ {
		if info[i].Name == name {
			times = i
			enfo := info[i]
			BytesSent1 = float32(enfo.BytesSent)
			BytesRecv1 = float32(enfo.BytesRecv)
		}
	}

	time.Sleep(time.Second)
	info, _ = net.IOCounters(true)
	enfo := info[times]
	BytesSent2 := float32(enfo.BytesSent)
	BytesRecv2 := float32(enfo.BytesRecv)
	send := (BytesSent2 - BytesSent1) * 8
	rec := (BytesRecv2 - BytesRecv1) * 8
	sp := (send + rec) / 1048576
	usage := sp / speedLen

	return sp, usage

}

//Check whether the NIC is a physical NIC and query the nic speed
func NetC() {

	var message []NetWordMessage
	//
	baseNicPath := "/sys/class/net/"
	cmd := exec.Command("ls", baseNicPath)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}
	output := string(buf)
	//
	for _, device := range strings.Split(output, "\n") {
		if len(device) > 1 {
			ethHandle, err := ethtool.NewEthtool()
			if err != nil {
				panic(err.Error())
			}
			defer ethHandle.Close()
			me1 := strings.Contains(device, "enp")
			me2 := strings.Contains(device, "eth")
			if me1 || me2 {
				netSpeed := LinuxSpeed(device)

				Speed, Usage := SpeedComputation(device, netSpeed)

				message = append(message, NetWordMessage{device, Speed, Usage})
			}
		}
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		klog.Errorf("err: %v", err)
	}
	fmt.Println(string(bytes))
}

func main() {
	NetC()
}
