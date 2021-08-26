package main

import (
	"fmt"
	"linkedList/list"
)

func main() {
	list := list.NewLinkedList()
	list.Add("00")
	list.Add("aa")
	list.Add("bb")
	list.Add("cc")
	list.InsertAfter(4, "dd")
	list.InsertAfter(3, "dd")
	list.Print()
	fmt.Println(list.Contains("cc"))
	list.Pop()
	list.Print()
	fmt.Println(list.Contains("00"))
	fmt.Println(list.IndexOf("cc"))
	fmt.Println(list.IndexOf("00"))
	list.Remove("dd")
	list.Print()
	list.Remove("ee")
	list.Push("22")
	list.Push("11")
	list.Push("00")
	list.RemoveByIndex(1)
	list.Print()
	list.Set(6, "11")
	list.Set(1, "11")
	list.Print()
	fmt.Println("Length = ", list.GetLength())
}
