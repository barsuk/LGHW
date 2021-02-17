package oldmain

import "fmt"

type R interface {
	Reduce() interface{}
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
	return fmt.Sprintf("%d + %d", a.Left, a.Right)
}

func (a Add) Reduce() interface{} {
	if l, ok := a.Left.(R); ok {
		return Add{
			l.Reduce(),
			a.Right,
		}
	} else if r, ok := a.Right.(R); ok {
		return Add{
			a.Left,
			r.Reduce(),
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
	return fmt.Sprintf("%d < %d", lt.Left, lt.Right)
}

func (lt LessThan) Reduce() interface{} {
	if l, ok := lt.Left.(R); ok {
		return LessThan{
			l.Reduce(),
			lt.Right,
		}
	} else if r, ok := lt.Right.(R); ok {
		return LessThan{
			lt.Left,
			r.Reduce(),
		}
	} else {
		return Boolean{
			int(lt.Left.(Num).Value) < int(lt.Right.(Num).Value),
		}
	}
}

type Machine struct {
	Expr *interface{}
}

func (m *Machine) step() {
	var res interface{}
	expr := *m.Expr
	fmt.Printf("BEFORE %s\n", expr)
	if e, ok := expr.(R); ok {
		res = e.Reduce()
	}
	m.Expr = &res
}

func (m Machine) run() {
	var check bool
	for _, check = (*m.Expr).(R); check; {
		m.step()
		_, check = (*m.Expr).(R)
	}
	fmt.Printf("AFTER %s\n", *m.Expr)
}

func main() {
	fmt.Println("hola, el mundo")

	var expr interface{}
	expr = LessThan{
		Left:  Add{Num{1}, Num{3},},
		Right: Num{3,},
	}

	m := Machine{&expr, }
	m.run()

	//
	expr = LessThan{
		Left:  Add{Num{1}, Num{3},},
		Right: Add{Num{4}, Num{5},},
	}

	m = Machine{&expr, }
	m.run()

	//
	expr = LessThan{
		Left: Add{Num{1}, Num{3},},
		Right: Add{
			Add{
				Num{8}, Num{0},
			},
			Num{5},
		},
	}

	m = Machine{&expr, }
	m.run()
}
