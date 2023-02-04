package user

import (
	"context"
	"fmt"
	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"strings"
	"time"

	uniModel "github.com/Mahamadou828/AOAC/business/data/v1/models/university"

	model "github.com/Mahamadou828/AOAC/business/data/v1/models/user"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func Create(ctx context.Context, cfg *lambda.Config, r events.APIGatewayProxyRequest, now time.Time) (model.User, error) {
	id := validate.GenerateID()

	var nu model.NewUserDTO
	var as []model.Application

	if err := config.ParseForm(&nu, r); err != nil {
		return model.User{}, fmt.Errorf("failed to parse form: %v", err)
	}
	if err := validate.Check(nu); err != nil {
		return model.User{}, fmt.Errorf("invalid user: %v", err)
	}

	//Get all the universities he applied and format the applications
	//@todo to-refacto to do a single request to get all the universities
	for _, id := range nu.SelectedUniversities {
		u, err := uniModel.FindByID(ctx, cfg.Db, id)
		if err != nil {
			return model.User{}, fmt.Errorf("can't find selected uniform: %v", err)
		}
		as = append(as, model.Application{
			ID:             validate.GenerateID(),
			UserID:         id,
			UniversityName: u.Name,
			EmailContact:   u.DetailsURL,
			Status:         model.APPLICATION_OPEN_STATUS,
			CreatedAt:      now,
			DeleteAt:       time.Time{},
			UpdatedAt:      now,
		})
	}

	//upload the profile picture
	if _, err := cfg.AWSClient.S3.UploadToBucket(
		strings.NewReader(nu.ProfilePick),
		cfg.AWSClient.S3.S3UserProfilePictureBucket,
		id,
		"image/jpeg"); err != nil {
		return model.User{}, fmt.Errorf("can't upload profile picture: %v", err)
	}

	//upload all documents and format them to the document struct
	noteCertif := model.Document{
		ID:        validate.GenerateID(),
		Name:      "note certificate of" + nu.Name,
		UserID:    id,
		CreatedAt: now,
		DeleteAt:  time.Time{},
		UpdatedAt: now,
	}
	baccCertif := model.Document{
		ID:        validate.GenerateID(),
		Name:      "note certificate of" + nu.Name,
		UserID:    id,
		CreatedAt: now,
		DeleteAt:  time.Time{},
		UpdatedAt: now,
	}
	var err error

	fmt.Println(nu.NoteCertificate)

	if noteCertif.S3URL, err = cfg.AWSClient.S3.UploadToBucket(
		strings.NewReader(nu.NoteCertificate),
		cfg.AWSClient.S3.S3UserDocumentBucket,
		validate.GenerateID(),
		"application/pdf"); err != nil {
		return model.User{}, fmt.Errorf("can't upload the note certificate: %v", err)
	}

	if baccCertif.S3URL, err = cfg.AWSClient.S3.UploadToBucket(
		strings.NewReader(nu.BaccalaureateCertificate),
		cfg.AWSClient.S3.S3UserDocumentBucket,
		validate.GenerateID(),
		"application/pdf"); err != nil {
		return model.User{}, fmt.Errorf("can't upload the baccalaureate certificate: %v", err)
	}

	//Create the user inside cognito
	co := aws.CognitoUser{
		ID:          id,
		Email:       nu.Email,
		PhoneNumber: nu.PhoneNumber,
		Password:    nu.Password,
		IsActive:    true,
		Name:        nu.Name,
	}

	if err := cfg.AWSClient.Cognito.CreateUser(co); err != nil {
		return model.User{}, fmt.Errorf("can't create user: %v", err)
	}

	//save the user, document, application
	user := model.User{
		Id:             id,
		Name:           nu.Name,
		Email:          nu.Email,
		Town:           nu.Town,
		Country:        nu.Country,
		PhoneNumber:    nu.PhoneNumber,
		Birthday:       nu.Birthday,
		University:     nu.University,
		GraduationDate: nu.GraduationDate,
		Section:        nu.Section,
		EnrolledBy:     nu.EnrolledBy,
		CognitoID:      id,
		CreatedAt:      now,
		DeleteAt:       time.Time{},
		UpdatedAt:      now,
	}
	if err := model.Create(ctx, cfg.Db, user); err != nil {
		return model.User{}, fmt.Errorf("failed to save user %v", err)
	}

	if err := model.CreateDocs(ctx, cfg.Db, []model.Document{noteCertif, baccCertif}); err != nil {
		return model.User{}, fmt.Errorf("can't save documents: %v", err)
	}

	if err := model.CreateApplications(ctx, cfg.Db, as); err != nil {
		return model.User{}, fmt.Errorf("can't save applications")
	}

	return user, nil
}
