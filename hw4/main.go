package main

import (
	"fmt"
	l "hw4/doublelinkedlist"
)

func main() {
	fmt.Println("Реализовываем двусвязный список")

	list := l.NewList()
	fmt.Println("Created list")
	list.PushFront("Item PushFront")
	list.PushBack("Item PushBack")
	list.PushFront("Item PushFront 2")
	itA := list.PushBack("Item A")
	list.AddValueList("Item C")
	list.AddValueList("Item B")
	fmt.Println("ShowList:")
	fmt.Print(list.ToString())
	fmt.Println("Count item:", list.Len())
	fmt.Println("--------------------------")
	list.Remove(*itA)
	fmt.Print(list.ToString())
	fmt.Println("Count item:", list.Len())
}
