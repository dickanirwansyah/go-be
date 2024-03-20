package controller

import (
	"log"

	"github.com/dickanirwansyah/blogspot/database"
	"github.com/dickanirwansyah/blogspot/model"
	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	return c.Status(200).JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blogspot successfully !",
	}

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request !")
		context["statusText"] = "Failed"
		context["msg"] = "Something went wrong !"
		return c.Status(500).JSON(context)
	}

	result := database.DBConn.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data !")
		context["statusText"] = "Failed"
		context["msg"] = "Something went wrong !"
		return c.Status(500).JSON(context)
	}

	context["msg"] = "Record is saved successfully !"
	context["data"] = record

	return c.Status(200).JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update a Blogspot successfully !",
	}

	//http://localhost:8000/2

	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found !")
		context["statusText"] = "Failed"
		context["msg"] = "Sorry record not found !"
		return c.Status(404).JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request !")
		context["statusText"] = "Failed"
		context["msg"] = "Error in parsing request !"
		return c.Status(500).JSON(context)
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data !")
		context["statusText"] = "Failed"
		context["msg"] = "Something went wrong !"
		return c.Status(500).JSON(context)
	}

	context["data"] = record
	return c.Status(200).JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete a Blogspot successfully !",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found !")
		context["statusText"] = "Failed"
		context["msg"] = "Record not found !"
		return c.Status(404).JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		log.Println("Error in delete data !")
		context["statusText"] = "Failed"
		context["msg"] = "Something went wrong !"
		return c.Status(500).JSON(context)
	}

	return c.Status(200).JSON(context)
}
