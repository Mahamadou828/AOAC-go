package admin

import (
	"context"
	"fmt"
	"github.com/Mahamadou828/AOAC/app/tools/config"
	"github.com/aws/aws-lambda-go/events"
	"strings"
	"time"

	"github.com/Mahamadou828/AOAC/business/data/v1/models/admin"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
	"github.com/Mahamadou828/AOAC/business/sys/validate"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
)

//@todo implement pagination

type Session struct {
	Admin          admin.Admin `json:"admin"`
	RefreshToken   string      `json:"refreshToken"`
	ExpiresIn      int64       `json:"expiresIn"`
	Token          string      `json:"token"`
	ProfilePickUrl string      `json:"profilePickUrl"`
}

func Create(ctx context.Context, cfg *lambda.Config, r events.APIGatewayProxyRequest, now time.Time) (admin.Admin, error) {
	//Generate admin id
	id := validate.GenerateID()

	na := admin.NewAdminDTO{}

	if err := config.ParseForm(&na, r); err != nil {
		return admin.Admin{}, fmt.Errorf("failed to parse form: %v", err)
	}

	if err := validate.Check(na); err != nil {
		return admin.Admin{}, fmt.Errorf("invalid body: %v", err)
	}

	//upload profile picture to aws s3, the file key will be the id of the admin
	if _, err := cfg.AWSClient.S3.UploadToBucket(strings.NewReader(na.ProfilePick), cfg.AWSClient.S3.S3AdminProfilePictureBucket, id, "image/jpeg"); err != nil {
		return admin.Admin{}, fmt.Errorf("failed to upload profile picture: %v", err)
	}

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
		ProfilePick:  id,
		CreatedAt:    now,
		DeleteAt:     time.Time{},
		UpdatedAt:    now,
	}
	if err := admin.Create(ctx, cfg.Db, newAdmin); err != nil {
		return admin.Admin{}, fmt.Errorf("error creating admin: %v", err)
	}

	return newAdmin, nil
}

func Find(ctx context.Context, cfg *lambda.Config, email, startKey string, limit int64) (lambda.FindResponse[admin.Admin], error) {

	switch true {
	case len(email) > 0:
		result, lastEk, err := admin.FindByEmail(ctx, cfg.Db, email, startKey, limit)
		return lambda.FindResponse[admin.Admin]{LastEvaluatedKey: lastEk, Data: result}, err
	default:
		result, lastEK, err := admin.Find(ctx, cfg.Db, startKey, limit)
		if err != nil {
			return lambda.FindResponse[admin.Admin]{}, fmt.Errorf("can't retrieve admin: %v", err)
		}
		return lambda.FindResponse[admin.Admin]{LastEvaluatedKey: lastEK, Data: result}, nil
	}

}

func FindByID(ctx context.Context, cfg *lambda.Config, id string) (admin.Admin, error) {
	res, err := admin.FindByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't retrieve admin: %v", err)
	}
	return res, nil
}

func Update(ctx context.Context, cfg *lambda.Config, id string, r events.APIGatewayProxyRequest, now time.Time) (admin.Admin, error) {
	var ua admin.UpdateAdminDTO

	if err := config.ParseForm(&ua, r); err != nil {
		return admin.Admin{}, fmt.Errorf("can't parse form: %v", err)
	}
	if err := validate.Check(ua); err != nil {
		return admin.Admin{}, fmt.Errorf("invalid request: %v", err)
	}

	res, err := admin.FindByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't retrieve admin: %v", err)
	}
	cognitoData := aws.CognitoUser{
		ID:          res.CognitoID,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Name:        res.Name,
	}

	if ua.Name != "" {
		cognitoData.Name = ua.Name
		res.Name = ua.Name
	}
	if ua.Email != "" {
		cognitoData.Email = ua.Email
		res.Email = ua.Email
	}
	if ua.Surname != "" {
		res.Surname = ua.Surname
	}
	if ua.Role != "" {
		res.Role = ua.Role
	}
	if ua.PhoneNumber != "" {
		cognitoData.PhoneNumber = ua.PhoneNumber
		res.PhoneNumber = ua.PhoneNumber
	}

	if ua.ProfilePick != "" {
		err := cfg.AWSClient.S3.UpdateObject(strings.NewReader(ua.ProfilePick), cfg.AWSClient.S3.S3AdminProfilePictureBucket, res.ProfilePick, "image/jpg")

		if err != nil {
			return admin.Admin{}, fmt.Errorf("failed to update profile: %v", err)
		}
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
	res, err := admin.FindByID(ctx, cfg.Db, id)
	if err != nil {
		return admin.Admin{}, fmt.Errorf("can't find admin: %v", err)
	}

	if err := cfg.AWSClient.Cognito.DeleteUser(res.CognitoID); err != nil {
		return admin.Admin{}, fmt.Errorf("can't delete admin from cognito: %v", err)
	}

	if err := admin.Delete(ctx, cfg.Db, id); err != nil {
		return admin.Admin{}, fmt.Errorf("can't delete admin: %v", err)
	}

	if err := cfg.AWSClient.S3.DeleteObject(cfg.AWSClient.S3.S3AdminProfilePictureBucket, res.ProfilePick); err != nil {
		return admin.Admin{}, fmt.Errorf("can't delete admin profile picture: %v", err)
	}
	return res, nil
}

func Login(ctx context.Context, cfg *lambda.Config, data admin.LoginAdminDTO) (Session, error) {
	res, err := admin.FindOneByEmail(ctx, cfg.Db, data.Email)
	if err != nil {
		return Session{}, fmt.Errorf("can't find admin: %v", err)
	}

	tokens, err := cfg.AWSClient.Cognito.AuthenticateUser(res.ID, data.Password)
	if err != nil {
		return Session{}, fmt.Errorf("can't authenticate user: %v", err)
	}

	url, _ := cfg.AWSClient.S3.GeneratePresignedUrl(cfg.AWSClient.S3.S3AdminProfilePictureBucket, res.ProfilePick)

	return Session{
		Admin:          res,
		RefreshToken:   tokens.RefreshToken,
		Token:          tokens.Token,
		ExpiresIn:      tokens.ExpireIn,
		ProfilePickUrl: url,
	}, nil
}

func RefreshToken(ctx context.Context, cfg *lambda.Config, data admin.RefreshTokenDTO) (Session, error) {
	res, err := admin.FindByID(ctx, cfg.Db, data.ID)
	if err != nil {
		return Session{}, fmt.Errorf("can't find admin: %v", err)
	}

	tokens, err := cfg.AWSClient.Cognito.RefreshToken(data.RefreshToken)
	if err != nil {
		return Session{}, fmt.Errorf("can't authenticate user: %v", err)
	}

	url, _ := cfg.AWSClient.S3.GeneratePresignedUrl(cfg.AWSClient.S3.S3AdminProfilePictureBucket, res.ProfilePick)

	return Session{
		Admin:          res,
		RefreshToken:   tokens.RefreshToken,
		Token:          tokens.Token,
		ExpiresIn:      tokens.ExpireIn,
		ProfilePickUrl: url,
	}, nil
}

func ConfirmSignUp(data admin.ConfirmSignupDTO, cfg *lambda.Config) error {
	if err := cfg.AWSClient.Cognito.ConfirmSignUp(data.Code, data.UserID); err != nil {
		return fmt.Errorf("can't confirm sign up: %v", err)
	}

	return nil
}
