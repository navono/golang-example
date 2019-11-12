package gorm3

import (
	"github.com/jinzhu/gorm"
)

type (
	NetworkTemplate struct {
		gorm.Model
		UID  string `gorm:"unique_index;not null"`
		Name string

		// networkPb.NetType
		Type int32

		Config string
	}

	net2 struct {
		Addrs     []netElement
		NTPServer netElement
	}
	net3 struct {
		Addrs     []netElement
		NTPServer []netElement
	}

	ethernet struct {
		Addrs     []netElement
		NTPServer []netElement
	}

	serial struct {
		// 串口通信时的速率
		BaudRate int32
		Verify   string
		DataBit  string
		StopBit  string
	}

	netElement struct {
		Name          string `json:"name"`
		Ip            string `json:"ip"`
		CanBeDisabled bool   `json:"canBeDisabled"`
	}
)
