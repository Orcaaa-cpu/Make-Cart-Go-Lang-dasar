package main

import "fmt"

func main() {
	cart := newCart()
	cart.add("kucing")
	cart.add("babi")
	cart.add("lancek")
	cart.remove("babi")
	cart.show()

}

type cart struct {
	items *[]string
}

func newCart() cart {
	return cart{
		items: &[]string{},
	}
}

func (c cart) add(item string) {
	*c.items = append(*c.items, item)
	fmt.Println("Item succsesfully added")
}

func (c cart) remove(item string) {
	for i, v := range *c.items {
		if item == v {
			tmp := make([]string, len(*c.items))
			copy(tmp, *c.items)
			*c.items = append(append([]string{}, tmp[:i]...), tmp[i+1:]...)
		}
	}
}

func (c cart) show() {
	fmt.Println(c.items)
}
