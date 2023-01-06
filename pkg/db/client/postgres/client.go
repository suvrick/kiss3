package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/suvrick/kiss/pkg/db"
)

type Client struct {
	Pool *pgxpool.Pool
}

func NewClient(config *db.DBConfig) (*Client, error) {

	dns := GetDNC(config)

	client, err := pgxpool.Connect(context.TODO(), dns)
	if err != nil {
		return nil, err
	}

	c := &Client{
		Pool: client,
	}

	err = c.initilize()

	return c, err
}

func (c *Client) Close() {
	c.Pool.Close()
}

func GetDNC(config *db.DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName)
}

func (c *Client) initilize() error {

	cmd := `
		CREATE TABLE IF NOT EXISTS PROXIES (
			id serial primary key,
			scheme text not null,
			host text not null,
			port int,
			username text,
			password text,
			at_create timestamp,
			at_update timestamp,
			is_bad bool
	);`

	_, err := c.Pool.Exec(context.Background(), cmd)

	if err != nil {
		return fmt.Errorf("error create proxies table. %s", err.Error())
	}

	cmd = `
	CREATE TABLE IF NOT EXISTS BOTS (
		id serial primary key,
		game_id int not null,
		name text not null,
		balance int not null,
		avatar text,
		profile text,
		at_create timestamp,
		at_update timestamp
	);`

	_, err = c.Pool.Exec(context.Background(), cmd)

	if err != nil {
		return fmt.Errorf("error create proxies table. %s", err.Error())
	}

	cmd = `
	CREATE TABLE IF NOT EXISTS USERS (
		id serial primary key,
		email text unique not null,
		password_hash text not null,
		role text not null,
		at_create timestamp not null,
		at_update timestamp not null
	);`

	_, err = c.Pool.Exec(context.Background(), cmd)

	if err != nil {
		return fmt.Errorf("error create proxies table. %s", err.Error())
	}

	return nil
}
