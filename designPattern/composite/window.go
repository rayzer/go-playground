package main

import (
	"log"
)

type element interface {
	print()
}

type attributes struct {
	elementType string
	name        string
}

func (a attributes) print() {
	log.Println("print " + a.elementType + "(" + a.name + ")")
}

type button struct {
	attributes
}

type picture struct {
	attributes
}

type Label struct {
	attributes
}

type textBox struct {
	attributes
}

type passwordBox struct {
	attributes
}

type frame struct {
	attributes
	subElements
}

func (f *frame) print() {
	f.attributes.print()
	f.subElements.print()
}

func (f *frame) add(e element) {
	f.elements = append(f.elements, e)
}

type window struct {
	attributes
	subElements
}

type subElements struct {
	elements []element
}

func (s *subElements) add(e element) {
	s.elements = append(s.elements, e)
}

func (s subElements) print() {
	for _, e := range s.elements {
		e.print()
	}
}

func (w *window) print() {
	w.attributes.print()
	w.subElements.print()
}

func newWinForm() *window {
	instance := &window{
		attributes{
			"WinFrom",
			"WINDOW窗口",
		},
		subElements{},
	}

	instance.subElements.add(picture{attributes{"Picture", "LOGO图片"}})
	instance.subElements.add(button{attributes{"Button", "登录"}})
	instance.subElements.add(button{attributes{"Button", "注册"}})

	f := &frame{attributes{"FRAME", "FRAME1"}, subElements{}}
	f.subElements.add(Label{attributes{"Lable", "用户名"}})
	f.subElements.add(textBox{attributes{"TextBox", "文本框"}})
	f.subElements.add(Label{attributes{"Lable", "密码"}})
	f.subElements.add(passwordBox{attributes{"PasswordBox", "密码框"}})

	instance.add(f)
	return instance
}

func main() {
	log.SetFlags(0)
	winForm := newWinForm()
	winForm.print()
}
