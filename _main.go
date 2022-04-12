// package main

import (
	user "go-api/User"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection to database
	dsn := "root:@tcp(127.0.0.1:3306)/goapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Mahardhika",
	}
	userInput := user.RegisterUserInput{}
	userInput.Name = "tes simpan dari service"
	userInput.Email = "testing@example.com"
	userInput.Occupation = "direktur"
	userInput.Password = "password"
	userService.RegisterUserInput(userInput)
	userRepository.Save(user)

	// fmt.Println("connected to database")

	// var users []user.User
	// db.Find(&users)
	// for _, user := range users {
	// 	fmt.Println(user.Email)
	// }

	// create an API with GIN
	// route := gin.Default()
	// route.GET("/users", handler)
	// route.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/goapi?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
