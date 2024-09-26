package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/lesson/model"
)

type LessonRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewLessonRepo(db *sqlx.DB) *LessonRepo {
	return &LessonRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *LessonRepo) Create(lesson model.Lesson) (domain.ID, error) {
	query, args, err := r.psql.Insert("lessons").
		Columns("name", "location", "teacher_id", "lesson_type").
		Values(lesson.Name, lesson.Location, lesson.TeacherID, lesson.LessonType).
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

func (r *LessonRepo) GetByID(id domain.ID) (model.Lesson, error) {
	query, args, err := r.psql.Select("*").From("lessons").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return model.Lesson{}, err
	}

	var lesson model.Lesson
	err = r.db.QueryRow(query, args...).Scan(&lesson.ID, &lesson.Name, &lesson.Location, &lesson.TeacherID, &lesson.LessonType)
	if err != nil {
		return model.Lesson{}, err
	}
	return lesson, nil
}

func (r *LessonRepo) Update(id domain.ID, update model.LessonUpdate) error {
	q := r.psql.Update("lessons").Where(squirrel.Eq{"id": id})

	if update.Name != nil {
		q = q.Set("name", *update.Name)
	}
	if update.Location != nil {
		q = q.Set("location", *update.Location)
	}
	if update.TeacherID != nil {
		q = q.Set("teacher_id", *update.TeacherID)
	}
	if update.LessonType != nil {
		q = q.Set("lesson_type", *update.LessonType)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *LessonRepo) Delete(id domain.ID) error {
	query, args, err := r.psql.Delete("lessons").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
