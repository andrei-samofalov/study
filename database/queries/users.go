package queries

import (
	"fmt"
	"log"
	conn "study/database/connection"
	"study/database/handlers"
	"study/database/models"
	"study/database/schemas"
)

func AddUser(userData *schemas.UserSchema) *models.UserModel {
	db := conn.DB()
	defer conn.CloseDB(db)

	tx := conn.GetTX(db)

	result := tx.QueryRow(
		`INSERT INTO users ("firstName", "lastName", "age", "balance") 
		VALUES ($1, $2, $3, $4) RETURNING "id";`,
		userData.FirstName, userData.LastName, userData.Age, userData.Balance,
	)

	var newUser models.UserModel
	err := result.Scan(&newUser.UserID)

	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			panic(err2)
		}
	}
	handlers.CheckError(err)
	err = tx.Commit()
	handlers.CheckError(err)
	fmt.Println("newID", newUser.UserID)

	return &newUser
}

func GetUser(userID int) *models.UserModel {

	db := conn.DB()
	defer conn.CloseDB(db)

	tx := conn.GetTX(db)

	var udata models.UserModel
	result := tx.QueryRow(
		`SELECT "firstName", "lastName", "age", "balance" 
			FROM users 
			WHERE id = $1;`,
		userID,
	)

	err := result.Scan(&udata.FirstName, &udata.LastName, &udata.Age, &udata.Balance)

	if err != nil {
		err2 := tx.Rollback() // Откат транзакции в случае ошибки
		if err2 != nil {
			log.Fatal(err2)
		}
		handlers.CheckError(err)
	}
	udata.UserID = userID
	return &udata
}
