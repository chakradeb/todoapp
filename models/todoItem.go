package models

type TodoItem struct {
	objective string
	isDone bool
}

func NewTodoItem(obj string) *TodoItem {
	return &TodoItem{objective: obj, isDone: false}
}

func (item *TodoItem) changeDescription(newObj string) {
	item.objective = newObj
}

func (item *TodoItem) changeStatus() {
	item.isDone = !item.isDone
}
