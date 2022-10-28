package admin

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/business/sys/database"
)

// Create a new admin
func Create(ctx context.Context, db *database.Database, na Admin) error {
	if err := database.PutOrCreateItem[Admin](ctx, db, "admin", na); err != nil {
		return fmt.Errorf("unable to create admin: %v", err)
	}
	return nil
}

// Query fetch all admin, no pagination for now
// @todo add pagination
func Query(ctx context.Context, db *database.Database) ([]Admin, error) {
	var as []Admin
	if err := database.GetItems[[]Admin](ctx, db, "admin", &as); err != nil {
		return as, fmt.Errorf("can't get admin: %v", err)
	}
	return as, nil
}

// QueryByID fetch an admin by the given ID
func QueryByID(ctx context.Context, db *database.Database, id string) (Admin, error) {
	var a Admin
	if err := database.GetItemByUniqueKey[Admin](ctx, db, id, "id", "admin", &a); err != nil {
		return a, fmt.Errorf("can't get admin by ID: %v", err)
	}
	return a, nil
}

// QueryByEmail fetch an admin by the given Email
func QueryByEmail(ctx context.Context, db *database.Database, email string) (Admin, error) {
	var a Admin
	if err := database.GetItemByIndex[Admin](ctx, db, email, "email", "emailIndex", "admin", &a); err != nil {
		return a, fmt.Errorf("can't get admin by email: %v", err)
	}
	return a, nil
}

// Update admin info
func Update(ctx context.Context, db *database.Database, na Admin) error {
	if err := database.PutOrCreateItem[Admin](ctx, db, "admin", na); err != nil {
		return fmt.Errorf("unable to update admin info: %v", err)
	}
	return nil
}

func Delete(ctx context.Context, db *database.Database, id string) error {
	if err := database.DeleteItem(ctx, db, "admin", id); err != nil {
		return fmt.Errorf("unable to delete admin: %v", err)
	}
	return nil
}
