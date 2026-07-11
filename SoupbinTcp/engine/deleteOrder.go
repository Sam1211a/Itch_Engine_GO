package engine

func (e *Engine) HandleDelete(orderID uint64) {
	order, ok := e.Book.Get(orderID)
	if !ok {
		return
	}

	if order.Side == 'B' {

		e.RemoveBid(
			order.BookID,
			order.Price,
			order.Qty,
		)

	} else {

		e.RemoveAsk(
			order.BookID,
			order.Price,
			order.Qty,
		)
	}

	e.Book.CancelOrder(orderID)
}
