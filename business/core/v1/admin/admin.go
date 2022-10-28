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

//@todo implement pagination

type Session struct {
	Admin        admin.Admin `json:"admin"`
	RefreshToken string      `json:"refreshToken"`
	ExpiresIn    int64       `json:"expiresIn"`
	Token        string      `json:"token"`
}

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
		Name:        fmt.Sprintf("%s %s", na.Name, na.Surname),
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

func Query(ctx context.Context, cfg *lambda.Config, page, rowsPerPage int) ([]admin.Admin, error) {
	res, err := admin.Query(ctx, cfg.Db)
	if err != nil {
		return nil, fmt.Errorf("can't retrieve admin: %v", err)
	}
	return res, nil
}

func QueryByID(ctx context.Context, cfg *lambda.Config, id string) (admin.Admin, error) {
	res, err := admin.QueryByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't retrieve admin: %v", err)
	}
	return res, nil
}

func Update(ctx context.Context, cfg *lambda.Config, id string, ua admin.UpdateAdminDTO, now time.Time) (admin.Admin, error) {
	res, err := admin.QueryByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't retrieve admin: %v", err)
	}
	cognitoData := aws.CognitoUser{
		ID:          res.CognitoID,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Name:        res.Name,
	}

	if ua.Name != nil {
		cognitoData.Name = *ua.Name
		res.Name = *ua.Name
	}
	if ua.Email != nil {
		cognitoData.Email = *ua.Email
		res.Email = *ua.Email
	}
	if ua.Surname != nil {
		res.Surname = *ua.Surname
	}
	if ua.Role != nil {
		res.Role = *ua.Role
	}
	if ua.PhoneNumber != nil {
		cognitoData.PhoneNumber = *ua.PhoneNumber
		res.PhoneNumber = *ua.PhoneNumber
	}

	//Update user data in dynamodb
	if err := admin.Update(ctx, cfg.Db, res); err != nil {
		return admin.Admin{}, fmt.Errorf("can't update admin: %v", err)
	}
	//Update user data in cognito
	if err := cfg.AWSClient.Cognito.UpdateUser(cognitoData); err != nil {
		return admin.Admin{}, fmt.Errorf("can't update user: %v", err)
	}

	return res, nil
}

func Delete(ctx context.Context, cfg *lambda.Config, id string, now time.Time) (admin.Admin, error) {
	res, err := admin.QueryByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't find admin: %v", err)
	}

	if err := cfg.AWSClient.Cognito.DeleteUser(res.CognitoID); err != nil {
		return admin.Admin{}, fmt.Errorf("can't delete admin from cognito: %v", err)
	}

	if err := admin.Delete(ctx, cfg.Db, id); err != nil {
		return admin.Admin{}, fmt.Errorf("can't delete admin: %v", err)
	}

	return res, nil
}

func Login(ctx context.Context, cfg *lambda.Config, data admin.LoginAdminDTO) (Session, error) {
	res, err := admin.QueryByEmail(ctx, cfg.Db, data.Email)
	if err != nil {
		return Session{}, fmt.Errorf("can't find admin: %v", err)
	}

	tokens, err := cfg.AWSClient.Cognito.AuthenticateUser(res.ID, data.Password)
	if err != nil {
		return Session{}, fmt.Errorf("can't authenticate user: %v", err)
	}

	return Session{
		Admin:        res,
		RefreshToken: tokens.RefreshToken,
		Token:        tokens.Token,
		ExpiresIn:    tokens.ExpireIn,
	}, nil
}

func RefreshToken(ctx context.Context, cfg *lambda.Config, data admin.RefreshTokenDTO) (Session, error) {
	res, err := admin.QueryByID(ctx, cfg.Db, data.ID)
	if err != nil {
		return Session{}, fmt.Errorf("can't find admin: %v", err)
	}

	tokens, err := cfg.AWSClient.Cognito.RefreshToken(data.RefreshToken)
	if err != nil {
		return Session{}, fmt.Errorf("can't authenticate user: %v", err)
	}

	return Session{
		Admin:        res,
		RefreshToken: tokens.RefreshToken,
		Token:        tokens.Token,
		ExpiresIn:    tokens.ExpireIn,
	}, nil
}
