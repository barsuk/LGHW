package dka

import "fmt"

type FARule struct {
	State     int
	Character string
	NextState int
}

func (rule FARule) AppliesTo(state int, character string) bool {
	return rule.State == state && rule.Character == character
}

func (rule FARule) Follow() int {
	return rule.NextState
}

func (rule FARule) inspect() {
	fmt.Printf("#<FARule %d --%s--> %d", rule.State, rule.Character, rule.NextState)
}

type DFARuleBook struct {
	Rules []FARule
}

func (b DFARuleBook) nextState(state int, character string) int {
	return b.ruleFor(state, character).Follow()
}

func (b DFARuleBook) ruleFor(state int, character string) FARule {
	for _, rule := range b.Rules {
		if rule.AppliesTo(state, character) {
			return rule
		}
	}
	panic("нет такого правила")
}

type DFA struct {
	CurrentState int
	AcceptStates []int
	RuleBook DFARuleBook
}

func (dfa *DFA) accepting() bool {
	for _, v := range dfa.AcceptStates {
		if v == dfa.CurrentState {
			return true
		}
	}
	return false
}

func (dfa *DFA) readCharacter(character string) {
	//fmt.Printf("BEFORE cur st: %d\n", dfa.CurrentState)
	dfa.CurrentState = dfa.RuleBook.nextState(dfa.CurrentState, character)
	//fmt.Printf("AFTER cur st: %d\n", dfa.CurrentState)
}

func (dfa *DFA) readString(str string) {
	//fmt.Printf("reading string: %q\n", str)
	for _, v := range str {
		//fmt.Printf("char: %d, %s\n", v, string(v))
		dfa.readCharacter(string(v))
	}
}

type DFADesign struct {
	StartState int
	AcceptStates []int
	RuleBook DFARuleBook
}

func (d *DFADesign) toDFA() DFA {
	return DFA{
		CurrentState: d.StartState,
		AcceptStates: d.AcceptStates,
		RuleBook:     d.RuleBook,
	}
}

func (d *DFADesign) accepts(str string) bool {
	dfa := d.toDFA()
	dfa.readString(str)
	return dfa.accepting()
}