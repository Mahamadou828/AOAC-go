package user

import "time"

const APPLICATION_OPEN_STATUS = "open"

// User represent the student that apply in the agencies to obtain a scholarships.
// Users are created by admin and linked by them the enrolledBy key
type User struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Town           string    `json:"town"`
	Country        string    `json:"country"`
	PhoneNumber    string    `json:"phoneNumber"`
	Birthday       string    `json:"birthday"`
	University     string    `json:"university"`
	GraduationDate string    `json:"graduationDate"`
	Section        string    `json:"section"`
	EnrolledBy     string    `json:"enrolledBy"`
	CognitoID      string    `json:"cognitoID"`
	CreatedAt      time.Time `json:"createdAt"`
	DeleteAt       time.Time `json:"deleteAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// Application represent a user ( student ) application to a single university.
// A student can apply to a maximum of 5 university.
type Application struct {
	ID             string    `json:"id"`
	UserID         string    `json:"userID"`
	UniversityName string    `json:"universityName"`
	EmailContact   string    `json:"emailContact"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	DeleteAt       time.Time `json:"deleteAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// Document represent a document send by a user.
type Document struct {
	ID        string    `json:"id"`
	S3URL     string    `json:"s3URL"`
	Name      string    `json:"name"`
	UserID    string    `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	DeleteAt  time.Time `json:"deleteAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUserDTO represent the data needed to create a new user.
type NewUserDTO struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Town           string `json:"town"`
	Country        string `json:"country"`
	PhoneNumber    string `json:"phoneNumber"`
	Birthday       string `json:"birthday"`
	University     string `json:"university"`
	GraduationDate string `json:"graduationDate"`
	Section        string `json:"section"`
	//EnrolledBy is the admin ID that has created the user and his application
	EnrolledBy string `json:"enrolledBy"`
	//list of universities id selected by the applicant.
	//the limit of university is 5
	SelectedUniversities []string `json:"selectedUniversities"`
	//all those fields are document needed to create an application
	ProfilePick              string `json:"profilePick"`
	NoteCertificate          string `json:"noteCertificate"`
	BaccalaureateCertificate string `json:"baccalaureateCertificate"`
}
