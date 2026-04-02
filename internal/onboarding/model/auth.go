package model

type RegisterRequest struct {
	WorkEmail    string `json:"work_email"     validate:"required,email"`
	Password     string `json:"password"       validate:"required,min=8,max=72"`
	EmployeeCode string `json:"employee_code"  validate:"required"`
}

type LoginRequest struct {
	WorkEmail string `json:"work_email" validate:"required,email"`
	Password  string `json:"password"   validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type Claims struct {
	EmployeeID int    `json:"employee_id"`
	Role       string `json:"role"`
}

type AuthEmployee struct {
	ID           int
	WorkEmail    string
	EmployeeCode string
	PasswordHash string
	Role         string
}