package snowflake

import (
	"nes/pkg/log"
	"strconv"

	"github.com/sony/sonyflake"
)

// 雪花ID工具

var _sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	_sf = sonyflake.NewSonyflake(st)
}

func NextId() string {
	id, err := _sf.NextID()
	if err != nil {
		log.Logger.Errorf("snowflake NextId ：%s", err.Error())
	}
	return strconv.FormatUint(id, 10)
}
