package lead

import (
	"github.com/ImArnav19/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Age     int    `json:"age"`
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)

}
func GetLeads(c *fiber.Ctx) {

	var leads []Lead
	db := database.DBConn

	db.Find(&leads)
	c.JSON(leads)

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.Find(&lead, id)

	if lead.Name == "" {
		c.Status(500).Send("No File named so!!")
		return
	}
	db.Delete(&lead)
	c.Send("File deletion success!")

}

func NewLead(c *fiber.Ctx) {

	db := database.DBConn

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)

}
