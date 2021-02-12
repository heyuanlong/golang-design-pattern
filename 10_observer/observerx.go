package observer

import "fmt"

//--------------------------------------------------
//抽象观察者（Observer）角色
type Observer interface {
	Update(string)
}
//--------------------------------------------------
//主题（Subject）角色
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Add(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s.context)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
//--------------------------------------------------
//具体观察者（Concrete Observer）角色
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s string) {
	fmt.Printf("receive %s\n", s)
}
