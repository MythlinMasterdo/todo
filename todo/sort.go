package todo

type ByOrder []Todo

func (todos ByOrder) Len() int {
	return len(todos)
}

func (todos ByOrder) Swap(i, j int) {
	todos[i], todos[j] = todos[j], todos[i]
}

func (todos ByOrder) Less(i, j int) bool {
	todo1, todo2 := todos[i], todos[j]
	orders1, orders2 := todo1.GetOrders(), todo2.GetOrders()

	var maxDepth int
	if len(orders1) > len(orders2) {
		maxDepth = len(orders1)
	} else {
		maxDepth = len(orders2)
	}

	for depth := 0; depth < maxDepth; depth++ {
		if depth > len(orders1)-1 {
			return true
		}

		if depth > len(orders2)-1 {
			return false
		}

		order1, order2 := orders1[depth], orders2[depth]
		if order1 == order2 {
			continue
		}

		return order1 < order2
	}

	return true
}
