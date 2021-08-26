package main

import "fmt"

type node struct {
	data string
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

type LinkedList interface {
	getLength() int
	insertAfter(index int, data string) error
	pop() error
	add(data string)
	push(data string)
	contains(data string) bool
	indexOf(data string) int
	remove(data string) error
	removeByIndex(index int) error
	set(index int, data string) error
	print()
}

func (l *linkedList) print() {
	fmt.Printf("...Printing...\n")
	pointer := l.head
	for pointer != nil {
		fmt.Printf("%v -> ", pointer.data)
		pointer = pointer.next
	}
	fmt.Printf("(null)\n")
}

func (l *linkedList) add(data string) {
	fmt.Printf("...Adding \"%v\" to the list...", data)
	node := &node{
		data: data,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length++
	fmt.Println("[Success!]")
}

func (l *linkedList) getLength() (len int) {
	return l.length
}

func (l *linkedList) push(data string) {
	fmt.Printf("...Pushing \"%v\" to the list...", data)
	node := &node{
		data: data,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		currentHead := l.head
		l.head = node
		l.head.next = currentHead
	}
	l.length++
	fmt.Println("[Success!]")
}

func (l *linkedList) insertAfter(index int, data string) error {
	fmt.Printf("...Inserting \"%v\" after node(%v) to the list...", data, index)
	if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", l.length-1)
		fmt.Println(err)
		return err
	}
	node := &node{
		data: data,
	}
	pointer := l.head
	for i := 0; i <= index; i++ {
		if i == index {
			node.next = pointer.next
			pointer.next = node
			l.length++
			fmt.Println("[Success!]")
		} else {
			pointer = pointer.next
		}
	}

	return nil
}

func (l *linkedList) pop() error {
	fmt.Printf("...Poping the first node...")
	if l.head == nil {
		err := fmt.Errorf("[error: empty list]\n")
		return err
	}
	data := l.head.data
	fmt.Printf("\"%v\"...", data)
	l.head = l.head.next
	l.length--
	fmt.Println("[Success!]")
	return nil
}

func (l *linkedList) contains(data string) bool {
	fmt.Printf("...Checking if the list contains \"%v\"...\n", data)
	pointer := l.head
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			return true
		} else {
			pointer = pointer.next
		}
	}
	return false
}

func (l *linkedList) indexOf(data string) int {
	fmt.Printf("...Searching \"%v\" in the list...\n", data)
	pointer := l.head
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			return i
		} else {
			pointer = pointer.next
		}
	}
	return -1
}
func (l *linkedList) remove(data string) error {
	fmt.Printf("...Removing \"%v\" from the list...", data)
	pointer := l.head
	var prevPointer *node
	for i := 0; i < l.length; i++ {
		if pointer.data == data {
			prevPointer.next = pointer.next
			l.length--
			fmt.Println("[Success!]")
			return nil
		} else {
			prevPointer = pointer
			pointer = pointer.next
		}
	}
	err := fmt.Errorf("[error: \"%v\" not found]", data)
	fmt.Println(err)
	return err
}
func (l *linkedList) removeByIndex(index int) error {
	fmt.Printf("...Removing node(%v) from the list...", index)
	if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", index)
		fmt.Println(err)
		return err
	}
	pointer := l.head
	var prevPointer *node
	for i := 0; i < l.length; i++ {
		if i == index {
			prevPointer.next = pointer.next
			l.length--
			fmt.Println("[Success!]")
			return nil
		} else {
			prevPointer = pointer
			pointer = pointer.next
		}
	}

	return nil
}

func (l *linkedList) set(index int, data string) error {
	fmt.Printf("...Setting node(%v) as \"%v\"...", index, data)
	if index >= l.length {
		err := fmt.Errorf("[invalid index: exceeded maximum index (%v)]", l.length-1)
		fmt.Println(err)
		return err
	}
	pointer := l.head
	for i := 0; i < l.length; i++ {
		if i == index {
			pointer.data = data
			fmt.Println("[Success!]")
			return nil
		} else {
			pointer = pointer.next
		}
	}
	return nil
}

func main() {
	list := linkedList{}
	list.add("00")
	list.add("aa")
	list.add("bb")
	list.add("cc")
	list.insertAfter(4, "dd")
	list.insertAfter(3, "dd")
	list.print()
	fmt.Println(list.contains("cc"))
	list.pop()
	list.print()
	fmt.Println(list.contains("00"))
	fmt.Println(list.indexOf("cc"))
	fmt.Println(list.indexOf("00"))
	list.remove("dd")
	list.print()
	list.remove("ee")
	list.push("22")
	list.push("11")
	list.push("00")
	list.removeByIndex(1)
	list.print()
	list.set(6, "11")
	list.set(1, "11")
	list.print()
	fmt.Println("Length = ", list.getLength())
}