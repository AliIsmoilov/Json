package controller

import (
	"errors"

	"app/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func(c *Controller) GetListRequest(offset, limit int) (*models.GetListResponse, error){

	users, err := c.store.User.GetListRequest(offset, limit)

	if err != nil{
		return nil, err
	} else{
		return users, nil
	}
}

func (c *Controller) GetByKey(id int) (*models.GetByKey, error){

	users, err := c.store.User.GetAllUsers()

	if err != nil{
		return nil, err
	}

	for _, user := range users.Users{
		if user.Id == id{
			
			found_user := *&models.GetByKey{Name: user.Name, Surname: user.Surname}

			return &found_user, nil
		}
	}

	return nil, errors.New("Not found such user")
}


func (c *Controller) DeleteUser(id int) (error){

	err := c.store.User.DeleteUser(&models.Delete{Id: id})

	return err
	
}

func (c *Controller) UpdateUser(update *models.UpdateUser, id int) error{

	err := c.store.User.UpdateUser(&models.UpdateUser{Name: update.Name, Surname: update.Surname}, id)

	return err
	
}
