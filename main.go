package main

import (
	"fmt"
	"user_management/users"
	"os"
	"github.com/sirupsen/logrus"
)

//AppLogger is a globally accessible logger
var AppLogger *logrus.Logger

func init() {
AppLogger = logrus.New()
AppLogger.SetFormatter(&logrus.JSONFormatter{}) // Set the log format to JSON
AppLogger.SetOutput(os.Stdout) // Set the output to stdout
AppLogger.SetLevel(logrus.DebugLevel) // Set the log level to debug
}

func main(){

	AppLogger.Info("Starting the application...")

	users.AddUser(1, "John")
	users.AddUser(2, "Doe")
	
	fmt.Println(users.GetAllUsers())
	fmt.Println(users.GetUserById(2))
}