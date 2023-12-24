package postgres

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	timezone                 = "utc"
	pingTimeout              = 1000 * time.Millisecond
	maxPingAttempts          = 10
	sleepTimeBetweenAttempts = 100
)

type (
	Config struct {
		User         string
		Password     string
		Host         string
		DBName       string
		Schema       string
		MaxIdleConns int
		MaxOpenConns int
		DisableTLS   bool
	}

	Client struct {
		*sqlx.DB
	}
)

func NewClient(config Config) (*Client, error) {
	sslMode := "require"
	if config.DisableTLS {
		sslMode = "disable"
	}

	query := make(url.Values)

	query.Set("sslmode", sslMode)
	query.Set("timezone", timezone)

	if config.Schema != "" {
		query.Set("search_path", config.Schema)
	}

	URL := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.User, config.Password),
		Host:     config.Host,
		Path:     config.DBName,
		RawQuery: query.Encode(),
	}

	db, err := sqlx.Open("postgres", URL.String())
	if err != nil {
		return &Client{}, err
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	client := &Client{db}

	if err := client.Ping(context.Background()); err != nil {
		return &Client{}, errors.New("failed to connect to the database")
	}

	return client, nil
}

func (client *Client) Ping(ctx context.Context) error {
	if _, ok := ctx.Deadline(); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, pingTimeout)
		defer cancel()
	}

	for attempts := 1; ; attempts++ {
		err := client.DB.Ping()
		if err == nil {
			break
		}

		if attempts == maxPingAttempts {
			return err
		}

		time.Sleep(time.Duration(attempts) * sleepTimeBetweenAttempts * time.Microsecond)
	}

	var result bool
	return client.DB.QueryRowContext(ctx, "SELECT true").Scan(&result)
}

func (client *Client) Close() error {
	return client.DB.Close()
}
