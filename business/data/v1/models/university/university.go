package university

import (
	"context"
	"fmt"

	"github.com/Mahamadou828/AOAC/business/sys/database"
)

func Create(ctx context.Context, db *database.Database, nu University) error {
	if err := database.UpdateOrCreate[University](ctx, db, "university", nu); err != nil {
		return fmt.Errorf("can't save university in database: %v", err)
	}

	return nil
}

func Find(ctx context.Context, db *database.Database, startKey string, limit int64) ([]University, string, error) {
	var us []University

	lastEk, err := database.Find[[]University](ctx, db, "university", startKey, limit, &us)

	if err != nil {
		return us, "", fmt.Errorf("can't get university: %v", err)
	}

	return us, lastEk, nil
}

func FindByID(ctx context.Context, db *database.Database, id string) (University, error) {
	var us University

	if err := database.FindByID[University](ctx, db, id, "university", &us); err != nil {
		return us, fmt.Errorf("can't get university: %v", err)
	}

	return us, nil
}

func FindByCountry(ctx context.Context, db *database.Database, country, startEK string, limit int64) ([]University, string, error) {
	var us []University
	fmt.Println(startEK)
	inp := database.FindByIndexInput[University]{
		KeyName:   "country",
		KeyValue:  country,
		Index:     "countryIndex",
		TableName: "university",
		Dest:      &us,
		Limit:     limit,
		StartEK:   startEK,
	}

	lastEK, err := database.FindByIndex[University](ctx, db, inp)

	if err != nil {
		return us, "", fmt.Errorf("can't get for country %s: %v", country, err)
	}

	return us, lastEK, nil
}
