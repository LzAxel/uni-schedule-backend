package schedule

import (
	"database/sql"
	"errors"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ScheduleRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewScheduleRepo(db *sqlx.DB) *ScheduleRepo {
	return &ScheduleRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *ScheduleRepo) Create(schedule domain.CreateScheduleDTO) (uint64, error) {
	var id uint64
	query, args, err := r.psql.Insert("schedules").
		Columns("slug", "user_id").
		Values(schedule.Slug, schedule.UserID).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	err = r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, apperror.ErrNotFound
		}
		return 0, err
	}

	return id, nil
}

func (r *ScheduleRepo) GetByID(id uint64) (domain.Schedule, error) {
	var schedule domain.Schedule
	query, args, err := r.psql.Select("*").
		From("schedules").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return domain.Schedule{}, err
	}

	err = r.db.Get(&schedule, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Schedule{}, apperror.ErrNotFound
		}
		return domain.Schedule{}, err
	}

	return schedule, nil
}

func (r *ScheduleRepo) GetBySlug(slug string) (domain.Schedule, error) {
	var schedule domain.Schedule
	query, args, err := r.psql.Select("*").
		From("schedules").
		Where(squirrel.Eq{"slug": slug}).
		ToSql()

	if err != nil {
		return domain.Schedule{}, err
	}

	err = r.db.Get(&schedule, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Schedule{}, apperror.ErrNotFound
		}
		return domain.Schedule{}, err
	}

	return schedule, nil
}

func (r *ScheduleRepo) GetAll(limit uint64, offset uint64, filters domain.ScheduleGetAllFilters) ([]domain.Schedule, uint64, error) {
	var (
		schedules []domain.Schedule = make([]domain.Schedule, 0)
		total     uint64
	)

	countBuilder := r.psql.Select("COUNT(*)").From("schedules")

	if filters.UserID != nil {
		countBuilder = countBuilder.Where(squirrel.Eq{"user_id": *filters.UserID})
	}

	countQuery, countArgs, err := countBuilder.ToSql()
	if err != nil {
		return schedules, total, err
	}

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return schedules, total, apperror.ErrNotFound
		}
		return schedules, total, err
	}

	queryBuilder := r.psql.Select("*")

	if filters.UserID != nil {
		queryBuilder = queryBuilder.Where(squirrel.Eq{"user_id": *filters.UserID})
	}

	query, args, err := queryBuilder.From("schedules").
		Limit(limit).
		Offset(offset).
		OrderBy("id DESC").
		ToSql()

	if err != nil {
		return schedules, total, err
	}

	err = r.db.Select(&schedules, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return schedules, total, apperror.ErrNotFound
		}
		return schedules, total, err
	}

	return schedules, total, nil
}

func (r *ScheduleRepo) Update(id uint64, update domain.UpdateScheduleDTO) error {
	q := r.psql.Update("schedules").Where(squirrel.Eq{"id": id})

	if update.Slug != nil {
		q = q.Set("slug", *update.Slug)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ScheduleRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("schedules").
		Where(squirrel.Eq{"id": id}).
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
