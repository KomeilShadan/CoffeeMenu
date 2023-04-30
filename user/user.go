package user

import "coffeeMenu/utils"

type User struct {
	name           string
	age            int
	favoriteNumber float64
	ownsADog       bool
}

func GetUserInfo() User {

	var user User
	user.name 			= GetUserName()
	user.age  			= GetUserAge()
	user.favoriteNumber = GetUserFavNumber()
	user.ownsADog 		= GetUserDogOwnership()
}

func GetUserName() string {
	return utils.ReadStringInput("Enter your name.")
}

func GetUserAge() int {
	return utils.ReadIntInput("Enter your age.")
}

func GetUserFavNumber() float64 {
	return utils.ReadFloatInput("Enter your favorite number.")
}

func GetUserDogOwnership() bool {
	return utils.ReadBoolInput("Do you have a dog?")
}
