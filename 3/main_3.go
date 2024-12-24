package main

import (
	"fmt"
	"go_core_task/3/stringIntMap"
)

func main() {
	m := stringIntMap.New()

	m.Add("one", 1)
	m.Add("two", 2)

	fmt.Print("1. ")
	m.Print()

	m.Remove("two")
	fmt.Print("2. ")
	m.Print()

	tmp := m.Copy()
	fmt.Println("3. ", tmp)

	fmt.Println("4. ", m.Exists("one"))

	v, ok := m.Get("one")
	fmt.Println("5. ", v, ok)

}
