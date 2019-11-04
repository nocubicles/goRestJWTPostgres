package user

import (
	"context"
	"database/sql"
	"goRestJWTPostgres/db"
	"goRestJWTPostgres/middleware/errors"
	"goRestJWTPostgres/models"
	"goRestJWTPostgres/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type pgsqlUserRepo struct {
	Conn *sql.DB
}

func NewSQLUserRepo(Conn *sql.DB) repository.UserRepo {
	return &pgsqlUserRepo{
		Conn: db.Database,
	}
}

func (m *pgsqlUserRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.User, 0)

	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.ID,
			&data.FirstName,
			&data.LastName,
		)
		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func (m *pgsqlUserRepo) Fetch(ctx context.Context, num int64) ([]*models.User, error) {
	query := "Select id, first_name, last_name, email, city From users limit ?"

	return m.fetch(ctx, query, num)
}

func (m *pgsqlUserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := "Select id, title, content From users where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}

	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, errors.ErrNotFound
	}

	return payload, nil
}

func (m *pgsqlUserRepo) Create(ctx context.Context, u *models.User) (int64, error) {
	var err error

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Passwordhash), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	query := "INSERT INTO users(first_name, last_name, email, passwordHash) VALUES ($1, $2, $3, $4) RETURNING id"

	stmt, err := m.Conn.Prepare(query)

	if err != nil {
		return -1, err
	}

	var id int64

	err = stmt.QueryRow(u.FirstName, u.LastName, u.Email, passwordHash).Scan(&id)

	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *pgsqlUserRepo) Update(ctx context.Context, u *models.User) (*models.User, error) {
	query := "Update users set first_name=?, last_name=? city=? email=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		u.FirstName,
		u.LastName,
		u.Email,
	)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return u, nil
}

func (m *pgsqlUserRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From users Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		return false, err
	}
	return true, nil
}
