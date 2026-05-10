package domain

import "time"

type Product struct {
    Base
    Name  string
    Price Price
    Stock uint8
}

func NewProduct(name string, price float64, stock int) *Product {
    return &Product{
        Base: Base{
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        Name:  name,
        Price: Price(price),
        Stock: uint8(stock),
    }
}