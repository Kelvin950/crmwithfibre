package lead

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/kelvin950/crmwithfibre/database"

_	"github.com/jinzhu/gorm/dialects/mysql"
	
)

type Lead struct {

	gorm.Model 
	Name string  `json:"name"`
	Company string  `json:"company"`
	Email string 	`json:"email"`
	Phone int	`json:"phone"`

}



func GetLeads(c *fiber.Ctx)error{
    
	db := database.DBConn ;

	var leads []Lead 
    
	db.Find(&leads) 
	return c.JSON(leads)
}



func GetLead(c *fiber.Ctx)error{
 id := c.Params("id") 

	db := database.DBConn  
	var lead Lead 
	db.Find(&lead , id) 
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx)error{
 
	db:=  database.DBConn 
	lead :=new(Lead) 

	if err := c.BodyParser(lead); err!=nil{
		c.Status(503).Send([]byte(fmt.Sprintf("err %s", err))) 
		return  err
	}

	db.Create(&lead)
	return c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx)error{
  
	id:= c.Params("id") 
	db:= database.DBConn

	var lead Lead
	db.First(&lead ,id)
	if lead.Name == ""{
		c.Status(500).Send([]byte("No lead found with id"))
	}
     
	 db.Delete(&lead)
	 type message struct{
		message string
	 } 
return 	 c.JSON(message{message:"deleted"})
}