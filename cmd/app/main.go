package main

import (
	"fmt"

	"github.com/Yandex-Practicum/final-project-encoding-go/internal"
	"github.com/Yandex-Practicum/final-project-encoding-go/internal/database"
)

func main() {
	db, err := database.Open()
	internal.CheckOpenedDBError(err)
	defer db.Close()

	//добавление нового клиента
	newClient := internal.NewClient("Stephen Howking", "stiveh", "08.01.1942", "stivenwilh@gmail.com")

	id, err := database.InsertClient(db, newClient)
	internal.HandleError(err)

	// получение клиента по идентификатору и вывод на консоль
	client, err := database.SelectClient(db, id)
	internal.HandleError(err)
	fmt.Println(client)

	// обновление логина клиента
	newLogin := "STEVENWILLH" // укажите новый логин
	err = database.UpdateClientLogin(db, newLogin, id)
	internal.HandleError(err)

	// получение клиента по идентификатору и вывод на консоль
	client, err = database.SelectClient(db, id)
	internal.HandleError(err)
	fmt.Println(client)

	// удаление клиента
	err = database.DeleteClient(db, 251)
	internal.HandleError(err)

	// получение клиента по идентификатору и вывод на консоль
	_, err = database.SelectClient(db, 251)
	internal.HandleError(err)
}
