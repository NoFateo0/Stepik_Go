package main

import "fmt"

type myStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (m *myStruct) Shoot() bool {
	if m.On == false {
		return false
	}
	if m.Ammo > 0 {
		m.Ammo -= 1
	} else {
		return false
	}
	return true
}

func (m *myStruct) RideBike() bool {
	if m.On == false {
		return false
	}
	if m.Power > 0 {
		m.Power -= 1
	} else {
		return false
	}
	return true
}

func main() {
	testStruct := &myStruct{true, 0, 0}
	a := testStruct.Shoot()
	b := testStruct.RideBike()
	fmt.Println(*testStruct, a, b)
}
