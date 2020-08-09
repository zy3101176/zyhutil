package zlog

import (
	"testing"
)

func TestDebug(t *testing.T) {
	InitLogger("/Users/zhuyuanhan/learn/zyhutil/zyhutil/zlog/test.log", "debug")
	Debugf("test Debug %d", 1)
}
