package design_pattern

import "time"

//三个工厂模式的区别，都和为了解决一个具体的问题。
//简单工厂 -开闭-> 每个产品一个工厂（工厂模式） -需要多个产品完成一次操作-> 一个工厂生产多个产品（抽象工厂）

//simple factory
type Phone interface {
	Call(phoneNum string) bool
}

type MiPhone struct {
}

func (m *MiPhone) Call(phoneNum string) bool {
	return true
}

//简单工厂,可以定义结构体，也可以直接定义函数

//方式一
// SimplePhoneFactory 获取产品的工厂结构体
type SimplePhoneFactory struct {
}

func (p *SimplePhoneFactory) Create(brandName string) Phone {
	switch brandName {
	case "Mi":
		return &MiPhone{}
	}
	return nil
}

// 方式二  获取产品的函数
func Create(brandName string) Phone {
	switch brandName {
	case "Mi":
		return &MiPhone{}
	}
	return nil
}

//工厂方法
type PhoneFactory interface {
	Create(brandName string) Phone
}

type MiPhoneFactory struct {
}

func (p *MiPhoneFactory) Create(brandName string) Phone {
	return &MiPhone{}
}

//工厂模式，为了避免"简单工厂方法" 在新增方法时需要修改工厂，所以就每个方法一个工厂，由客户端决定使用哪个
//这里不应该定义 GetPhoneFactory 这样就回到简单工厂模式，没有意义了
//func GetPhoneFactory(brandName string) PhoneFactory {
//	switch brandName {
//	case "Mi":
//		return &MiPhoneFactory{}
//	}
//	return nil
//}

type Watch interface {
	ReportTime() time.Time
}

type MiWatch struct {
}

func (m *MiWatch) ReportTime() time.Time {
	return time.Now()
}

type Factory interface {
	CreatePhone() Phone
	CreateWatch() Watch
}

type MiFactory struct {
}

func (p *MiFactory) CreatePhone() Phone {
	return &MiPhone{}
}
func (p *MiFactory) CreateWatch() Watch {
	return &MiWatch{}
}

// 这里也是不必要的，抽象工厂的目的就是一个工厂多个方法，比如数据库可能操作多个表。
// AbstractPhoneFactory 反倒有限制了
//type AbstractPhoneFactory struct {
//}
//
//func (a *AbstractPhoneFactory) Create(branchName string) Phone {
//	switch branchName {
//	case "Mi":
//		return (&MiFactory{}).CreatePhone()
//	}
//	return nil
//}
//
//func GetFactory(brandName string) Factory {
//	switch brandName {
//	case "Mi":
//		return &MiFactory{}
//	}
//	return nil
//}
