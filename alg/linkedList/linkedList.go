package linkedList

type Element struct {
	next *Element
	// The list to which this element belongs.
	list *linkedList
	// The value stored with this element.
	Value interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type linkedList struct {
	root Element
	len  int
}

func (l *linkedList) init() *linkedList {
	l.root.next = &l.root
	l.len = 0
	return l
}

// insert inserts e after at, increments l.len, and returns e.
func (l *linkedList) insert(e, at *Element) *Element {
	e.next = at.next
	at.next = e
	e.list = l
	l.len++
	return e
}

func New() *linkedList { return new(linkedList).init() }

func main() {

}
