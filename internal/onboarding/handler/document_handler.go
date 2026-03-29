package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) UploadDocument(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "file required"})
	}

	employeeID, _ := strconv.Atoi(c.FormValue("employee_id"))
	docCategory := c.FormValue("doc_category")

	err = h.Service.UploadDocument(file, model.UploadDocumentRequest{
		EmployeeID:  employeeID,
		DocCategory: docCategory,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "uploaded"})
}

func (h *OnboardingHandler) GetDocuments(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Query("employee_id"))

	data, err := h.Service.GetDocuments(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (h *OnboardingHandler) DeleteDocument(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.Service.DeleteDocument(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}

func (h *OnboardingHandler) VerifyDocument(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	var body struct {
		Status string `json:"status"`
		Note   string `json:"note"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := h.Service.VerifyDocument(id, body.Status, body.Note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "updated"})
}