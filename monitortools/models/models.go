package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type HostStatus struct {
	Id         int64
	HostIp     string
	HostPort   string
	PortStatus string
}

func RegisterDb() {
	orm.RegisterModel(new(HostStatus))
	//orm.RegisterModel(new(monitorHosts))
	orm.RegisterDataBase("default", "mysql", "root:pppppp@tcp"+
		"(192.168.213.101:3306)/imooc?charset=utf8")
}

func GetHostStatus() ([]*HostStatus, error) {
	o := orm.NewOrm()
	hostStatus := make([]*HostStatus, 0)
	qs := o.QueryTable("host_status")
	var err error
	_, err = qs.All(&hostStatus)

	hostStatusNew := make([]*HostStatus, 0)
	for a,i :=range hostStatus {
		status:=checkPort(i.HostIp,i.HostPort)
		tmp:=&HostStatus{
			Id:int64(a),
			HostIp:     i.HostIp,
			HostPort:   i.HostPort,
			PortStatus: status,
		}
		hostStatusNew=append(hostStatusNew,tmp)
	}
	return hostStatusNew, err
}

func AddHosts(hostIp, port string) error {
	portStatus := checkPort(hostIp, port)
	o := orm.NewOrm()
	insertTmp := &HostStatus{
		HostIp:     hostIp,
		HostPort:   port,
		PortStatus: portStatus,
	}
	_, err := o.Insert(insertTmp)
	return err
}
