package idgen

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	//if sf == nil {
	//	panic("sonyflake not created")
	//}
}

func Next() (id uint64, err error) {
	var i uint64
	i, err = sf.NextID()
	if err == nil {
		id = i
	}
	return
}

func NextString() (id string, err error) {
	var i uint64
	i, err = sf.NextID()
	if err == nil {
		id = string(i)
	}
	return
}
