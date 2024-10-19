package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
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

func (r *UserRepo) Create(user domain.UserCreate) (uint64, error) {
	query, args := r.psql.Insert("users").
		Columns("username", "password_hash", "role", "created_at").
		Values(user.Username, user.PasswordHash, user.Role, user.CreatedAt).
		Suffix("RETURNING id").
		MustSql()

	var id uint64
	err := r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		if psql.IsPgErrorCode(err, pgerrcode.UniqueViolation) {
			return 0, apperror.ErrAlreadyExists
		}
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) GetByID(id uint64) (domain.User, error) {
	query, args := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"id": id}).
		MustSql()

	var user domain.User
	err := r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetByUsername(username string) (domain.User, error) {
	query, args := r.psql.Select("*").From("users").
		Where(squirrel.Eq{"username": username}).
		MustSql()

	var user domain.User
	err := r.db.QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		if psql.IsNoRows(err) {
			return domain.User{}, apperror.ErrNotFound
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepo) Update(id uint64, update domain.UserUpdateDTO) error {
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

func (r *UserRepo) Delete(id uint64) error {
	query, args := r.psql.Delete("users").
		Where(squirrel.Eq{"id": id}).
		MustSql()

	_, err := r.db.Exec(query, args...)
	return err
}
