package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/kelvin950/crmwithfibre/database"
	"github.com/kelvin950/crmwithfibre/lead"
	
_	"github.com/jinzhu/gorm/dialects/mysql"
)



 func setupRoutes(app *fiber.App){

	app.Get("/api/v1/lead" ,lead.GetLeads)
	app.Get("api/v1/lead/:id" , lead.GetLead)
	app.Post("api/v1/lead" , lead.NewLead)
	app.Delete("api/v1/lead/:id" ,lead.DeleteLead)
}

func initDb(){
   var err error 

   database.DBConn ,err =  gorm.Open("mysql" , "myuser:mypassword@tcp(localhost:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
   if err !=nil{
	panic(err)
} 

fmt.Println("connection opened to database")
database.DBConn.AutoMigrate(&lead.Lead{})
fmt.Println("Database Migrated")


}

func main(){

	 app :=  fiber.New()

	 initDb()
	 
	 defer database.DBConn.Close()
	 setupRoutes(app)

	 app.Listen(":3000")
	 
    
}