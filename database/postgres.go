package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/cecardev/go-rest-server/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) (id int64, err error) {
	lastInsertId := 0
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	err = repo.db.QueryRowContext(ctx, query, user.Email, user.Password).Scan(&lastInsertId)

	if err != nil {
		log.Fatal(err)
	}
	return int64(lastInsertId), err

}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,  email FROM users WHERE id =$1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err != nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, err
}

func (repo *PostgresRepository) Close() error {
	return repo.Close()
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,  email, password FROM users WHERE email =$1", email)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err != nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, err
}
