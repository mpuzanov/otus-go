package doublelinkedlist

// Item Элемент списка
type Item struct {
	next  *Item       // A reference to the next node
	prev  *Item       // A reference to the previous node
	value interface{} // Data or a reference to data
}

//Value возвращает значение
func (i *Item) Value() interface{} {
	return i.value
}

// Next следующий Item
func (i *Item) Next() *Item {
	return i.next
}

//Prev предыдущий
func (i *Item) Prev() *Item {
	return i.prev
}
