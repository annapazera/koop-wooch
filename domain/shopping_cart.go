package domain

type ShoppingCart struct {
	Items map[string]CartItem
}

func (s *ShoppingCart) Add(ci CartItem) {
	if value, exists := s.Items[ci.Product.ID]; exists {
		value.Quantity++
		s.Items[ci.Product.ID] = value
	} else {
		s.Items[ci.Product.ID] = ci
	}
}

type CartItem struct {
	Product  Product
	Quantity uint
}

func (c CartItem) Overall() float64 {
	return float64(c.Quantity) * c.Product.Price
}

func (s *ShoppingCart) Sum (Items map[string]CartItem) float64 {
	var sum float64

	for _,item := range s.Items {
	sum += item.Overall()
	}
	return sum
}
