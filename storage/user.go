package storage

import (
	
	"io/ioutil"
	"encoding/json"
	"os"
	"errors"

	"app/models"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return 0, err
	}

	if len(users) > 0 {
		id = users[len(users)-1].Id + 1
		users = append(users, &models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
		})
	} else {
		id = 1
		users = append(users, &models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
		})
	}

	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepo) GetListRequest(offset, limit int) (*models.GetListResponse, error){

	
	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	
	if err != nil {
		return nil, err
	} 
	return &models.GetListResponse{Users: users,Count: len(users)}, nil
	
}


func (u *userRepo) GetAllUsers() (*models.GetListResponse, error){

	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	
	if err != nil {
		return nil, err
	} 
	return &models.GetListResponse{Users: users,Count: len(users)}, nil

}


func (u *userRepo) DeleteUser(input_id *models.Delete) (error){

	var users []*models.User
	err1 := json.NewDecoder(u.file).Decode(&users)

	if len(users) < input_id.Id{
		return errors.New("Index out of range")
	}

	if err1 != nil{
		return err1
	}

	for in, user := range users{

		if user.Id == input_id.Id{
			

			users = append(users[:in], users[in+1:]...)

			
			body, err := json.MarshalIndent(users, "", "	")
			err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
			
			if err != nil {
				return err
			}

			return nil
		}
	}
	
	return errors.New("Something is wrong")
}


func (u *userRepo) UpdateUser(UpdateUser *models.UpdateUser, id int) error{

	var users []*models.User
	err1 := json.NewDecoder(u.file).Decode(&users)

	if len(users) < id{
		return errors.New("Index out of range")
	}

	if err1 != nil{
		return err1
	}

	for in, user := range users{

		if user.Id == id{
			users[in].Name = UpdateUser.Name
			users[in].Surname = UpdateUser.Surname

			body, err := json.MarshalIndent(users, "", "	")
			err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
			
			if err != nil {
				return err
			}

			return nil
			
		}
	}

	return errors.New("Index out of range")
}