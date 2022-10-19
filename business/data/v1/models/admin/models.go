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
	Name         string   `json:"name"`
	Surname      string   `json:"surname"`
	Email        string   `json:"email"`
	PhoneNumber  string   `json:"phoneNumber"`
	Role         string   `json:"role"`
	EnrolledUser []string `json:"enrolledUser"`
	CognitoID    string   `json:"cognitoID"`
	ProfilePick  string   `json:"profilePick"`
	Password     string   `json:"password"`
}
