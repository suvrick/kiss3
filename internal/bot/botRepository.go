package bot

import (
	"context"
	"fmt"

	"github.com/suvrick/kiss/pkg/db/client/postgres"
)

type BotRepository struct {
	client *postgres.Client
}

func NewBotRepository(db *postgres.Client) *BotRepository {
	return &BotRepository{
		client: db,
	}
}

func (repo *BotRepository) Create(ctx context.Context, p Bot) (Bot, error) {

	q := `INSERT INTO bots(
			game_id, 
			name, 
			balance, 
			avatar, 
			profile,
			at_create, 
			at_update 
		) 
		VALUES($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id`

	row := repo.client.Pool.QueryRow(ctx, q,
		&p.GameID,
		&p.Name,
		&p.Balance,
		&p.Avatar,
		&p.Profile,
		&p.AtCreate,
		&p.AtUpdate)

	err := row.Scan(&p.ID)

	if err != nil {
		return Bot{}, fmt.Errorf("Create bot fail")
	}

	return p, nil
}

func (repo *BotRepository) Update(ctx context.Context, bot Bot) (Bot, error) {

	q := `
			UPDATE 
				bots 
			SET 
				balance = $1, 
				avatar = $2,
				at_update = $3
			WHERE 
				id = $4
	`

	_, err := repo.client.Pool.Exec(ctx, q,
		&bot.Balance,
		&bot.Avatar,
		&bot.AtUpdate,
		&bot.ID)

	if err != nil {
		return bot, err
	}

	return bot, nil
}

func (repo *BotRepository) Delete(ctx context.Context, botID uint64) error {

	q := "DELETE FROM bots WHERE id=$1"

	_, err := repo.client.Pool.Exec(ctx, q, &botID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BotRepository) Get(ctx context.Context, limit int) ([]Bot, error) {

	cmd := "SELECT * FROM bots"

	rows, err := repo.client.Pool.Query(ctx, cmd)
	if err != nil {
		return nil, err
	}

	rows.Close()

	proxies := make([]Bot, 0)

	for rows.Next() {
		p := Bot{}

		err = rows.Scan(
			&p.ID,
			&p.GameID,
			&p.Name,
			&p.Balance,
			&p.Avatar,
			&p.Profile,
			&p.AtCreate,
			&p.AtUpdate,
		)

		if err != nil {
			return nil, err
		}

		proxies = append(proxies, p)
	}

	return proxies, nil
}

func (repo *BotRepository) GetByID(ctx context.Context, botID uint64) (Bot, error) {

	p := Bot{}

	cmd := `
	SELECT (
		game_id, 
		name, 
		balance, 
		avatar, 
		profile,
		at_create, 
		at_update, 
	) 
	FROM 
		proxies 
	WHERE 
		id = $1
	LIMIT 
		1`

	row := repo.client.Pool.QueryRow(ctx, cmd, botID)

	err := row.Scan(
		&p.ID,
		&p.GameID,
		&p.Name,
		&p.Balance,
		&p.Avatar,
		&p.Profile,
		&p.AtCreate,
		&p.AtUpdate,
	)

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

// 	row.Scan(&p.ID, &p.Scheme, &p.Host, &p.Port, &p.Botname, &p.Password, &p.AtCreate, &p.AtUpdate, &p.IsBad)

// 	return p, nil
// }
