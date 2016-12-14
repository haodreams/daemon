//守护进程
package daemon

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//生成PID文件
func MakePID(pidPath string) (err error) {
	f, err := os.Create(pidPath)
	if err != nil {
		return
	}
	f.WriteString(fmt.Sprintf("%d", os.Getpid()))
	f.Close()
	return
}

//获取PID
func GetPID(pidPath string) (pid int, err error) {
	data, err := ioutil.ReadFile(pidPath)
	if err != nil {
		return
	}
	pid, err = strconv.Atoi(string(data))
	return
}

//注意: 参数中不能有空格,等带有分割标记的字符
func runProcess(binPath string) {
	ss := strings.Fields(binPath)
	var cmd *exec.Cmd
	if len(ss) > 1 {
		cmd = exec.Command(ss[0], ss[1:]...)
	} else {
		cmd = exec.Command(ss[0])
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return
	}
	cmd.Wait()
}

//守护
func Watch(binPath string) {
	ss := strings.Fields(binPath)
	pidPath := ss[0] + ".wj.pid"
	for {
		if !IsExistProcess(pidPath) {
			runProcess(binPath)
		}
		time.Sleep(time.Second)
	}
}

//启动守护
func Run() (err error) {
	//首先验证是否有写入权限
	test := "test.tmp"
	f, err := os.Create(test)
	if err != nil {
		log.Println("没有写入权限,程序停止运行")
		os.Exit(0)
		return
	}
	f.Close()
	os.Remove(test)
	deamonPath := os.Args[0] + ".wj.did"
	if IsExistProcess(deamonPath) {
		pidPath := os.Args[0] + ".wj.pid"
		if IsExistProcess(pidPath) {
			err = errors.New("Process exist")
			return
		}
		MakePID(pidPath)
		return
	} else {
		MakePID(deamonPath)
		cmd := ""
		for _, s := range os.Args {
			cmd += " " + s
		}
		Watch(cmd)
		err = errors.New("OK")
	}
	return
}
