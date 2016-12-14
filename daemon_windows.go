//+build windows
package daemon

//是否存在进程
import (
	"io/ioutil"
	"os"
	"strconv"
)

//判断指定pid的进程是否存在
func IsExistProcess(pidPath string) bool {
	data, err := ioutil.ReadFile(pidPath)
	if err != nil {
		return false
	}
	pid, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return false
	}
	_, err = os.FindProcess(int(pid))
	if err != nil {
		return false
	}
	return true
}
