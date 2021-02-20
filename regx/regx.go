package regx

import (
	"fmt"
	"strings"
)

type ToStringer interface {
	String() string
	Precedence() int
	Bracket(int) string
}

//func (p Pattern) String() string {
//	return fmt.Sprintf("/%s/", p.String())
//}

type Empty struct {
}

func (e Empty) Bracket(outerPrecedence int) string {
	if e.Precedence() < outerPrecedence {
		return fmt.Sprintf("(%s)", e)
	}
	return fmt.Sprintf("%s", e)
}

func (e Empty) Precedence() int {
	return 3
}

func (e Empty) String() string {
	return ""
}

type Literal struct {
	Character string
}

func (l Literal) Bracket(outerPrecedence int) string {
	if l.Precedence() < outerPrecedence {
		return fmt.Sprintf("(%s)", l)
	}
	return fmt.Sprintf("%s", l)
}

func (l Literal) Precedence() int {
	return 3
}

func (l Literal) String() string {
	return l.Character
}

type Concatenate struct {
	First, Second ToStringer
}

func (c Concatenate) Bracket(outerPrecedence int) string {
	if c.Precedence() < outerPrecedence {
		return fmt.Sprintf("(%s)", c)
	}
	return fmt.Sprintf("%s", c)
}

func (c Concatenate) Precedence() int {
	return 1
}

func (c Concatenate) String() string {
	return c.First.Bracket(c.First.Precedence()) + c.Second.Bracket(c.Second.Precedence())
}

type Choose struct {
	First, Second ToStringer
}

func (c Choose) Bracket(outerPrecedence int) string {
	if c.Precedence() < outerPrecedence {
		return fmt.Sprintf("(%s)", c)
	}
	return fmt.Sprintf("%s", c)
}

func (c Choose) Precedence() int {
	return 0
}

func (c Choose) String() string {
	return strings.Join([]string{
		c.First.Bracket(c.First.Precedence()),
		c.Second.Bracket(c.Second.Precedence()),
	}, "|")
}

type Repeat struct {
	Pattern ToStringer
}

func (r Repeat) Bracket(outerPrecedence int) string {
	if r.Precedence() < outerPrecedence {
		return fmt.Sprintf("(%s)", r)
	}
	return fmt.Sprintf("%s", r)
}

func (r Repeat) Precedence() int {
	return 2
}

func (r Repeat) String() string {
	return r.Pattern.Bracket(r.Precedence()) + "*"
}