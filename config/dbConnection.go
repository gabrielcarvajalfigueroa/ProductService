package database

import(

	"fmt"
	"os"

	"github.com/gabrielcarvajalfigueroa/ProductService/models"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var dbUrl = os.Getenv("DB_URL")

	if dbUrl == ""{
		panic("DB_URl enviroment variables is missing")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	autoMigrate(DB)
}

func autoMigrate(connection *gorm.DB) {
	//connection.Debug().AutoMigrate(&models.User{}, &models.Task{}) //Informacion por consola
	connection.AutoMigrate(&models.Product{})
}