package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/user/model"
	"uni-schedule-backend/pkg/psql"
)

type UserRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *UserRepo) Create(user model.UserCreate) (domain.ID, error) {
	query, args := r.psql.Insert("users").
		Columns("username", "password_hash", "role", "created_at").
		Values(user.Username, user.PasswordHash, user.Role, user.CreatedAt).
		Suffix("RETURNING id").
		MustSql()

	var id domain.ID
	err := r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		if psql.IsPgErrorCode(err, pgerrcode.UniqueViolation) {
			return 0, apperror.ErrAlreadyExists
		}
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) GetByID(id domain.ID) (model.User, error) {
	query, args := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"id": id}).
		MustSql()

	var user model.User
	err := r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetByUsername(username string) (model.User, error) {
	query, args := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"username": username}).
		MustSql()

	var user model.User
	err := r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		if psql.IsNoRows(err) {
			return model.User{}, apperror.ErrNotFound
		}
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
	query, args := q.MustSql()

	_, err := r.db.Exec(query, args...)
	if psql.IsPgErrorCode(err, pgerrcode.UniqueViolation) {
		return apperror.ErrAlreadyExists
	}

	return err
}

func (r *UserRepo) Delete(id domain.ID) error {
	query, args := r.psql.Delete("users").
		Where(squirrel.Eq{"id": id}).
		MustSql()

	_, err := r.db.Exec(query, args...)
	return err
}
