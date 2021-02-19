package smysl

import "fmt"

type R interface {
	Reduce(env map[string]interface{}) interface{}
}

type Num struct {
	Value int
}

func (n Num) String() string {
	return fmt.Sprintf("%d", n)
}

type Add struct {
	Left interface{}
	Right interface{}
}

func (a Add) String() string {
	//return fmt.Sprintf("%d + %d", a.Left, a.Right)
	return fmt.Sprintf("%v + %v", a.Left, a.Right)
}

func (a Add) Reduce(env map[string]interface{}) interface{} {
	if l, ok := a.Left.(R); ok {
		return Add{
			l.Reduce(env),
			a.Right,
		}
	} else if r, ok := a.Right.(R); ok {
		return Add{
			a.Left,
			r.Reduce(env),
		}
	} else {
		return Num{
			a.Left.(Num).Value + a.Right.(Num).Value,
		}
	}
}

type Boolean struct {
	Value bool
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}

type LessThan struct {
	Left interface{}
	Right interface{}
}

func (lt LessThan) String() string {
	return fmt.Sprintf("%v < %v", lt.Left, lt.Right)
}

func (lt LessThan) Reduce(env map[string]interface{}) interface{} {
	if l, ok := lt.Left.(R); ok {
		return LessThan{
			l.Reduce(env),
			lt.Right,
		}
	} else if r, ok := lt.Right.(R); ok {
		return LessThan{
			lt.Left,
			r.Reduce(env),
		}
	} else {
		return Boolean{
			int(lt.Left.(Num).Value) < int(lt.Right.(Num).Value),
		}
	}
}

type Machine struct {
	Expr *interface{}
	Env map[string]interface{}
}

func (m *Machine) step(env map[string]interface{}) {
	var res interface{}
	expr := *m.Expr
	fmt.Printf("STEP %s\n", expr)
	if e, ok := expr.(R); ok {
		res = e.Reduce(env)
	}
	m.Expr = &res
}

func (m Machine) run() {
	var check bool
	for _, check = (*m.Expr).(R); check; {
		m.step(m.Env)
		_, check = (*m.Expr).(R)
	}
	fmt.Printf("END WITH %s\n", *m.Expr)
}

type Variable struct {
	Name string
}

func (v Variable) String() string {
	return fmt.Sprintf("%v", v.Name)
}

func (v Variable) Reduce(env map[string]interface{}) interface{} {
	return env[v.Name]
}