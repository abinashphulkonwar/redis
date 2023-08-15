package storage

import (
	"fmt"
	"reflect"
)

type Node struct {
	Next  *Node
	Prev  *Node
	Value interface{}
}

type List struct {
	head   *Node
	tail   *Node
	Length int
}

func Init() *List {
	return &List{
		head:   nil,
		tail:   nil,
		Length: 0,
	}
}
func (list *List) isPointer(value interface{}) bool {
	return reflect.ValueOf(value).Kind() == reflect.Ptr
}

func (list *List) dataTypeCheck(value interface{}) bool {
	if list.isPointer(value) {
		return false
	}
	switch value.(type) {
	case int:
		return true
	case []byte:
		return true
	case string:
		return true
	default:
		println("Unsupported type")

		return false
	}

}

func (list *List) RPush(value interface{}) {

	if !list.dataTypeCheck(value) {
		return
	}

	node := &Node{
		Next:  nil,
		Prev:  nil,
		Value: value,
	}

	if list.Length == 0 {
		list.tail = node
		list.head = node
		list.Length++

		return
	}
	list.Length++

	tail := list.tail

	tail.Next = node
	node.Prev = tail
	list.tail = node

}
func (list *List) LPush(value interface{}) {

	if !list.dataTypeCheck(value) {
		return
	}
	node := &Node{
		Next:  nil,
		Prev:  nil,
		Value: value,
	}

	if list.Length == 0 {
		list.tail = node
		list.head = node
		list.Length++

		return
	}
	list.Length++

	head := list.head

	head.Prev = node
	node.Next = head
	list.head = node

}

func (list *List) RRemove() {

	if list.Length == 0 {

		return
	}
	list.Length--

	if list.Length == 0 {
		list.head = nil
		list.tail = nil
		return
	}

	tail := list.tail

	tail.Prev.Next = nil
	list.tail = tail.Prev
	tail.Prev = nil

}
func (list *List) LRemove() {

	if list.Length == 0 {

		return
	}
	list.Length--

	if list.Length == 0 {
		list.head = nil
		list.tail = nil
		return
	}

	head := list.head

	head.Next.Prev = nil
	list.head = head.Next
	head.Next = nil

}

func (list *List) Travers() {
	current := list.head
	for i := 0; i < list.Length; i++ {
		switch v := current.Value.(type) {
		case int:
			fmt.Printf("%d\n", v)
		case []byte:
			fmt.Printf("%s\n", v)
		case string:
			fmt.Printf("%s\n", v)
		default:
			fmt.Println("Unsupported type")
		}
		current = current.Next
	}
	current = list.tail
	for i := 0; i < list.Length; i++ {
		switch v := current.Value.(type) {
		case int:
			fmt.Printf("%d\n", v)
		case []byte:
			fmt.Printf("%s\n", v)
		case string:
			fmt.Printf("%s\n", v)
		default:
			fmt.Println("Unsupported type")
		}
		current = current.Prev
	}
}
