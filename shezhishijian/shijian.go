package main

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/text/gstr"
)

func main() {
	Cmd("False") //you can choose the "False"or "True"
	UpdateSystemDate("2020-03-20 11:02")
}

func UpdateSystemDate(dateTime string) bool {
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			_, err1 := gproc.ShellExec(`date  ` + gstr.Split(dateTime, " ")[0])
			_, err2 := gproc.ShellExec(`time  ` + gstr.Split(dateTime, " ")[1])
			if err1 != nil && err2 != nil {
				glog.Info("更新系统时间错误:请用管理员身份启动程序!")
				return false
			}
			return true
		}
	case "linux":
		{
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Info("更新系统时间错误:", err1.Error())
				return false
			}
			return true
		}
	case "darwin":
		{
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Info("更新系统时间错误:", err1.Error())
				return false
			}
			return true
		}
	}
	return false
}

func Cmd(para string) {
	cmd := exec.Command("timedatectl", "set-ntp", para)
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
