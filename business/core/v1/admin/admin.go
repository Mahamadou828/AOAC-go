package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/Mahamadou828/AOAC/business/data/v1/models/admin"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
)

func Create(ctx context.Context, cfg *lambda.Config, na admin.NewAdminDTO, now time.Time) (admin.Admin, error) {
	//Generate admin id
	id := validate.GenerateID()

	//create user inside cognito
	co := aws.CognitoUser{
		ID:          id,
		Email:       na.Email,
		PhoneNumber: na.PhoneNumber,
		Password:    na.Password,
		IsActive:    true,
	}
	if err := cfg.AWSClient.Cognito.CreateUser(co); err != nil {
		return admin.Admin{}, fmt.Errorf("can't create user, exception during cognito registration: %v", err)
	}

	//upload profile picture to aws s3 and store the file name

	//register user inside dynamodb
	newAdmin := admin.Admin{
		ID:           id,
		Name:         na.Name,
		Surname:      na.Surname,
		Email:        na.Email,
		PhoneNumber:  na.PhoneNumber,
		Role:         na.Role,
		EnrolledUser: nil,
		CognitoID:    id,
		ProfilePick:  "",
		CreatedAt:    now,
		DeleteAt:     time.Time{},
		UpdatedAt:    now,
	}
	if err := admin.Create(ctx, cfg.Db, newAdmin); err != nil {
		return admin.Admin{}, fmt.Errorf("error creating admin: %v", err)
	}

	return newAdmin, nil
}

//func Query(ctx context.Context, cfg *lambda.Config) error {
//	return admin.Query(ctx, cfg.Db)
//}

func QueryByID(ctx context.Context, cfg *lambda.Config) {

}

func Update() {

}

func Delete() {

}

func Login() {

}

func RefreshToken() {

}
