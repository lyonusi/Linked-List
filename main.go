package main

import (
	"fmt"
	"linkedList/list"
)

func main() {
	list := list.NewLinkedList()
	fmt.Println("===============================================================")
	fmt.Println("===============================================================")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	list.Add(123)
	list.Add("aa")
	list.Add("bb")
	list.Add("cc")
	// list.InsertAfter(-4, "dd")
	// list.Print()
	// fmt.Println(list.Contains("cc"))
	// list.Pop()
	// list.Print()
	// fmt.Println(list.Contains("00"))
	// fmt.Println(list.IndexOf("cc"))
	// fmt.Println(list.IndexOf("00"))
	// list.Print()
	// list.Remove("ee")
	// list.Push("22")
	// list.Push("11")
	// list.Push("00")
	// list.RemoveByIndex(1)
	// list.Print()
	// list.Set(6, "11")

	type test struct {
		num  int
		text string
	}

	t := test{
		num:  3,
		text: "three"}

	list.Set(1, t)
	fmt.Println(list.IndexOf(test{num: 3, text: "three"}))
	list.Print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("===============================================================")
	fmt.Println("===============================================================")
}
