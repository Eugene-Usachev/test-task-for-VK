package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContainerRepository struct {
	pool *pgxpool.Pool
}

var _ Container = &ContainerRepository{}

func NewContainerRepository(pool *pgxpool.Pool) *ContainerRepository {
	return &ContainerRepository{
		pool: pool,
	}
}

func (c *ContainerRepository) RegisterContainer(ctx context.Context, container *model.RegisterContainer) error {
	const query = `INSERT INTO containers
    			   (ip_address)
				   VALUES ($1)
				   ON CONFLICT DO NOTHING`

	_, err := c.pool.Exec(ctx, query, container.GetIpAddress())

	return err
}

func (c *ContainerRepository) GetContainers(ctx context.Context) ([]model.GetContainer, error) {
	const query = `SELECT id, ip_address FROM containers`

	containers := make([]model.GetContainer, 0, 8)
	i := 0

	rows, err := c.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		containers = append(containers, model.GetContainer{})

		err = rows.Scan(&containers[i].Id, &containers[i].IpAddress)
		if err != nil {
			return nil, err
		}

		i++
	}

	return containers, nil
}

func (c *ContainerRepository) GetContainersWithLatestPing(
	ctx context.Context,
) (successfulContainers []model.GetContainerWithLatestPing, invalidContainers []model.GetContainer, err error) {
	const query = `SELECT DISTINCT ON (c.id) 
					       c.id,
					       c.ip_address,
					       p.ping_time,
					       p.date
					FROM containers c
					LEFT JOIN pings p ON c.id = p.container_id AND p.was_successful = TRUE
					ORDER BY c.id, p.date DESC`

	rows, err := c.pool.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return
		}

		return
	}

	var (
		successSliceIndex = 0
		containers        = make([]model.GetContainerWithLatestPing, 8)
		pingTime          sql.NullInt64
		date              sql.NullTime
	)

	for rows.Next() {
		if successSliceIndex == len(containers) {
			containers = append(containers, containers...)
		}

		err = rows.Scan(
			&containers[successSliceIndex].Id,
			&containers[successSliceIndex].IpAddress,
			&pingTime,
			&date,
		)
		if err != nil {
			return containers[:successSliceIndex], invalidContainers, err
		}

		if pingTime.Valid {
			containers[successSliceIndex].PingTime = pingTime.Int64
			containers[successSliceIndex].Date = timestamppb.New(date.Time)

			successSliceIndex++
		} else {
			invalidContainers = append(
				invalidContainers,
				model.GetContainer{
					Id:        containers[successSliceIndex].GetId(),
					IpAddress: containers[successSliceIndex].GetIpAddress(),
				},
			)
		}
	}

	return containers[:successSliceIndex], invalidContainers, nil
}

func (c *ContainerRepository) UnregisterContainer(ctx context.Context, containerID int) error {
	const query = `DELETE FROM containers
				   WHERE id = $1`

	_, err := c.pool.Exec(ctx, query, containerID)

	return err
}
