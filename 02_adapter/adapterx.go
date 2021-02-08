package adapter

//----------------------------------------------
//原始接口
type Origin interface {
	OriginMethod()string
}
func NewOriginStruct() *OriginStruct {
	return &OriginStruct{}
}
type OriginStruct struct {
}
func (ts *OriginStruct) OriginMethod() string  {
	return "origin"
}

//----------------------------------------------
//Target 是适配的目标接口
type Target interface {
	Request() string
}

//----------------------------------------------
//适配器
func NewAdapter(ori Origin) Target {
	return &adapter{
		Origin: ori,
	}
}
type adapter struct {
	Origin
}
func (a *adapter) Request() string {
	return a.OriginMethod()
}
