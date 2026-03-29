package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"hrms/internal/onboarding/model"
	
)


func (s *OnboardingService) UploadDocument(file *multipart.FileHeader, req model.UploadDocumentRequest) error {

	// validation
	if file.Size > 5*1024*1024 {
		return fmt.Errorf("file too large")
	}

	ext := filepath.Ext(file.Filename)
	path := fmt.Sprintf("uploads/%d/%s%s", req.EmployeeID, req.DocCategory, ext)

	url, err := s.Storage.Upload(file, path)
	if err != nil {
		return err
	}

	doc := model.EmployeeDocument{
		EmployeeID:  req.EmployeeID,
		DocCategory: req.DocCategory,
		FileName:    file.Filename,
		S3URL:       url,
		FileSizeKB:  int(file.Size / 1024),
		MimeType:    file.Header.Get("Content-Type"),
	}

	return s.Repo.SaveDocument(doc)
}

func (s *OnboardingService) GetDocuments(employeeID int) ([]model.EmployeeDocument, error) {
	return s.Repo.GetDocuments(employeeID)
}

func (s *OnboardingService) DeleteDocument(id int) error {
	return s.Repo.DeleteDocument(id)
}

func (s *OnboardingService) VerifyDocument(id int, status string, note string) error {
	return s.Repo.VerifyDocument(id, status, note)
}