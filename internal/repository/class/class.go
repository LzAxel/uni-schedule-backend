package class

import (
	"context"
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
		Columns("schedule_id", "subject_id", "teacher_id", "class_type", "day_of_week", "class_number", "even_week").
		Values(class.ScheduleID, class.SubjectID, class.TeacherID, class.ClassType, class.DayOfWeek, class.Number, class.EvenWeek).
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

func (r *ClassRepo) CreateOrSplit(class domain.CreateClassDTO) (uint64, error) {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	foundClasses, err := r.GetAllByDayAndNumber(class.ScheduleID, class.DayOfWeek, class.Number)
	if err != nil {
		return 0, err
	}

	if len(foundClasses) > 1 {
		return 0, apperror.ErrClassesAlreadySet
	}

	if len(foundClasses) == 1 {
		otherClass := foundClasses[0]
		if (otherClass.EvenWeek != nil &&
			class.EvenWeek != nil &&
			*otherClass.EvenWeek == *class.EvenWeek) || (otherClass.EvenWeek == nil && class.EvenWeek == nil) {
			return 0, apperror.ErrCannotSetSameClassesPositions
		}
		if class.EvenWeek == nil && otherClass.EvenWeek != nil {
			return 0, apperror.ErrClassOnlyEvenOrSingleClass
		}
		if class.EvenWeek != nil && otherClass.EvenWeek == nil {
			reversedEvenWeek := !*class.EvenWeek

			query, args, _ := r.psql.Update("classes").
				Where(squirrel.Eq{"id": otherClass.ID}).
				Set("even_week", reversedEvenWeek).ToSql()
			_, err = tx.Exec(query, args...)
			if err != nil {
				return 0, err
			}

		}
	}

	query, args, err := r.psql.Insert("classes").
		Columns("schedule_id", "subject_id", "teacher_id", "class_type", "day_of_week", "class_number", "even_week").
		Values(class.ScheduleID, class.SubjectID, class.TeacherID, class.ClassType, class.DayOfWeek, class.Number, class.EvenWeek).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	var id uint64
	err = tx.QueryRow(query, args...).Scan(&id)
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

func (r *ClassRepo) GetAllByDayAndNumber(scheduleID uint64, day domain.Day, number uint64) ([]domain.Class, error) {
	query, args, err := r.psql.Select("*").From("classes").
		Where(squirrel.Eq{"schedule_id": scheduleID, "day_of_week": day, "class_number": number}).
		ToSql()

	if err != nil {
		return []domain.Class{}, err
	}

	var classes []domain.Class
	err = r.db.Select(&classes, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Class{}, apperror.ErrNotFound
		}
		return []domain.Class{}, err
	}

	return classes, nil
}

type GetAllViewsStruct struct {
	ID         uint64           `db:"class_id"`
	ScheduleID uint64           `db:"schedule_id"`
	ClassType  domain.ClassType `db:"class_type"`
	DayOfWeek  domain.Day       `db:"day_of_week"`
	Number     uint64           `db:"class_number"`
	EvenWeek   *bool            `db:"even_week"`

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
		ID:        c.ID,
		DayOfWeek: c.DayOfWeek,
		Number:    c.Number,
		EvenWeek:  c.EvenWeek,
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
func (r *ClassRepo) GetAllViews(scheduleID uint64) (domain.ClassViews, uint64, error) {
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
	`
	if err := r.db.Get(&total, query, scheduleID); err != nil {
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
			c.day_of_week,	
			c.class_number,
			c.even_week,
		
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
		`
	err := r.db.Select(&classes, query, scheduleID)
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

	q = q.Set("subject_id", update.SubjectID)
	q = q.Set("teacher_id", update.TeacherID)
	q = q.Set("class_type", update.ClassType)
	q = q.Set("day_of_week", update.DayOfWeek)
	q = q.Set("class_number", update.Number)
	q = q.Set("even_week", update.EvenWeek)

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

func (r *ClassRepo) UpdateOrSwitch(id uint64, scheduleID uint64, update domain.UpdateClassDTO) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		tx.Rollback()
	}()

	q1 := r.psql.Update("classes").Where(squirrel.Eq{"id": id}).
		Set("subject_id", update.SubjectID).
		Set("teacher_id", update.TeacherID).
		Set("class_type", update.ClassType).
		Set("day_of_week", update.DayOfWeek).
		Set("class_number", update.Number).
		Set("even_week", update.EvenWeek)

	query1, args1, err := q1.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build update query: %w", err)
	}

	if _, err = tx.Exec(query1, args1...); err != nil {
		return fmt.Errorf("failed to update class: %w:%s:%+v", err, query1, args1)
	}

	foundClasses, err := r.GetAllByDayAndNumber(scheduleID, update.DayOfWeek, update.Number)
	if err != nil {
		return err
	}

	if len(foundClasses) > 1 {
		otherClass := foundClasses[0]
		if otherClass.ID == id {
			otherClass = foundClasses[1]
		}

		if update.EvenWeek == nil && otherClass.EvenWeek != nil {
			return apperror.ErrClassOnlyEvenOrSingleClass
		}

		if update.EvenWeek != nil && otherClass.EvenWeek == nil || ((update.EvenWeek != nil && otherClass.EvenWeek != nil) &&
			(*update.EvenWeek == *otherClass.EvenWeek)) {
			reversedEvenWeek := !*update.EvenWeek
			q2, args2, _ := r.psql.Update("classes").Where(squirrel.Eq{"id": otherClass.ID}).
				Set("even_week", reversedEvenWeek).ToSql()

			if _, err = tx.Exec(q2, args2...); err != nil {
				return fmt.Errorf("failed to update secondary class: %w", err)
			}
		}

	}
	if err := tx.Commit(); err != nil {
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
