package repository

import (
	"context"
	"errors"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
				   VALUES ($1)`

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
) ([]model.GetContainerWithLatestPing, error) {
	const query = `SELECT 
    			   c.id, c.ip_address, p.ping_time, p.was_successful, p.date
				   FROM containers c 
				   LEFT JOIN pings p ON c.id = p.container_id
				   ORDER BY p.ping_time DESC`

	containers := make([]model.GetContainerWithLatestPing, 0, 8)
	i := 0

	rows, err := c.pool.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []model.GetContainerWithLatestPing{}, nil
		}

		return nil, err
	}

	for rows.Next() {
		containers = append(containers, model.GetContainerWithLatestPing{})

		err = rows.Scan(
			&containers[i].Id,
			&containers[i].IpAddress,
			&containers[i].PingTime,
			&containers[i].WasSuccessful,
			&containers[i].Date,
		)
		if err != nil {
			return nil, err
		}

		i++
	}

	return containers, nil
}

func (c *ContainerRepository) UnregisterContainer(ctx context.Context, containerID int) error {
	const query = `DELETE FROM containers
				   WHERE id = $1`

	_, err := c.pool.Exec(ctx, query, containerID)

	return err
}
