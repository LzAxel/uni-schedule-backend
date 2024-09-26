package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/auth/model"
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

func (r *TokenRepo) CreateOrUpdate(token model.RefreshToken) error {
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
func (r *TokenRepo) GetByUserID(userID domain.ID) (model.RefreshToken, error) {
	query, args, err := r.psql.
		Select("user_id", "updated_at", "token").
		From("refresh_tokens").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return model.RefreshToken{}, err
	}

	var token model.RefreshToken
	err = r.db.Get(&token, query, args...)

	return token, err
}
func (r *TokenRepo) Delete(userID domain.ID) error {
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
