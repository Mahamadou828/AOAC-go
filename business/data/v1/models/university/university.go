package university

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/business/sys/database"
)

func Create(ctx context.Context, db *database.Database, nu University) error {
	if err := database.PutOrCreateItem[University](ctx, db, "university", nu); err != nil {
		return fmt.Errorf("can't save university in database: %v", err)
	}

	return nil
}

func Query(ctx context.Context, db *database.Database) ([]University, error) {
	var us []University

	if err := database.GetItems[[]University](ctx, db, "university", &us); err != nil {
		return us, fmt.Errorf("can't get university: %v", err)
	}

	return us, nil
}

func QueryByID(ctx context.Context, db *database.Database, id string) (University, error) {
	var us University

	if err := database.GetItemByUniqueKey[University](ctx, db, id, "id", "university", &us); err != nil {
		return us, fmt.Errorf("can't get university: %v", err)
	}

	return us, nil
}

func QueryByCountry(ctx context.Context, db *database.Database, country string) ([]University, error) {
	var us []University

	if err := database.GetItemsByIndex[[]University](ctx, db, country, "country", "countryIndex", "university", &us); err != nil {
		return us, fmt.Errorf("can't get for country %s: %v", country, err)
	}

	return us, nil
}
