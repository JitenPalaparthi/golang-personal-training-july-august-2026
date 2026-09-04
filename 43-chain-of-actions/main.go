package main

import "fmt"

func main() {

	c := NewCalc(100).Add(100).Sub(100).Mul(2).Div(2).Add(100).Get()

	fmt.Println(c)

}

type ICalc interface {
	Add(int) ICalc
	Sub(int) ICalc
	Mul(int) ICalc
	Div(int) ICalc
	Get() int
}

type Calc struct {
	Data int
}

func NewCalc(v int) ICalc {
	return &Calc{v}
}

func (c *Calc) Add(n int) ICalc {
	c.Data += n
	return c
}

func (c *Calc) Sub(n int) ICalc {
	c.Data -= n
	return c
}
func (c *Calc) Mul(n int) ICalc {
	c.Data *= n
	return c
}
func (c *Calc) Div(n int) ICalc {
	c.Data /= n
	return c
}
func (c *Calc) Get() int {
	return c.Data
}

// *Calc --> Concrete Type, ICalc --> Interface Type
// chain of actions
// builder pattern
// fluent api
