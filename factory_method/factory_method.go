package factory_method

import "fmt"

// 工厂方法模式

// 动物
type animal struct {
	name 	string
	master 	string
}

// 给动物设置基本信息
func (a *animal) SetInfo(name,master string) {
	a.name = name
	a.master = master
}

// 具体物种
type dog struct {
	animal
}

type cat struct {
	animal
}

type pig struct {
	animal
}


// 动物都有 Speake() 方法
type Speaker interface {
	Speak()
	SetInfo(name,master string)   // 这里实际上是调用 animal 类型的方法
}

func (d *dog) Speak()  {
	fmt.Printf("我叫[%v], 我的主人是[%s], 汪汪汪~\n", d.name, d.master)
}

func (c *cat) Speak()  {
	fmt.Printf("我叫[%v], [%v]才不是我的主人呢, 明明是我的铲屎官\n", c.name, c.master)
}

func (p *pig) Speak() {
	fmt.Printf("我叫[%v], 每天都是[%v]给饭我吃, 我吃饱了就想睡觉\n", p.name, p.master)
}


// 这里要注意，有一个抽象工厂的接口：

// 抽象工厂接口
type AnimalFactory interface {
	GetAnimal() Speaker
}


// 这个抽象工厂接口有不同的实现，每个实现都是一个创建具体动物的工厂

// 生产具体动物的具体工厂
type DogFactory struct {}
type CatFactory struct {}
type PigFactory struct {}

// 狗工厂
func (d *DogFactory) GetAnimal() Speaker  {
	return &dog{}
}

func (c *CatFactory) GetAnimal() Speaker {
	return &cat{}
}

func (p *PigFactory) GetAnimal() Speaker  {
	return &pig{}
}













