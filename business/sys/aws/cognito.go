package aws

import (
	"fmt"
	"strconv"

	sdkaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Cognito struct {
	ClientID         string
	UserPoolID       string
	identityProvider *cognitoidentityprovider.CognitoIdentityProvider
}

type CognitoUser struct {
	ID          string
	Email       string
	PhoneNumber string
	Password    string
	IsActive    bool
	Name        string
}

type Session struct {
	Token        string
	RefreshToken string
	ExpireIn     int64
}

func NewCognito(sess *session.Session, clientID string, userPoolID string) *Cognito {
	return &Cognito{
		ClientID:         clientID,
		UserPoolID:       userPoolID,
		identityProvider: cognitoidentityprovider.New(sess),
	}
}

// CreateUser creates a new user inside the cognito pool and return his id.
func (c *Cognito) CreateUser(u CognitoUser) error {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: sdkaws.String(c.ClientID),
		Password: sdkaws.String(u.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  sdkaws.String("email"),
				Value: sdkaws.String(u.Email),
			},
			{
				Name:  sdkaws.String("phone_number"),
				Value: sdkaws.String(u.PhoneNumber),
			},
			{
				Name:  sdkaws.String("name"),
				Value: sdkaws.String(u.Name),
			},
			{
				Name:  sdkaws.String("custom:isActive"),
				Value: sdkaws.String(strconv.FormatBool(u.IsActive)),
			},
		},
		Username: sdkaws.String(u.ID),
	}

	if _, err := c.identityProvider.SignUp(input); err != nil {
		return fmt.Errorf("can't create new cognito user: %v", err)
	}

	return nil
}

// ConfirmSignUp validate a newly create account
func (c *Cognito) ConfirmSignUp(code, id string) error {
	inp := cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         sdkaws.String(c.ClientID),
		ConfirmationCode: sdkaws.String(code),
		Username:         sdkaws.String(id),
	}

	if _, err := c.identityProvider.ConfirmSignUp(&inp); err != nil {
		return err
	}

	//Update the user isActive attribute to true
	attr := []*cognitoidentityprovider.AttributeType{
		{
			Name:  sdkaws.String("isActive"),
			Value: sdkaws.String("true"),
		},
	}

	if err := c.updateUserAttribute(id, attr); err != nil {
		return err
	}

	return nil
}

// AuthenticateUser authenticate a new user, and return
// identification data.
func (c *Cognito) AuthenticateUser(id, password string) (Session, error) {
	inp := cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: sdkaws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": sdkaws.String(id),
			"SRP_A":    sdkaws.String(password),
			"PASSWORD": sdkaws.String(password),
		},
		ClientId: sdkaws.String(c.ClientID),
	}

	out, err := c.identityProvider.InitiateAuth(&inp)

	if err != nil {
		return Session{}, err
	}

	return Session{
		Token:        *out.AuthenticationResult.AccessToken,
		RefreshToken: *out.AuthenticationResult.RefreshToken,
		ExpireIn:     *out.AuthenticationResult.ExpiresIn,
	}, nil
}

func (c *Cognito) ForgotPassword(id string) error {
	inp := cognitoidentityprovider.ForgotPasswordInput{
		ClientId: sdkaws.String(c.ClientID),
		Username: sdkaws.String(id),
	}

	if _, err := c.identityProvider.ForgotPassword(&inp); err != nil {
		return err
	}

	return nil
}

func (c *Cognito) RefreshToken(token string) (Session, error) {
	inp := cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: sdkaws.String(cognitoidentityprovider.AuthFlowTypeRefreshToken),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": sdkaws.String(token),
		},
		ClientId: sdkaws.String(c.ClientID),
	}

	out, err := c.identityProvider.InitiateAuth(&inp)

	if err != nil {
		return Session{}, err
	}
	return Session{
		Token:        *out.AuthenticationResult.AccessToken,
		ExpireIn:     *out.AuthenticationResult.ExpiresIn,
		RefreshToken: token,
	}, nil
}

// DeleteUser will completely delete the user from the pool
// the user will not be able to recover the account.
func (c *Cognito) DeleteUser(id string) error {
	inp := cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: sdkaws.String(c.UserPoolID),
		Username:   sdkaws.String(id),
	}

	if _, err := c.identityProvider.AdminDeleteUser(&inp); err != nil {
		return err
	}

	return nil
}

func (c *Cognito) ResendValidateCode(sub string) error {
	inp := cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId: sdkaws.String(c.ClientID),
		Username: sdkaws.String(sub),
	}

	if _, err := c.identityProvider.ResendConfirmationCode(&inp); err != nil {
		return err
	}

	return nil
}

func (c *Cognito) UpdateUser(user CognitoUser) error {
	attr := []*cognitoidentityprovider.AttributeType{
		{Name: sdkaws.String("email"), Value: sdkaws.String(user.Email)},
		{Name: sdkaws.String("phone_number"), Value: sdkaws.String(user.PhoneNumber)},
		{Name: sdkaws.String("name"), Value: sdkaws.String(user.Email)},
	}

	return c.updateUserAttribute(user.ID, attr)
}

func (c *Cognito) updateUserAttribute(sub string, attr []*cognitoidentityprovider.AttributeType) error {
	inp := cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: attr,
		UserPoolId:     sdkaws.String(c.UserPoolID),
		Username:       sdkaws.String(sub),
	}

	if _, err := c.identityProvider.AdminUpdateUserAttributes(&inp); err != nil {
		return err
	}
	return nil
}
