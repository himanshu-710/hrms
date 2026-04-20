package handler

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "hrms/internal/onboarding/model"
)

func (h *OnboardingHandler) UploadDocument(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "file required"})
    }
    docCategory := c.FormValue("doc_category")
    err = h.Service.UploadDocument(file, model.UploadDocumentRequest{
        EmployeeID:  id,
        DocCategory: docCategory,
    })
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "uploaded"})
}

func (h *OnboardingHandler) GetDocuments(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("employeeId"))  
    if err != nil || id == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "invalid employee id"})
    }
    data, err := h.Service.GetDocuments(id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(data)
}

func (h *OnboardingHandler) DeleteDocument(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    if err := h.Service.DeleteDocument(id); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "deleted"})
}

func (h *OnboardingHandler) VerifyDocument(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    var req model.VerifyDocumentRequest  
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    if err := h.Service.VerifyDocument(id, req.Status, req.Note); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "updated"})
}

// Add to document_handler.go
func (h *OnboardingHandler) GetPendingDocuments(c *fiber.Ctx) error {
    data, err := h.Service.GetPendingDocuments()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(data)
}