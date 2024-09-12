package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/user/model"
)

type UserRepo struct {
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *UserRepo) Create(user model.User) (domain.ID, error) {
	query, args, err := r.psql.Insert("users").
		Columns("username", "password_hash", "role", "created_at").
		Values(user.Username, user.PasswordHash, user.Role, user.CreatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var id domain.ID
	err = r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) GetByID(id domain.ID) (model.User, error) {
	query, args, err := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetByUsername(username string) (model.User, error) {
	query, args, err := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) Update(id domain.ID, update model.UserUpdateDTO) error {
	q := r.psql.Update("users").Where(squirrel.Eq{"id": id})

	if update.Role != nil {
		q = q.Set("role", *update.Role)
	}
	if update.Username != nil {
		q = q.Set("username", *update.Username)
	}
	if update.PasswordHash != nil {
		q = q.Set("password_hash", *update.PasswordHash)
	}
	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *UserRepo) Delete(id domain.ID) error {
	query, args, err := r.psql.Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
