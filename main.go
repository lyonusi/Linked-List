package main

import (
	"fmt"

	"github.com/lyonusi/Linked-List/list"
)

type testData struct {
	data string
}

func (t *testData) Compare(d list.Data) (bool, error) {
	// Convert input d (a struct implementing list.Data interface) to TestData, "ok" will be false if fields of d not match with fields in TestData
	compareData, ok := d.(*testData)
	if !ok {
		return false, fmt.Errorf("Input cannot be converted to TestData")
	} else {
		return t.data == compareData.data, nil
	}
}

func main() {
	list := list.NewLinkedList()
	list.Add(&testData{"00"})
	list.Add(&testData{"aa"})
	list.Add(&testData{"bb"})
	list.Add(&testData{"cc"})
	list.InsertAfter(-4, &testData{"dd"})
	list.Print()
	fmt.Println(list.Contains(&testData{"cc"}))
	list.Pop()
	list.Print()
	fmt.Println(list.Contains(&testData{"00"}))
	fmt.Println(list.IndexOf(&testData{"cc"}))
	fmt.Println(list.IndexOf(&testData{"00"}))
	list.Print()
	list.Remove(&testData{"ee"})
	list.Push(&testData{"22"})
	list.Push(&testData{"11"})
	list.Push(&testData{"00"})
	list.RemoveByIndex(1)
	list.Print()
	list.Set(6, &testData{"11"})
	list.Set(1, &testData{"11"})
	list.Print()
}
