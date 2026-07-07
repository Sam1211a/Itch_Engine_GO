package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleDelete(order *model.DeleteOrder) {
	e.Book.CancelOrder(order.OrderId)
	fmt.Println("Engine -> DeleteOrder")
}
