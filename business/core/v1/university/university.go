package university

import (
	"context"
	"fmt"

	model "github.com/Mahamadou828/AOAC/business/data/v1/models/university"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
)

func Find(ctx context.Context, cfg *lambda.Config, country, startEK string, limit int64) (lambda.FindResponse[model.University], error) {
	var us []model.University
	var lastEK string
	var err error

	switch true {
	case len(country) > 0:
		us, lastEK, err = model.FindByCountry(ctx, cfg.Db, country, startEK, limit)
		if err != nil {
			return lambda.FindResponse[model.University]{}, fmt.Errorf("can't get university by country %s: %v", country, err)
		}
		break
	default:
		us, lastEK, err = model.Find(ctx, cfg.Db, startEK, limit)
		if err != nil {
			return lambda.FindResponse[model.University]{}, fmt.Errorf("can't get university by country %s: %v", country, err)
		}
		break
	}

	return lambda.FindResponse[model.University]{LastEvaluatedKey: lastEK, Data: us}, nil
}

func FindByID(ctx context.Context, cfg *lambda.Config, id string) (model.University, error) {
	u, err := model.FindByID(ctx, cfg.Db, id)
	if err != nil {
		return model.University{}, fmt.Errorf("can't get university by id %s: %v", id, err)
	}

	return u, nil
}
