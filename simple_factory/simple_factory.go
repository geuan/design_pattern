package simple_factory

import "fmt"

// 简单工厂模式实质是由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类（这些产品类继承自一个父类或接口）的实例


// go 语言没有构造函数一说，所以一般会定义 NewXXX 函数来初始化相关类。
// NewXXX 函数返回接口时就是简单工厂模式，也就是说 Golang的一般推荐做法就是简单工厂。
// 在这个 simplefactory包中只有API接口和NewAPI函数为包外可见，封装了实现细节。

type API interface {
	Say(name string) string
}

func NewAPI(t int) API  {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {}

func (*hiAPI) Say(name string) string  {
	return fmt.Sprintf("Hi,%s",name)
}

type helloAPI struct {}

func (*helloAPI) Say(name string) string  {
	return fmt.Sprintf("Hello,%s",name)
}





