package doublelinkedlist_test

import (
	"fmt"
	dl "hw4/doublelinkedlist"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	list0, list1        *dl.List
	itemFirst, itemLast *dl.Item
)

func init() {
	list0 = dl.NewList()
	list1 = dl.NewList()
	itemFirst = list1.PushBack("Item 1")
	list1.PushBack("Item 2")
	itemLast = list1.PushBack("Item 3")
}

func TestFirst(t *testing.T) {

	testCases := []struct {
		desc string
		lst  dl.List
		wont interface{}
	}{
		{desc: "Last", lst: *list1, wont: itemLast},
		{desc: "First", lst: *list1, wont: itemFirst},
		{desc: "Len3", lst: *list1, wont: 3},
		{desc: "Len0", lst: *list0, wont: 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			switch tC.desc {
			case "Last":
				assert.Equal(t, tC.wont, tC.lst.Last())
			case "First":
				assert.Equal(t, tC.wont, tC.lst.First())
			case "Len3":
				assert.Equal(t, tC.wont, tC.lst.Len())
			case "Len0":
				assert.Equal(t, tC.wont, tC.lst.Len())
			}
		})
	}
}
func Example() {
	list := dl.NewList()
	list.PushFront("Item PushFront 1")
	list.PushBack("Item PushBack")
	list.PushFront("Item PushFront 2")
	itA := list.PushBack("Item A")
	list.AddValueList("Item B")
	fmt.Println("ShowList:")
	fmt.Print(list.ToString())
	fmt.Println("Count item:", list.Len())
	fmt.Println("Remove Item A")
	fmt.Println("ShowList:")
	list.Remove(*itA)
	fmt.Print(list.ToString())
	fmt.Println("Count item:", list.Len())
	// Output:
	// ShowList:
	// Item PushFront 2
	// Item PushFront 1
	// Item PushBack
	// Item A
	// Item B
	// Count item: 5
	// Remove Item A
	// ShowList:
	// Item PushFront 2
	// Item PushFront 1
	// Item PushBack
	// Item B
	// Count item: 4
}
