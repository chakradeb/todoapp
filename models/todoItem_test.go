package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoItem(t *testing.T) {
	item := NewTodoItem("This is my job")
	expectedItem := &TodoItem{
		objective: "This is my job",
	}

	assert.Equal(t, expectedItem, item, "Invalid item")
}

func TestTodoItem_ChangeDescription(t *testing.T) {
	item := NewTodoItem("This is my job")
	item.changeDescription("Job has been changed")

	expectedItem := &TodoItem{
		objective: "Job has been changed",
		isDone: false,
	}
	assert.Equal(t, expectedItem, item, "Invalid todo description")
}

func TestTodoItem_ChangeStatus(t *testing.T) {
	item := NewTodoItem("This is my job")
	item.changeStatus()

	expectedItem := &TodoItem{
		objective: "This is my job",
		isDone: true,
	}
	assert.Equal(t, expectedItem, item, "Invalid todo status")
}
