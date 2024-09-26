package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/schedule/model"
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

func (r *ScheduleRepo) Create(schedule model.Schedule) (domain.ID, error) {
	query, args, err := r.psql.Insert("schedules").
		Columns("creator_id", "name", "slug").
		Values(schedule.CreatorID, schedule.Name, schedule.Slug).
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

func (r *ScheduleRepo) GetByID(id domain.ID) (model.Schedule, error) {
	query, args, err := r.psql.Select("*").From("schedules").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return model.Schedule{}, err
	}

	var schedule model.Schedule
	err = r.db.QueryRow(query, args...).Scan(&schedule.ID, &schedule.CreatorID, &schedule.Name, &schedule.Slug)
	if err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (r *ScheduleRepo) Update(id domain.ID, update model.ScheduleUpdate) error {
	q := r.psql.Update("schedules").Where(squirrel.Eq{"id": id})

	if update.CreatorID != nil {
		q = q.Set("creator_id", *update.CreatorID)
	}
	if update.Name != nil {
		q = q.Set("name", *update.Name)
	}
	if update.Slug != nil {
		q = q.Set("slug", *update.Slug)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *ScheduleRepo) Delete(id domain.ID) error {
	query, args, err := r.psql.Delete("schedules").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
