package models

import "fmt"

type UserModel struct {
	UserID    int    `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       uint   `json:"Age"`
	Balance   string `json:"Balance"`
}

func (m UserModel) String() string {
	return fmt.Sprintf(
		"User (id %d, name %s %s, age %d, balance %s)",
		m.UserID, m.FirstName, m.LastName, m.Age, m.Balance,
	)
}

func (m UserModel) Val() string {
	return fmt.Sprintf(
		"UserModel (id=%d, firstName=%s, lastName=%s, age=%d, balance=%s)",
		m.UserID, m.FirstName, m.LastName, m.Age, m.Balance,
	)
}
