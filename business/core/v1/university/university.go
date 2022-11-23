package university

import (
	"context"
	"fmt"

	model "github.com/Mahamadou828/AOAC/business/data/v1/models/university"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
)

func Query(ctx context.Context, cfg *lambda.Config, page, rowsPerPage int, country string) ([]model.University, error) {
	var us []model.University
	var err error

	switch true {
	case len(country) > 0:
		us, err = model.QueryByCountry(ctx, cfg.Db, country)
		if err != nil {
			return us, fmt.Errorf("can't get university by country %s: %v", country, err)
		}
		break
	default:
		us, err = model.Query(ctx, cfg.Db)
		if err != nil {
			return us, fmt.Errorf("can't get university by country %s: %v", country, err)
		}
		break
	}

	return us, nil
}

func QueryByID(ctx context.Context, cfg *lambda.Config, id string) (model.University, error) {
	u, err := model.QueryByID(ctx, cfg.Db, id)
	if err != nil {
		return model.University{}, fmt.Errorf("can't get university by id %s: %v", id, err)
	}

	return u, nil
}
