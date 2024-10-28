package class

import (
	"database/sql"
	"errors"
	"fmt"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ClassRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewClassRepo(db *sqlx.DB) *ClassRepo {
	return &ClassRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *ClassRepo) Create(class domain.CreateClassDTO) (uint64, error) {
	query, args, err := r.psql.Insert("classes").
		Columns("schedule_id", "subject_id", "teacher_id", "class_type").
		Values(class.ScheduleID, class.SubjectID, class.TeacherID, class.ClassType).
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

func (r *ClassRepo) GetByID(id uint64) (domain.Class, error) {
	query, args, err := r.psql.Select("*").From("classes").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return domain.Class{}, err
	}

	var class domain.Class
	err = r.db.Get(&class, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Class{}, apperror.ErrNotFound
		}
		return domain.Class{}, err
	}

	return class, nil
}

type GetAllViewsStruct struct {
	ID         uint64           `db:"class_id"`
	ScheduleID uint64           `db:"schedule_id"`
	ClassType  domain.ClassType `db:"class_type"`

	// Subject details
	SubjectID   uint64 `db:"subject_id"`
	SubjectName string `db:"subject_name"`

	// Teacher details
	TeacherID uint64 `db:"teacher_id"`
	FirstName string `db:"teacher_first_name"`
	LastName  string `db:"teacher_last_name"`
	Surname   string `db:"teacher_surname"`
}

func (c GetAllViewsStruct) ToView() domain.ClassView {
	return domain.ClassView{
		ID: c.ID,
		Subject: domain.SubjectView{
			ID:   c.SubjectID,
			Name: c.SubjectName,
		},
		Teacher: domain.TeacherView{
			ID:        c.TeacherID,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Surname:   c.Surname,
		},
		ClassType: c.ClassType,
	}
}
func (r *ClassRepo) GetAllViews(scheduleID uint64, limit uint64, offset uint64) ([]domain.ClassView, uint64, error) {
	var (
		classes []GetAllViewsStruct = make([]GetAllViewsStruct, 0)
		total   uint64
	)

	query := `
		SELECT COUNT(c.*) OVER() AS total
		
		FROM classes c
		JOIN subjects s ON c.subject_id = s.id
		JOIN teachers t ON c.teacher_id = t.id
		
		WHERE c.schedule_id = $1
		LIMIT $2
		OFFSET $3;
	`
	if err := r.db.Get(&total, query, scheduleID, limit, offset); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			total = 0
		} else {
			return []domain.ClassView{}, total, err
		}
	}

	query = `
		SELECT
			c.id AS class_id,
			c.schedule_id,
			c.class_type,
		
			-- Subject information
			s.id AS subject_id,
			s.name AS subject_name,
		
			-- Teacher information
			t.id AS teacher_id,
			t.first_name AS teacher_first_name,
			t.last_name AS teacher_last_name,
			t.surname AS teacher_surname
		
		FROM classes c
		JOIN subjects s ON c.subject_id = s.id
		JOIN teachers t ON c.teacher_id = t.id
		
		WHERE c.schedule_id = $1
		ORDER BY c.id DESC
		LIMIT $2
		OFFSET $3;
		`
	err := r.db.Select(&classes, query, scheduleID, limit, offset)
	if err != nil {
		return []domain.ClassView{}, total, fmt.Errorf("getting all classes views: %w", err)
	}

	classViews := make([]domain.ClassView, len(classes))
	for i := range classes {
		classViews[i] = classes[i].ToView()
	}

	return classViews, total, nil
}

func (r *ClassRepo) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Class, uint64, error) {
	var (
		classes []domain.Class = make([]domain.Class, 0)
		total   uint64
	)

	countQuery, countArgs, err := r.psql.Select("COUNT(*)").
		From("classes").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		ToSql()
	if err != nil {
		return classes, total, err
	}

	if err := r.db.Get(&total, countQuery, countArgs...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return classes, total, apperror.ErrNotFound
		}
		return classes, total, err
	}

	query, args, err := r.psql.Select("*").From("classes").
		Where(squirrel.Eq{"schedule_id": scheduleID}).
		Limit(limit).
		Offset(offset).
		OrderBy("id DESC").
		ToSql()

	if err != nil {
		return classes, total, err
	}

	err = r.db.Select(&classes, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return classes, total, apperror.ErrNotFound
		}
		return classes, total, err
	}

	return classes, total, nil
}

func (r *ClassRepo) Update(id uint64, update domain.UpdateClassDTO) error {
	q := r.psql.Update("classes").Where(squirrel.Eq{"id": id})

	if update.SubjectID != nil {
		q = q.Set("subject_id", *update.SubjectID)
	}
	if update.TeacherID != nil {
		q = q.Set("teacher_id", *update.TeacherID)
	}
	if update.ClassType != nil {
		q = q.Set("class_type", *update.ClassType)
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

func (r *ClassRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("classes").
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
