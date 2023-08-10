package main

import (
	"fmt"
)

func FuncsStart() {
	fmt.Println()
	fmt.Println(MergeNumberLists([]int{10, 20}, nil, []int{50, 60}))
	makeCopy()

	c := Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	fmt.Println(c.Value) // 42

	c.Dec(0)
	c.Dec(30)
	fmt.Println(c.Value) // 11

	c.Dec(100)
	fmt.Println(c.Value) // 0

	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}
	fmt.Println(pl.GetAgePopularity())
}

func MergeNumberLists(numberLists ...[]int) []int {
	result := make([]int, 0)
	for _, list := range numberLists {
		result = append(result, list...)
	}
	return result
}

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	par := &Parent{}
	par.Name = p.Name
	par.Children = append(par.Children, p.Children...)
	return *par
}

func makeCopy() {
	cp := CopyParent(nil)

	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}

	cp = CopyParent(p)

	cp.Children[0] = Child{}
	fmt.Println(p.Children)
}

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		c.Value += 1
	}
	c.Value += delta
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		c.Value -= 1
	}
	c.Value -= delta
	if c.Value < 0 {
		c.Value = 0
	}
}

type Person struct {
	Age uint8
}

type PersonList []Person 

func (pl PersonList) GetAgePopularity() map[uint8]int {
	finMap := map[uint8]int{}
	for _, pers := range pl {
		finMap[pers.Age]++
	}
	return finMap
}