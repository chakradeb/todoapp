package models

type User struct {
	name  string
	todos []*Todo
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (user *User) AddTodo(name string, desc string) {
	user.todos = append(user.todos, NewTodo(name, desc))
}

func (user *User) UpdateTodo(index int, updatedName string, updatedDesc string) {
	user.todos[index].ChangeName(updatedName)
	user.todos[index].ChangeDesc(updatedDesc)
}

func (user *User) RemoveTodo(index int) {
	user.todos = append(user.todos[:index], user.todos[index+1:]...)
}
