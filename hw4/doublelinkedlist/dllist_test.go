package doublelinkedlist_test

import (
	"fmt"
	l "hw4/doublelinkedlist"
)

func Example() {
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
	fmt.Println("Remove Item A")
	fmt.Println("ShowList:")
	list.Remove(*itA)
	fmt.Print(list.ToString())
	fmt.Println("Count item:", list.Len())
	// Output:
	// Created list
	// ShowList:
	// Item PushFront 2
	// Item PushFront
	// Item PushBack
	// Item A
	// Item C
	// Item B
	// Count item: 6
	// --------------------------
	// Remove Item A
	// ShowList:
	// Item PushFront 2
	// Item PushFront
	// Item PushBack
	// Item C
	// Item B
	// Count item: 5
}
