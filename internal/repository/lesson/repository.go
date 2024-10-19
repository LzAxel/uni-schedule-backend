package lesson

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
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

func (r *LessonRepo) Create(lesson domain.LessonCreate) (uint64, error) {
	query, args, err := r.psql.Insert("lessons").
		Columns("name", "location", "teacher_id", "lesson_type").
		Values(lesson.Name, lesson.Location, lesson.TeacherID, lesson.LessonType).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var id uint64
	err = r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *LessonRepo) GetByID(id uint64) (domain.Lesson, error) {
	query, args, err := r.psql.Select("*").From("lessons").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return domain.Lesson{}, err
	}

	var lesson domain.Lesson
	err = r.db.QueryRow(query, args...).Scan(&lesson.ID, &lesson.Name, &lesson.Location, &lesson.TeacherID, &lesson.LessonType)
	if err != nil {
		return domain.Lesson{}, err
	}
	return lesson, nil
}

func (r *LessonRepo) GetWithRelationsByID(id uint64) (domain.LessonView, error) {
	query, args, err := r.psql.
		Select(`
			l.id as lesson_id, 
			l.name as lesson_name, 
			l.location as lesson_location, 
			l.lesson_type as lesson_type, 
			t.id as teacher_id,
			t.short_name as teacher_short_name,
			t.full_name as teacher_full_name
		`).
		From("lessons as l").
		LeftJoin("teachers as t on l.teacher_id = t.id").
		Where(squirrel.Eq{"l.id": id}).
		ToSql()
	if err != nil {
		return domain.LessonView{}, err
	}

	var lesson lessonWithRelations
	err = r.db.Get(&lesson, query, args...)
	if err != nil {
		return domain.LessonView{}, err
	}
	return lesson.ToDomain(), nil
}

func (r *LessonRepo) Update(id uint64, update domain.LessonUpdate) error {
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

func (r *LessonRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("lessons").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
