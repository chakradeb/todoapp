package models

type Todo struct {
	name string
	description string
	items []*TodoItem
}

func NewTodo(name string, desc string) *Todo {
	return &Todo{name: name, description: desc}
}

func (t *Todo) ChangeName(newName string) {
	t.name = newName
}

func (t *Todo) ChangeDesc(newDesc string) {
	t.description = newDesc
}

func (t *Todo) AddItem(objective string) {
	t.items = append(t.items, NewTodoItem(objective))
}

func (t *Todo) UpdateItem(objective string, index int) {
	t.items[index].changeDescription(objective)
}

func (t *Todo) RemoveItem(index int) {
	t.items = append(t.items[:index], t.items[index+1:]...)
}

func (t *Todo) changeItemStatus(index int) {
	t.items[index].changeStatus()
}
