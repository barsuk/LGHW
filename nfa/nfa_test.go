package nfa

import (
	"LGHW/dka"
	"fmt"
	"testing"
)

var rulebook = RuleBook{Rules: []dka.FARule{
	{1, "a", 1},
	{1, "b", 1},
	{1, "b", 2},

	{2, "a", 3},
	{2, "b", 3},

	{3, "a", 4},
	{3, "b", 4},
}}

func TestDesign(t *testing.T) {
	fmt.Println(rulebook.nextStates(map[int]bool{1: false}, "b"))
	fmt.Println(rulebook.nextStates(map[int]bool{1: false, 2: false}, "a"))
	fmt.Println(rulebook.nextStates(map[int]bool{1: false, 3: false}, "b"))
}

func TestNFAAccepting(t *testing.T) {
	nfa := NFA{
		CurrentStates: map[int]bool{1:false},
		AcceptStates: []int{4},
		RuleBook:     rulebook,
	}
	fmt.Printf("%+v\n", nfa.accepting())
	nfa1 := NFA{
		CurrentStates: map[int]bool{1:false, 2: false, 4: false},
		AcceptStates: []int{4},
		RuleBook:     rulebook,
	}
	fmt.Printf("%+v\n", nfa1.accepting())
}

func TestStrReading(t *testing.T) {
	nfa := NFA{
		CurrentStates: map[int]bool{1:false},
		AcceptStates: []int{4},
		RuleBook:     rulebook,
	}
	fmt.Printf("STARTED new. Accepting? %t\n", nfa.accepting())
	for _, c := range []string{"b", "a", "b"} {
		nfa.readCharacter(c)
		fmt.Printf("char: %s, accepted: %t\n", c, nfa.accepting())
	}
	nfa = NFA{
		CurrentStates: map[int]bool{1:false},
		AcceptStates: []int{4},
		RuleBook:     rulebook,
	}
	fmt.Printf("STARTED new. Accepting? %t\n", nfa.accepting())
	str := "bbbbb"
	nfa.readString(str)
	fmt.Printf("str: %s. Accepting? %t\n", str, nfa.accepting())
}