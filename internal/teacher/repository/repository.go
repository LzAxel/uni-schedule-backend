package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/teacher/model"
)

type TeacherRepo struct {
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

func NewTeacherRepo(db *sql.DB) *TeacherRepo {
	return &TeacherRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *TeacherRepo) Create(teacher model.Teacher) (domain.ID, error) {
	query, args, err := r.psql.Insert("teachers").
		Columns("short_name", "full_name").
		Values(teacher.ShortName, teacher.FullName).
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

func (r *TeacherRepo) GetByID(id domain.ID) (model.Teacher, error) {
	query, args, err := r.psql.Select("*").From("teachers").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return model.Teacher{}, err
	}

	var teacher model.Teacher
	err = r.db.QueryRow(query, args...).Scan(&teacher.ID, &teacher.ShortName, &teacher.FullName)
	if err != nil {
		return model.Teacher{}, err
	}
	return teacher, nil
}

func (r *TeacherRepo) GetAll() ([]model.Teacher, error) {
	query, args, err := r.psql.Select("*").From("teachers").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []model.Teacher
	for rows.Next() {
		var teacher model.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.ShortName, &teacher.FullName); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func (r *TeacherRepo) Update(id domain.ID, update model.TeacherUpdate) error {
	q := r.psql.Update("teachers").Where(squirrel.Eq{"id": id})

	if update.ShortName != nil {
		q = q.Set("short_name", *update.ShortName)
	}
	if update.FullName != nil {
		q = q.Set("full_name", *update.FullName)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *TeacherRepo) Delete(id domain.ID) error {
	query, args, err := r.psql.Delete("teachers").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
