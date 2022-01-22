package singleton

import "testing"

func TestNewSingleton(t *testing.T) {
	instance1 := NewSingleton(1,"hello")
	// 查看其内存地址
	t.Logf("%p,%v",instance1,instance1)

	instance2 := NewSingleton(2,"hi")
	t.Logf("%p,%v",instance2,instance2)
}
