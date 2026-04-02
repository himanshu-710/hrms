package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"hrms/internal/onboarding/model"
)

func (s *OnboardingService) UploadDocument(file *multipart.FileHeader, req model.UploadDocumentRequest) error {

	if file.Size > 5*1024*1024 {
		return fmt.Errorf("file too large")
	}

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%s_%d%s", req.DocCategory, req.EmployeeID, ext)

	
	objectPath := fmt.Sprintf("uploads/%d/%s", req.EmployeeID, fileName)

	
	storedPath, err := s.Storage.Upload(file, objectPath)
	if err != nil {
		return err
	}

	doc := model.EmployeeDocument{
		EmployeeID:  req.EmployeeID,
		DocCategory: req.DocCategory,
		FileName:    file.Filename,
		S3URL:       storedPath, 
		FileSizeKB:  int(file.Size / 1024),
		MimeType:    file.Header.Get("Content-Type"),
	}

	return s.Repo.SaveDocument(doc)
}


func (s *OnboardingService) GetDocuments(employeeID int) ([]model.EmployeeDocument, error) {

	docs, err := s.Repo.GetDocuments(employeeID)
	if err != nil {
		return nil, err
	}

	for i := range docs {
    var presignedURL string
    var err error
    presignedURL, err = s.Storage.GetPresignedURL(docs[i].S3URL, 1*time.Hour)
    if err != nil {
        fmt.Printf("warning: could not generate presigned URL for doc %d: %v\n", docs[i].ID, err)
        continue
    }
    docs[i].PresignedURL = presignedURL
}

	return docs, nil
}

func (s *OnboardingService) DeleteDocument(id int) error {
	return s.Repo.DeleteDocument(id)
}

func (s *OnboardingService) VerifyDocument(id int, status string, note string) error {
	return s.Repo.VerifyDocument(id, status, note)
}