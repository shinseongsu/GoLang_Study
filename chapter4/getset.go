package main

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t *Item) Name() string {
	return t.name
}

func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.name = n
	}
}

func (t *Item) Price() float64 {
	return t.price
}

func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.price = p
	}
}

func (t *Item) Quantity() int {
	return t.quantity
}

func (t *Item) SetQuenatity(q int) {
	if q > 0 {
		t.quantity = q
	}
}
