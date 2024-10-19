package auth

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
)

type TokenRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewTokenRepo(db *sqlx.DB) *TokenRepo {
	return &TokenRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *TokenRepo) CreateOrUpdate(token domain.RefreshToken) error {
	query, args, err := r.psql.
		Insert("refresh_tokens").
		Columns("user_id", "token", "updated_at").
		Values(token.UserID, token.RefreshToken, token.UpdatedAt).
		Suffix("ON CONFLICT (user_id) DO UPDATE SET token = ?", token.RefreshToken).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
func (r *TokenRepo) GetByUserID(userID uint64) (domain.RefreshToken, error) {
	query, args, err := r.psql.
		Select("user_id", "updated_at", "token").
		From("refresh_tokens").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return domain.RefreshToken{}, err
	}

	var token domain.RefreshToken
	err = r.db.Get(&token, query, args...)

	return token, err
}
func (r *TokenRepo) Delete(userID uint64) error {
	query, args, err := r.psql.
		Delete("refresh_tokens").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)

	return err
}
