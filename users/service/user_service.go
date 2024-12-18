package service

import(
	"user_management/users/models"
	"github.com/sirupsen/logrus"

)

var Users []models.Users
func AddUser(id int, name string){
	user := models.Users{ID: id, Name: name}
	Users= append(Users, user)
	
	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
		"user_name": user.Name,
	}).Info("User added successfully")

}

func GetUserById (id int )models.Users{
	for _, user := range Users{
		if user.ID == id{
			return user
		}
	}
	return models.Users{}
}

func GetAllUsers() []models.Users{
	return Users
}
