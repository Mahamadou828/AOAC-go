package user

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/business/sys/database"
)

// Create a new user
func Create(ctx context.Context, db *database.Database, nu User) error {
	if err := database.UpdateOrCreate[User](ctx, db, "user", nu); err != nil {
		return fmt.Errorf("can't create an user: %v", err)
	}
	return nil
}

// Find return a list of users with pagination
func Find(ctx context.Context, db *database.Database, startKey string, limit int64) ([]User, string, error) {
	var us []User
	lastEK, err := database.Find[[]User](ctx, db, "user", startKey, limit, &us)
	if err != nil {
		return us, "", fmt.Errorf("can't find users: %v", err)
	}

	return us, lastEK, err
}

// FindByName return a list of users that match the given name
// @todo implement a search like system
func FindByName(ctx context.Context, db *database.Database, name, startKey string, limit int64) ([]User, string, error) {
	var us []User

	inp := database.FindByIndexInput[User]{
		KeyName:   "name",
		KeyValue:  name,
		Index:     "nameIndex",
		TableName: "user",
		Dest:      &us,
		Limit:     limit,
		StartEK:   startKey,
	}

	lastEK, err := database.FindByIndex[User](ctx, db, inp)
	if err != nil {
		return us, "", fmt.Errorf("can't find user by the given name: %v, %v", name, err)
	}

	return us, lastEK, nil
}

// FindByID returns a user by the given ID. The request will return a error if the user does not exist.
func FindByID(ctx context.Context, db *database.Database, id string) (User, error) {
	var u User

	if err := database.FindByID[User](ctx, db, id, "user", &u); err != nil {
		return User{}, fmt.Errorf("can't find user by the given id: %v, %v", id, err)
	}

	return u, nil
}

// FindOneByEmail return a user that match the given email
func FindOneByEmail(ctx context.Context, db *database.Database, email, startKey string, limit int64) ([]User, string, error) {
	var us []User

	inp := database.FindByIndexInput[User]{
		KeyName:   "email",
		KeyValue:  email,
		Index:     "emailIndex",
		TableName: "user",
		Dest:      &us,
		Limit:     limit,
		StartEK:   startKey,
	}

	lastEK, err := database.FindByIndex[User](ctx, db, inp)
	if err != nil {
		return us, "", fmt.Errorf("can't find user by the given email: %v, %v", email, err)
	}

	return us, lastEK, nil
}

// Update the data of a user
func Update(ctx context.Context, db *database.Database, nu User) (User, error) {
	if err := database.UpdateOrCreate[User](ctx, db, "user", nu); err != nil {
		return User{}, fmt.Errorf("can't update user: %v, %v", nu.Email, err)
	}

	return nu, nil
}

// Delete a given user
func Delete(ctx context.Context, db *database.Database, id string) error {
	if err := database.Delete(ctx, db, "user", id); err != nil {
		return fmt.Errorf("can't delete the given user: %v, %v", id, err)
	}
	return nil
}

// CreateApplications save a list of application
func CreateApplications(ctx context.Context, db *database.Database, nas []Application) error {
	if err := database.BatchWrite[Application](ctx, db, "application", nas); err != nil {
		return fmt.Errorf("can't create application: %v, %v", nas, err)
	}

	return nil
}

// CreateDocs save a list of document
func CreateDocs(ctx context.Context, db *database.Database, nds []Document) error {
	if err := database.BatchWrite[Document](ctx, db, "document", nds); err != nil {
		return fmt.Errorf("can't create documents: %v, %v", nds, err)
	}

	return nil
}

// FindApplications returns a list of application for the given user id
func FindApplications(ctx context.Context, db *database.Database, userID string) ([]Application, error) {
	var as []Application

	inp := database.FindByIndexInput[Application]{
		KeyName:   "userID",
		KeyValue:  userID,
		Index:     "userIDIndex",
		TableName: "application",
		Dest:      &as,
		//The limit is 5 because you can't have more than 5 applications per user.
		Limit: 5,
	}

	if _, err := database.FindByIndex(ctx, db, inp); err != nil {
		return as, fmt.Errorf("can't find application for user: %s, %v", userID, err)
	}

	return as, nil
}

// FindDocuments return a list of documents for the given user id.
func FindDocuments(ctx context.Context, db *database.Database, userID, startEK string, limit int64) ([]Document, string, error) {
	var res []Document

	inp := database.FindByIndexInput[Document]{
		KeyName:   "document",
		KeyValue:  userID,
		Index:     "userIDIndex",
		TableName: "document",
		Dest:      &res,
		Limit:     limit,
		StartEK:   startEK,
	}

	lastEK, err := database.FindByIndex(ctx, db, inp)

	if err != nil {
		return res, "", fmt.Errorf("can't find documents for user: %v, %v", userID, err)
	}

	return res, lastEK, nil
}

// DeleteDocument delete a given document
func DeleteDocument(ctx context.Context, db *database.Database, docID string) error {
	if err := database.Delete(ctx, db, "document", docID); err != nil {
		return fmt.Errorf("can't delete document: %v", err)
	}
	return nil
}

// UpdateApplication update a user application. For now you can update only one application at the time
func UpdateApplication(ctx context.Context, db *database.Database, na Application) error {
	if err := database.UpdateOrCreate[Application](ctx, db, "application", na); err != nil {
		return fmt.Errorf("can't update application: %v", err)
	}
	return nil
}
