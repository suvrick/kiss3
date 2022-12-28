package proxy

import (
	"context"
	"fmt"

	"github.com/suvrick/kiss/pkg/db/client/postgres"
)

type ProxyRepository struct {
	client *postgres.Client
}

func NewProxyRepository(db *postgres.Client) *ProxyRepository {
	return &ProxyRepository{
		client: db,
	}
}

func (repo *ProxyRepository) Create(ctx context.Context, p Proxy) (Proxy, error) {

	q := `INSERT INTO proxies(
			scheme, 
			host, 
			port, 
			username, 
			password, 
			at_create, 
			at_update, 
			is_bad
		) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id`

	row := repo.client.Pool.QueryRow(ctx, q, &p.Scheme, &p.Host, &p.Port, &p.Username, &p.Password, &p.AtCreate, &p.AtUpdate, &p.IsBad)

	err := row.Scan(&p.ID)

	if err != nil {
		return Proxy{}, fmt.Errorf("Create proxy fail")
	}

	return p, nil
}

func (repo *ProxyRepository) Update(ctx context.Context, proxy Proxy) (Proxy, error) {

	q := "UPDATE proxies SET is_bad = $1, at_update = $2 WHERE id = $3"

	_, err := repo.client.Pool.Exec(ctx, q, &proxy.IsBad, &proxy.AtUpdate, &proxy.ID)
	if err != nil {
		return proxy, err
	}

	return proxy, nil
}

func (repo *ProxyRepository) Delete(ctx context.Context, proxyID uint64) error {

	q := "DELETE FROM proxies WHERE id=$1"

	_, err := repo.client.Pool.Exec(ctx, q, &proxyID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProxyRepository) Get(ctx context.Context, limit int) ([]Proxy, error) {

	cmd := "SELECT * FROM proxies"

	rows, err := repo.client.Pool.Query(ctx, cmd)
	if err != nil {
		return nil, err
	}

	proxies := make([]Proxy, 0)

	for rows.Next() {
		p := Proxy{}

		err = rows.Scan(&p.ID, &p.Scheme, &p.Host, &p.Port, &p.Username, &p.Password, &p.AtCreate, &p.AtUpdate, &p.IsBad)
		if err != nil {
			return nil, err
		}

		proxies = append(proxies, p)
	}

	return proxies, nil
}

func (repo *ProxyRepository) GetByID(ctx context.Context, proxyID uint64) (Proxy, error) {

	p := Proxy{}

	cmd := `SELECT (
		scheme, 
		host, 
		port, 
		username, 
		password, 
		at_create, 
		at_update, 
		is_bad
	) 
	FROM proxies 
	WHERE id = $1
	LIMIT 1`

	row := repo.client.Pool.QueryRow(ctx, cmd, proxyID)

	err := row.Scan(&p.ID, &p.Scheme, &p.Host, &p.Port, &p.Username, &p.Password, &p.AtCreate, &p.AtUpdate, &p.IsBad)

	return p, err
}

// func (repo *ProxyRepository) Get(ctx context.Context) (Proxy, error) {

// 	q := `SELECT * FROM proxies WHERE is_bad NOT $1 AND at_update NOT $2 LIMIT 1`

// 	at_update := time.Now()

// 	is_bad := false

// 	row := repo.client.Pool.QueryRow(ctx, q, &is_bad, &at_update)

// 	p := Proxy{
// 		AtUpdate: at_update,
// 		IsBad:    is_bad,
// 	}

// 	row.Scan(&p.ID, &p.Scheme, &p.Host, &p.Port, &p.Username, &p.Password, &p.AtCreate, &p.AtUpdate, &p.IsBad)

// 	return p, nil
// }
