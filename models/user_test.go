package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := NewUser("Happy User")
	expectedUser := &User{
		name: "Happy User",
	}

	assert.Equal(t, expectedUser, user, "Invalid user")
}

func TestUser_AddTodo(t *testing.T) {
	user := NewUser("Happy User")
	user.AddTodo("Todo 1", "This is a simple todo")

	todo := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
	}

	expectedUser := &User{
		name:  "Happy User",
		todos: []*Todo{todo},
	}
	assert.Equal(t, expectedUser, user, "Invalid user todo")
}

func TestUser_UpdateTodo(t *testing.T) {
	user := NewUser("Happy User")
	user.AddTodo("Todo 1", "This is a simple todo")
	user.UpdateTodo(0, "Updated name", "This is a updated description")

	todo := &Todo{
		name:        "Updated name",
		description: "This is a updated description",
	}

	expectedUser := &User{
		name:  "Happy User",
		todos: []*Todo{todo},
	}
	assert.Equal(t, expectedUser, user, "Invalid user todo")
}

func TestUser_RemoveTodo(t *testing.T) {
	user := NewUser("Happy User")
	user.AddTodo("Todo 1", "This is a simple todo")
	user.AddTodo("Todo 2", "This is a simple todo")
	user.AddTodo("Todo 3", "This is a simple todo")
	user.RemoveTodo(1)

	todo1 := &Todo{
		name:        "Todo 1",
		description: "This is a simple todo",
	}
	todo2 := &Todo{
		name:        "Todo 3",
		description: "This is a simple todo",
	}

	expectedUser := &User{
		name:  "Happy User",
		todos: []*Todo{todo1, todo2},
	}
	assert.Equal(t, expectedUser, user, "Invalid user todo")
}
