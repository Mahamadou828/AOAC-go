package admin

import "time"

// Admin represents an admin whose has access to the admin application.
// An admin can manage and enrolled user.
// There's two possible role for an admin: ADMIN, SUPER_ADMIN
// ADMIN -> can add/manage users, their applications and can chat with other admin, users and tech support
// SUPER_ADMIN -> has all admin permissions plus he can manage admin themselves.
type Admin struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phoneNumber"`
	Role         string    `json:"role"`
	EnrolledUser []string  `json:"enrolledUser"`
	CognitoID    string    `json:"cognitoID"`
	ProfilePick  string    `json:"profilePick"`
	CreatedAt    time.Time `json:"createdAt"`
	DeleteAt     time.Time `json:"deleteAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// NewAdminDTO represents all data needed to create a new administrator
type NewAdminDTO struct {
	Name        string `json:"name" validate:"required"`
	Surname     string `json:"surname" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Role        string `json:"role" validate:"required"`
	ProfilePick string `json:"profilePick" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// UpdateAdminDTO defines what information may be provided to modify an existing
// Admin. All fields are optional so clients can send just the fields they want
// changed. Since the request have to support file upload we will not use pointer semantics field.
// Instance we will check each field against they zero value.
type UpdateAdminDTO struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email" validate:"omitempty,email"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
	ProfilePick string `json:"profilePick"`
}

// LoginAdminDTO defines all data needed to log an admin
type LoginAdminDTO struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
}

// RefreshTokenDTO defines all data needed to refresh a session
type RefreshTokenDTO struct {
	ID           string `json:"id" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// ConfirmSignupDTO defines all data needed to confirm a new user or admin account
type ConfirmSignupDTO struct {
	Code   string `json:"code" validate:"required"`
	UserID string `json:"userID" validate:"required"`
}
