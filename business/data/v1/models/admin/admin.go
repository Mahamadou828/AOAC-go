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
		return as, err
	}
	return as, nil
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
