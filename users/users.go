package users
import (
	"user_management/users/models"
	"user_management/users/service"
)

func AddUser(id int, name string){
	service.AddUser(id, name)
}

func GetUserById(id int) models.Users{
	return service.GetUserById(id)
}

func GetAllUsers() []models.Users{	
	return service.GetAllUsers()
}	