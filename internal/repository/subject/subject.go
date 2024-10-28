package subject

import (
	"database/sql"
	"errors"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SubjectRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewSubjectRepo(db *sqlx.DB) *SubjectRepo {
	return &SubjectRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *SubjectRepo) Create(subject domain.CreateSubjectDTO) (uint64, error) {
	var id uint64
	query, args, err := r.psql.Insert("subjects").
		Columns("name", "schedule_id").
		Values(subject.Name, subject.ScheduleID).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	if err := r.db.QueryRow(query, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SubjectRepo) GetByID(id uint64) (domain.Subject, error) {
	var subject domain.Subject
	query, args, err := r.psql.Select("*").
		From("subjects").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return domain.Subject{}, err
	}

	if err := r.db.Get(&subject, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Subject{}, apperror.ErrNotFound
		}
		return domain.Subject{}, err
	}

	return subject, nil
}
func (r *SubjectRepo) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Subject, uint64, error) {
	var (
		subjects []domain.Subject = make([]domain.Subject, 0)
		total    uint64
	)

	countQuery, countArgs, err := r.psql.Select("COUNT(*)").
		From("subjects").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		ToSql()
	if err != nil {
		return subjects, total, err
	}

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subjects, total, apperror.ErrNotFound
		}
		return subjects, total, err
	}

	query, args, err := r.psql.Select("*").
		From("subjects").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		Limit(limit).
		Offset(offset).
		OrderBy("id DESC").
		ToSql()

	if err != nil {
		return subjects, total, err
	}

	if err := r.db.Select(&subjects, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subjects, total, apperror.ErrNotFound
		}
		return subjects, total, err
	}

	return subjects, total, nil
}

func (r *SubjectRepo) Update(id uint64, update domain.UpdateSubjectDTO) error {
	q := r.psql.Update("subjects").Where(squirrel.Eq{"id": id})

	if update.Name != nil {
		q = q.Set("name", *update.Name)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func (r *SubjectRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("subjects").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
