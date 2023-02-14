package main

import (
	"fmt"
	"log"

	// "log"
	// "os/user"

	"app/config"
	"app/controller"
	"app/models"

	// "app/models"
	"app/storage"
)

func main() {

	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)

	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name:    "Abduqodir",
	// 		Surname: "Musayev",
	// 	},
	// )

	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// fmt.Println(id)
	
	// users, err := c.GetListRequest(1,10)

	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Println(users)
	// }
	
	// user, err := c.GetByKey(5)

	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(user)

	// err = c.DeleteUser(4)

	// if err != nil{
	// 	log.Println(err)
	// 	return
	// } else {
	// 	fmt.Println("User has been deleted")
	// }

	// update := &models.UpdateUser{Name: "Ali", Surname: "Ismoilov"}
	
	err = c.UpdateUser(&models.UpdateUser{Name: "Ali", Surname: "Ismoilov"},10)

	if err != nil{
		log.Println(err)
		return
	} else {
		fmt.Println("User has been updated")
	}
}