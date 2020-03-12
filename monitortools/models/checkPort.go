package models

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func checkPort(hostIp , hostPort string) string {
	hostPortStr,err:=strconv.Atoi(hostPort)
	_, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%v",hostIp,hostPortStr),time.Duration(time.Second*1))
	if err !=nil {
		return "error"
	} else {
		return "ok"
	}
}
