package models

type Product struct {
    BaseModel
    Name  string
    Price Price
    Stock uint8
}