package models

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUserManagement_AddUser(t *testing.T) {
	users := &UserManagement{}
	users.addUser("Happy User")

	user := &User{
		name:"Happy User",
	}

	expectedUsers := &UserManagement{
		users:[]*User{user},
	}

	assert.Equal(t, expectedUsers, users, "Invalid users")
}
