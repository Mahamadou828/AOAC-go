package admin

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/business/sys/database"
)

// Create a new admin
func Create(ctx context.Context, db *database.Database, na Admin) error {
	if err := database.UpdateOrCreate[Admin](ctx, db, "admin", na); err != nil {
		return fmt.Errorf("unable to create admin: %v", err)
	}
	return nil
}

// Find fetch all admin, send the start key to have the pagination added
func Find(ctx context.Context, db *database.Database, startKey string, limit int64) ([]Admin, string, error) {
	var as []Admin
	lastEK, err := database.Find[[]Admin](ctx, db, "admin", startKey, limit, &as)

	if err != nil {
		return as, "", fmt.Errorf("can't get admin: %v", err)
	}
	return as, lastEK, nil
}

// FindByEmail fetch all admin that match the given email
func FindByEmail(ctx context.Context, db *database.Database, email, startKey string, limit int64) ([]Admin, string, error) {
	var as []Admin

	lastEK, err := database.FindByIndex(ctx, db, database.FindByIndexInput[Admin]{
		KeyName:   "email",
		KeyValue:  email,
		Index:     "emailIndex",
		TableName: "admin",
		Dest:      &as,
		Limit:     limit,
		StartEK:   startKey,
	})

	if err != nil {
		return as, "", fmt.Errorf("can't get admin with the given email: %v, %v", email, err)
	}
	return as, lastEK, nil
}

// FindByID fetch an admin by the given ID
func FindByID(ctx context.Context, db *database.Database, id string) (Admin, error) {
	var a Admin
	if err := database.FindByID[Admin](ctx, db, id, "admin", &a); err != nil {
		return a, fmt.Errorf("can't get admin by ID: %v", err)
	}
	return a, nil
}

// FindOneByEmail fetch an admin by the given Email
func FindOneByEmail(ctx context.Context, db *database.Database, email string) (Admin, error) {
	var a Admin

	inp := database.FindOneByIndexInput[Admin]{
		KeyName:   "email",
		KeyValue:  email,
		Index:     "emailIndex",
		TableName: "admin",
		Dest:      &a,
	}

	if err := database.FindOneByIndex[Admin](ctx, db, inp); err != nil {
		return a, fmt.Errorf("can't get admin by email: %v", err)
	}
	return a, nil
}

// Update admin info
func Update(ctx context.Context, db *database.Database, na Admin) error {
	if err := database.UpdateOrCreate[Admin](ctx, db, "admin", na); err != nil {
		return fmt.Errorf("unable to update admin info: %v", err)
	}
	return nil
}

// Delete admin info
func Delete(ctx context.Context, db *database.Database, id string) error {
	if err := database.Delete(ctx, db, "admin", id); err != nil {
		return fmt.Errorf("unable to delete admin: %v", err)
	}
	return nil
}
