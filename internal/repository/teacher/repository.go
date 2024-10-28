package teacher

import (
	"database/sql"
	"errors"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type TeacherRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewTeacherRepo(db *sqlx.DB) *TeacherRepo {
	return &TeacherRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *TeacherRepo) Create(teacher domain.TeacherCreateDTO) (uint64, error) {
	var id uint64
	query, args, err := r.psql.Insert("teachers").
		Columns("first_name", "last_name", "surname", "schedule_id").
		Values(teacher.FirstName, teacher.LastName, teacher.Surname, teacher.ScheduleID).
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

func (r *TeacherRepo) GetByID(id uint64) (domain.Teacher, error) {
	var teacher domain.Teacher
	query, args, err := r.psql.Select("*").
		From("teachers").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return domain.Teacher{}, err
	}

	if err := r.db.Get(&teacher, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Teacher{}, apperror.ErrNotFound
		}
		return domain.Teacher{}, err
	}

	return teacher, nil
}

func (r *TeacherRepo) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Teacher, uint64, error) {
	var (
		teachers []domain.Teacher = make([]domain.Teacher, 0)
		total    uint64
	)
	countQuery, countArgs, err := r.psql.Select("COUNT(*)").
		From("teachers").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		ToSql()
	if err != nil {
		return teachers, total, err
	}

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teachers, total, apperror.ErrNotFound
		}
		return teachers, total, err
	}

	query, args, err := r.psql.Select("*").
		From("teachers").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		Limit(limit).
		Offset(offset).
		OrderBy("id DESC").
		ToSql()

	if err != nil {
		return teachers, total, err
	}

	if err := r.db.Select(&teachers, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teachers, total, apperror.ErrNotFound
		}
		return teachers, total, err
	}

	return teachers, total, nil
}

func (r *TeacherRepo) Update(id uint64, update domain.TeacherUpdateDTO) error {
	q := r.psql.Update("teachers").Where(squirrel.Eq{"id": id})

	if update.FirstName != nil {
		q = q.Set("first_name", *update.FirstName)
	}
	if update.LastName != nil {
		q = q.Set("last_name", *update.LastName)
	}
	if update.Surname != nil {
		q = q.Set("surname", *update.Surname)
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

func (r *TeacherRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("teachers").
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
