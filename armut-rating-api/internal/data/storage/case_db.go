package storage

import (
	"armut-rating-api/internal/util/env"
	"armut-rating-api/internal/util/logger"
	"armut-rating-api/internal/util/validator"
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"time"
)

type IRatingDb interface {
	SetRatings(ch chan *SetRatingDbResponse, model *SetRatingDbModel)
	GetRatingAverage(ch chan *GetRatingDbResponse, model *GetRatingDbModel)
}

type RatingDb struct {
	loggr            logger.ILogger
	validatr         validator.IValidator
	environment      env.IEnvironment
	connectionString string
	driverName       string
	timeout          time.Duration
}

func NewRatingDb(environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) IRatingDb {
	db := RatingDb{
		environment:      environment,
		loggr:            loggr,
		validatr:         validatr,
		driverName:       "postgres",
		connectionString: environment.Get(env.PostgresqlConnectionString),
		timeout:          time.Second * 5,
	}
	return &db
}

func (d *RatingDb) SetRatings(ch chan *SetRatingDbResponse, model *SetRatingDbModel) {
	modelErr := d.validatr.ValidateStruct(model)
	if modelErr != nil {
		ch <- &SetRatingDbResponse{Error: modelErr}
		d.loggr.Error("setRatings model is not valid.")
		return
	}

	connection, err := sql.Open(d.driverName, d.connectionString)
	if err != nil {
		ch <- &SetRatingDbResponse{Error: err}
		d.loggr.Error("db connection failed.")
		return
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	query := `INSERT INTO Rating (service_provider_id, service_provier_rating) values ($1, $2)`
	result, err := connection.ExecContext(ctx, query, model.ServiceProviderId, model.ServiceProviderRating)

	if err != nil {
		ch <- &SetRatingDbResponse{Error: err}
		d.loggr.Error("query failed.")
		return
	}
	rows, err := result.RowsAffected()

	if err != nil {
		ch <- &SetRatingDbResponse{Error: err}
		d.loggr.Error("query failed.")
		return
	}
	if rows != 1 {
		ch <- &SetRatingDbResponse{Error: errors.New("could not set rating.")}
		d.loggr.Error("could not set rating.")
		return
	}
	ch <- &SetRatingDbResponse{}
	d.loggr.Info("setRating db ops successful.")
	close(ch)
}
func (d *RatingDb) GetRatingAverage(ch chan *GetRatingDbResponse, model *GetRatingDbModel) {
	modelErr := d.validatr.ValidateStruct(model)
	if modelErr != nil {
		ch <- &GetRatingDbResponse{Error: modelErr}
		d.loggr.Error("getRatingAverage model is not valid.")
		return
	}

	connection, err := sql.Open(d.driverName, d.connectionString)
	if err != nil {
		ch <- &GetRatingDbResponse{Error: err}
		d.loggr.Error("db connection failed.")
		return
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	var averageData GetRatingDbResponse

	query := `SELECT AVG(service_provier_rating) as average_rating, COUNT(*) as total_rating_count FROM Rating WHERE service_provider_id = $1;`

	dbQueryErr := connection.QueryRowContext(ctx, query, model.ServiceProviderId).Scan(
		&averageData.AverageRating,
		&averageData.TotalRatingCount,
	)
	if dbQueryErr != nil {
		ch <- &GetRatingDbResponse{Error: dbQueryErr}
		d.loggr.Error("query failed.")
		return
	}

	ch <- &averageData
	d.loggr.Info("getRatingAverage db ops successful.")
	close(ch)
}
