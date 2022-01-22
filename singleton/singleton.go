package singleton

import (
	"fmt"
	"sync"
)

// 单例模式 golang 实现
/*
概念
单例模式：一个只允许创建一个实例的类叫做单例类，这种模式称为单例模式

作用
类的实例具有全局唯一性，适用于一些场景，如管理配置类，写log的类等，这种类只需要一个全局的实例来共享资源，无需创建多个实例，
单例模式有利用节约资源、防止多个实例产生冲突。

实现方式有懒汉式和饿汉式之分。
懒汉式：需要的时候才创建，在这种写法中，需要注意的是要注意并发调用创建函数的场景，
		需要用锁或其他并发控制原语来控制实例只会创建一次。延迟加载
饿汉式：事先创建好，需要时直接返回。
*/

var (
	lock 	sync.Mutex
	instance 	*singleton
)

type singleton struct {
	Data 	int
	name 	string
}

func (s *singleton) String() string  {
	return fmt.Sprintf("singlenton data:%v,name:%v",s.Data,s.name)
}

// 懒汉式
func NewSingleton(data int,name string) *singleton  {
	if instance == nil {
		lock.Lock()
		// 双重判断
		if instance == nil {
			instance = &singleton{
				Data: data,
				name: name,
			}
		}
		lock.Unlock()
	}
	return instance
}


// 使用 sync 标准库实现
// sync 库里面有 Once类，可以控制某个逻辑只会执行一次，需要单例时直接使用该方式会很方便。

var once sync.Once
func GetInstance() *singleton  {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance

}
