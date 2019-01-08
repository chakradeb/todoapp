package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
	}

	assert.Equal(t, expectedTodo, todo, "Invalid todo")
}

func TestTodo_ChangeName(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.ChangeName("New Todo")

	expectedTodo := &Todo{
		name:        "New Todo",
		description: "This is a simple todo",
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo name")
}

func TestTodo_ChangeDescription(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.ChangeDesc("New Description")

	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "New Description",
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo name")
}

func TestTodo_AddItem(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.AddItem("Todo one")

	todoItem := &TodoItem{
		objective: "Todo one",
	}

	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
		items:       []*TodoItem{todoItem},
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo item")
}

func TestTodo_UpdateItem(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.AddItem("Todo one")
	todo.UpdateItem("Updated Item", 0)

	todoItem := &TodoItem{
		objective: "Updated Item",
	}

	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
		items:       []*TodoItem{todoItem},
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo item")
}

func TestTodo_RemoveItem(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.AddItem("Todo one")
	todo.AddItem("Todo two")
	todo.AddItem("Todo three")
	todo.RemoveItem(1)

	todoItem1 := &TodoItem{
		objective: "Todo one",
	}
	todoItem2 := &TodoItem{
		objective: "Todo three",
	}

	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
		items:       []*TodoItem{todoItem1, todoItem2},
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo item")
}

func TestTodo_ChangeItemStatus(t *testing.T) {
	todo := NewTodo("Todo 1", "This is a simple todo")
	todo.AddItem("Todo one")
	todo.changeItemStatus(0)

	todoItem1 := &TodoItem{
		objective: "Todo one",
		isDone:    true,
	}

	expectedTodo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
		items:       []*TodoItem{todoItem1},
	}
	assert.Equal(t, expectedTodo, todo, "Invalid todo item")
}
