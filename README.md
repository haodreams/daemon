# daemon

package main

import (
	"fmt"
	"time"

	"github.com/haodreams/daemon"
)

func main() {
	if err := daemon.Run(); err != nil {	
		return		
	}
	
	fmt.Println("start")	
	var a *int	
	time.Sleep(time.Second * 5)	
	*a = 10	
}


//开机自启动 编辑 /etc/rc.local

/bin/rm -f /home/pi/test.wj.*

/home/pi/test&
