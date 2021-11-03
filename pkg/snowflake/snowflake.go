package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

// 生成ID值为 int64位
func GenID() int64 {
	return node.Generate().Int64()
}
