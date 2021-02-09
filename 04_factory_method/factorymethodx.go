package factorymethod

///-------------------------------------------
//抽象产品：提供了产品的接口
type Product  interface {
	SetA(int)
	SetB(int)
	Show() int
}

//ProductBase
type ProductBase struct {
	a, b int
}
func (o *ProductBase) SetA(a int) {
	o.a = a
}
func (o *ProductBase) SetB(b int) {
	o.b = b
}

//Product1
type Product1 struct {
	*ProductBase
}
func (o Product1) Show() int {
	return o.a + o.b
}

//Product2
type Product2 struct {
	*ProductBase
}
func (o Product2) Show() int {
	return o.a - o.b
}
///-------------------------------------------

//抽象工厂：提供了厂品的生成方法
type AbstractFactory interface {
	Create() Product
}

//Product1Factory 是 Product1 的工厂类
type Product1Factory struct{}
func (Product1Factory) Create() Product {
	return &Product1{
		ProductBase: &ProductBase{},
	}
}

//Product2Factory 是 Product2 的工厂类
type Product2Factory struct{}
func (Product2Factory) Create() Product {
	return &Product2{
		ProductBase: &ProductBase{},
	}
}
//-------------------------------------------------------


