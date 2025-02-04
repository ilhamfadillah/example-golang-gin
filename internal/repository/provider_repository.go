package repository

import (
	"database/sql"
	"example/apps/config"
	"example/apps/internal/model"
	"time"

	"github.com/google/uuid"
)

type ProviderRepository interface {
	GetAll() ([]*model.Provider, error)
	GetByUuid(uuid string) (*model.Provider, error)
	Store(provider *model.Provider) (*model.Provider, error)
	Update(uuid string, provider *model.Provider) (*model.Provider, error)
}

type providerRepository struct {
	db *sql.DB
}

func NewProviderRepository(db *sql.DB) ProviderRepository {
	return &providerRepository{db: db}
}

func (repo *providerRepository) GetAll() ([]*model.Provider, error) {
	rows, err := config.DB.Query("SELECT * FROM providers")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var providers []*model.Provider

	for rows.Next() {
		var provider = &model.Provider{}
		if err := rows.Scan(
			&provider.ID,
			&provider.UUID,
			&provider.Code,
			&provider.Name,
			&provider.Address,
			&provider.Phone,
			&provider.City,
			&provider.CreatedAt,
			&provider.UpdatedAt,
			&provider.DeletedAt,
		); err != nil {
			return nil, err
		}

		providers = append(providers, provider)

	}

	return providers, nil
}

func (repo *providerRepository) GetByUuid(uuid string) (*model.Provider, error) {
	var provider model.Provider

	err := config.DB.QueryRow("SELECT * FROM providers where uuid = ?", uuid).Scan(
		&provider.ID,
		&provider.UUID,
		&provider.Code,
		&provider.Name,
		&provider.Address,
		&provider.Phone,
		&provider.City,
		&provider.CreatedAt,
		&provider.UpdatedAt,
		&provider.DeletedAt,
	)

	return &provider, err
}

func (repo *providerRepository) Store(provider *model.Provider) (*model.Provider, error) {
	uuid := uuid.New().String()
	now := time.Now()

	query := "INSERT INTO providers (uuid, code, name, address, phone, city, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, uuid, provider.Code, provider.Name, provider.Address, provider.Phone, provider.City, now, now)
	if err != nil {
		return nil, err
	}

	return repo.GetByUuid(uuid)
}

func (repo *providerRepository) Update(uuid string, provider *model.Provider) (*model.Provider, error) {
	now := time.Now()

	query := "UPDATE providers SET code = ?, name = ?, address = ?, phone = ?, city = ?, updated_at = ? WHERE uuid = ?"
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(provider.Code, provider.Name, provider.Address, provider.Phone, provider.City, now, uuid)
	if err != nil {
		return nil, err
	}

	return repo.GetByUuid(uuid)
}
