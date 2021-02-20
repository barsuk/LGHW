package regx

import (
	"fmt"
	"testing"
)

/*
pattern = Repeat . new (
	Choose . new (
		Concatenate . new (
			Literal , new ( ' а ' ) ,
			Literal . new ( 'а ' )
			Literal . new ( ' b ' )
		)
	) ,
) => /(ab|a)*/

func TestPattern(t *testing.T) {
	pattern := Repeat{
		Pattern: Choose{
			First: Concatenate{
				First: Literal{
					Character: "a",
				},
				Second: Literal{
					Character: "b",
				},
			},
			Second: Literal{
				Character: "a",
			},
		},
	}
	fmt.Printf("%s\n", pattern)
}
