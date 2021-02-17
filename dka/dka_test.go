package dka

import (
	"fmt"
	"testing"
)

var rulebook = DFARuleBook{Rules: []FARule{
	{1, "a", 2},
	{1, "b", 1},
	{2, "a", 2},
	{2, "b", 3},
	{3, "a", 3},
	{3, "b", 3},
}}

func TestDKA(t *testing.T) {
	print(rulebook.nextState(1, "a"), "\n")
	print(rulebook.nextState(1, "b"), "\n")
	print(rulebook.nextState(2, "b"), "\n")
}

func TestReadChar(t *testing.T) {
	dfa := DFA{
		CurrentState: 1,
		AcceptStates: []int{3},
		RuleBook:     rulebook,
	}
	fmt.Printf("%+v\n", dfa.accepting())
	for _, c := range []string{"a", "a", "b"} {
		dfa.readCharacter(c)
		fmt.Printf("char %s, curstate: %d, accepting: %t\n", c, dfa.CurrentState, dfa.accepting())
	}
}

func TestReadStringFalse(t *testing.T) {
	dfa := DFA{
		CurrentState: 1,
		AcceptStates: []int{3},
		RuleBook:     rulebook,
	}
	fmt.Printf("%+v\n", dfa.accepting())
	dfa.readString("baaa")
	fmt.Printf("%+v\n", dfa.accepting())
}

func TestReadStringTrue(t *testing.T) {
	dfa := DFA{
		CurrentState: 1,
		AcceptStates: []int{3},
		RuleBook:     rulebook,
	}
	fmt.Printf("%+v\n", dfa.accepting())
	dfa.readString("baaab")
	fmt.Printf("%+v\n", dfa.accepting())
}