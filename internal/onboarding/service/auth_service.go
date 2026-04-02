package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"hrms/internal/onboarding/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *OnboardingService) Register(req model.RegisterRequest) error {

	emp, err := s.Repo.GetEmployeeByCode(req.EmployeeCode)
	if err != nil {
		fmt.Printf("DEBUG GetEmployeeByCode error: %v | code used: '%s'\n", err, req.EmployeeCode)
		fmt.Println()
		return fmt.Errorf("employee not found with that employee code")

	}

	if emp.WorkEmail != req.WorkEmail {
		return fmt.Errorf("work email does not match employee code")
	}

	if emp.PasswordHash != "" {
		return fmt.Errorf("employee already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	return s.Repo.SetPasswordHash(emp.ID, string(hash))
}

func (s *OnboardingService) Login(req model.LoginRequest) (*model.AuthResponse, error) {

	emp, err := s.Repo.GetEmployeeByWorkEmail(req.WorkEmail)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if emp.PasswordHash == "" {
		return nil, fmt.Errorf("employee not registered yet")
	}

	err = bcrypt.CompareHashAndPassword([]byte(emp.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	accessToken, err := generateAccessToken(emp.ID, emp.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshHash, expiry, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.Repo.StoreRefreshToken(emp.ID, refreshHash, expiry)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *OnboardingService) Refresh(req model.RefreshRequest) (*model.AuthResponse, error) {

	incomingHash := hashToken(req.RefreshToken)

	emp, expiry, err := s.Repo.GetEmployeeByRefreshHash(incomingHash)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token")
	}

	expiryTime, err := time.Parse("2006-01-02 15:04:05", expiry)
	if err != nil {
		return nil, fmt.Errorf("invalid token expiry")
	}

	if time.Now().After(expiryTime) {
		return nil, fmt.Errorf("refresh token expired")
	}

	accessToken, err := generateAccessToken(emp.ID, emp.Role)
	if err != nil {
		return nil, err
	}

	newRefreshToken, newHash, newExpiry, err := generateRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.Repo.StoreRefreshToken(emp.ID, newHash, newExpiry)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *OnboardingService) Logout(employeeID int) error {
	return s.Repo.ClearRefreshToken(employeeID)
}

func generateAccessToken(employeeID int, role string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"employee_id": employeeID,
		"role":        role,
		"exp":         time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func generateRefreshToken() (plainToken, hash, expiry string, err error) {
	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return
	}
	plainToken = hex.EncodeToString(b)
	hash = hashToken(plainToken)
	expiry = time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02 15:04:05")
	return
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}
