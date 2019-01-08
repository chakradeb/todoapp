package models

type UserManagement struct {
	users []*User
}

func (mgmt *UserManagement) addUser(name string) {
	mgmt.users = append(mgmt.users, NewUser(name))
}
