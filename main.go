package main 

import {
	"github.com/gofiber/fiber/v2"
	"github.com/kelvin950/crmwithfibre/databse"
}



 func setupRoutes(app *fiber.App){

	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDb(){

}

func main(){

	 app :=  fiber.New()

	 setupRoutes(app)

	 app.Listen(":3000")
    
}