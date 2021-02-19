package smysl

import "testing"

func TestMashine(t *testing.T) {
	var expr interface{}
	expr = LessThan{
		Left: Add{Num{1},Num{3},},
		Right: Num{3,},
	}

	m := Machine{ &expr, nil}
	m.run()

}

func TestLessThan(t *testing.T)  {
	var expr interface{}
	expr = LessThan{
		Left: Add{Num{1},Num{3},},
		Right: Add{Num{4},Num{5},},
	}

	m := Machine{ &expr, nil}
	m.run()
}

func TestVarExpr(t *testing.T)  {
	var expr interface{}
	//
	expr = LessThan{
		Left: Add{Num{1},Num{3},},
		Right: Add{
			Add{
				Variable{"x"},Num{0},
			},
			Variable{"y"},
		},
	}
	m := Machine{ &expr, map[string]interface{}{
		"x": Num{3},
		"y": Add{Num{8},Num{7},},
	}}
	m.run()
}
