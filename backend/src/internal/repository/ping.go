package repository

import (
	"context"
	"errors"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PingRepository struct {
	pool *pgxpool.Pool
}

var _ Ping = (*PingRepository)(nil)

func NewPingRepository(pool *pgxpool.Pool) *PingRepository {
	return &PingRepository{
		pool: pool,
	}
}

func (p *PingRepository) StorePings(ctx context.Context, pings []model.Ping) error {
	const query = `INSERT INTO pings
				   (container_id, ping_time, was_successful, date)
				   VALUES ($1, $2, $3, $4)`

	for i := range pings {
		_, err := p.pool.Exec(
			ctx,
			query,
			pings[i].GetContainerId(),
			pings[i].GetPingTime(),
			pings[i].GetWasSuccessful(),
			pings[i].GetDate(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PingRepository) GetPingsForContainer(
	ctx context.Context,
	offset int,
	containerID int,
) ([]model.Ping, error) {
	const limit = 20

	const query = `SELECT container_id, ping_time, was_successful, date
				   FROM pings 
				   WHERE container_id = $1
				   ORDER BY ping_time DESC
				   LIMIT $2
				   OFFSET $3`

	rows, err := p.pool.Query(ctx, query, containerID, limit, offset)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.Ping{}, nil
		}

		return nil, err
	}

	pings := make([]model.Ping, limit)
	i := 0

	for rows.Next() {
		err = rows.Scan(&pings[i].ContainerId, &pings[i].PingTime, &pings[i].WasSuccessful, &pings[i].Date)
		if err != nil {
			return nil, err
		}
	}

	return pings, nil
}
