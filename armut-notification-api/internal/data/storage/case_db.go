package storage

import (
	"armut-notification-api/internal/util/env"
	"armut-notification-api/internal/util/logger"
	"armut-notification-api/internal/util/validator"
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"errors"
)

type INotificationDb interface {
	SetRatings(ch chan *SetRatingDbResponse, model *SetRatingDbModel)
	GetRatings(ch chan *GetRatingDbResponse, model *GetRatingDbModel)
	DeleteRatings(ch chan *DeleteRatingDbResponse, model *DeleteRatingDbModel)
}

type NotificationDb struct {
	loggr            logger.ILogger
	validatr         validator.IValidator
	environment      env.IEnvironment
	connectionString string
	driverName       string
	timeout          time.Duration
}

func NewNotificationDb(environment env.IEnvironment, loggr logger.ILogger, validatr validator.IValidator) INotificationDb {
	db := NotificationDb{
		environment:      environment,
		loggr:            loggr,
		validatr:         validatr,
		driverName:       "postgres",
		connectionString: environment.Get(env.PostgresqlConnectionString),
		timeout:          time.Second * 5,
	}
	return &db
}

func (d *NotificationDb) SetRatings(ch chan *SetRatingDbResponse, model *SetRatingDbModel) {
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

	query := `INSERT INTO Notification (service_provider_id, service_provier_rating) values ($1, $2)`
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
		ch <- &SetRatingDbResponse{Error: errors.New("could not add record")}
		d.loggr.Error("query failed.")
		return
	}
	ch <- &SetRatingDbResponse{}
	d.loggr.Info("setRating db ops successful.")
	close(ch)
}

func (d *NotificationDb) DeleteRatings(ch chan *DeleteRatingDbResponse, model *DeleteRatingDbModel) {
	modelErr := d.validatr.ValidateStruct(model)

	if modelErr != nil {
		ch <- &DeleteRatingDbResponse{Error: modelErr}
		d.loggr.Error("DeleteRatings model is not valid.")
		return
	}

	connection, err := sql.Open(d.driverName, d.connectionString)
	if err != nil {
		ch <- &DeleteRatingDbResponse{Error: err}

		return
	}
	defer connection.Close()
	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	query := `DELETE FROM Notification WHERE service_provider_id = $1;`
	result, err := connection.ExecContext(ctx, query, model.ServiceProviderId)

	if err != nil {
		ch <- &DeleteRatingDbResponse{Error: err}
		d.loggr.Error("db connection failed.")
		return
	}
	rows, err := result.RowsAffected()

	if err != nil {
		ch <- &DeleteRatingDbResponse{Error: err}
		d.loggr.Error("query failed.")
		return
	}
	if rows != 1 {
		ch <- &DeleteRatingDbResponse{Error: errors.New("could not get record")}
		d.loggr.Error("query failed.")
		return
	}
	ch <- &DeleteRatingDbResponse{}
	close(ch)
	d.loggr.Info("DeleteRatings db ops successful.")
}

func (d *NotificationDb) GetRatings(ch chan *GetRatingDbResponse, model *GetRatingDbModel) {
	modelErr := d.validatr.ValidateStruct(model)
	if modelErr != nil {
		ch <- &GetRatingDbResponse{Error: modelErr}
		return
	}

	connection, err := sql.Open(d.driverName, d.connectionString)
	if err != nil {
		ch <- &GetRatingDbResponse{Error: err}
		return
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), d.timeout)
	defer cancel()

	var notification GetRatingDbResponse

	query := `SELECT service_provier_rating FROM Notification where service_provider_id = $1;`

	rows, dbQueryErr := connection.QueryContext(ctx, query, model.ServiceProviderId)
	for rows.Next() {
		var serviceProvierRating int
		if err := rows.Scan(&serviceProvierRating); err != nil {
			ch <- &GetRatingDbResponse{Error: dbQueryErr}
			return
		}
		notification.Notifications = append(notification.Notifications, serviceProvierRating)
	}
	ch <- &notification
	close(ch)
	d.loggr.Info("GetRatings db ops successful.")

}
