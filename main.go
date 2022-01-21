package main

import (
	"design_patterns/factory_method"
)

func main()  {
	// 抽象工厂
	var af factory_method.AnimalFactory

	// 狗工厂
	// 由于实现 AnimalFactory 接口的接收者都是指针类型，所以是对应的指针实现了接口，这里需要加上 &
	af = &factory_method.DogFactory{}
	dog := af.GetAnimal()
	dog.SetInfo("哈士奇","张三")
	dog.Speak()


	// 猫工厂
	af = &factory_method.CatFactory{}
	cat := af.GetAnimal()
	cat.SetInfo("英短","李四")
	cat.Speak()


	// 猪工厂
	af = &factory_method.PigFactory{}
	pig := af.GetAnimal()
	pig.SetInfo("猪英俊","王五")
	pig.Speak()
}