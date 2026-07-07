package engine

import (
	"soupbintcp/model"
	"soupbintcp/orderbook"
)

type Engine struct {
	Book       *orderbook.Orderbook
	Securities map[uint32]model.SecurityDirectory
}

func New() *Engine {
	return &Engine{
		Book:       orderbook.New(),
		Securities: make(map[uint32]model.SecurityDirectory),
	}
}

func (e *Engine) PrintBook() {
	e.Book.Print()
}
