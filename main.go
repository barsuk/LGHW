package main

import "fmt"

type R interface {
	reduce() interface{}
}

type Num struct {
	Value int
}

type Add struct {
	Left interface{}
	Right interface{}
}

type Boolean struct {
	Value bool
}

type LessThan struct {
	Left interface{}
	Right interface{}
}

func (a *Add) reduce() interface{} {
	if left, ok := (a.Left).(R); ok {
		return Add{ left.reduce(), a.Right, }
	} else if right, ok := (a.Right).(R); ok {
		return Add{ a.Left, right.reduce(), }
	} else {
		return Num{ a.Left.(int) + a.Right.(int) }
	}
}

func (lt *LessThan) reduce() interface{} {
	if left, ok := (lt.Left).(R); ok {
		return LessThan{ left.reduce(), lt.Right, }
	} else if right, ok := (lt.Right).(R); ok {
		return LessThan{ lt.Left, right.reduce(), }
	} else {
		return Boolean{ lt.Left.(int) < lt.Right.(int) }
	}
}

type Machine struct {
	Expr interface{}
}

func (m *Machine) step() {
	t := m.Expr.(R)
	m.Expr = t.reduce()
}

func (m *Machine) run() {
	_, ok := (m.Expr).(R)
	for {
		fmt.Println(m.Expr)
		m.step()

		_, ok = (m.Expr).(R)
		if !ok {
			break
		}
	}
	fmt.Println("last")
	fmt.Println(m.Expr)
}

func main() {
	fmt.Println("hola, el mundo")
	m := &Machine{
		&LessThan{
			&Add{1,3},
			3,
		},
	}
	m.run()
}
