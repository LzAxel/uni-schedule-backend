package entry

import (
	"database/sql"
	"errors"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type EntryRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewEntryRepo(db *sqlx.DB) *EntryRepo {
	return &EntryRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *EntryRepo) Create(entry domain.CreateScheduleEntryDTO) (uint64, error) {
	var id uint64
	query, args, err := r.psql.Insert("schedule_entries").
		Columns("day",
			"class_number",
			"even_class_id",
			"odd_class_id",
			"is_static",
			"schedule_id").
		Values(entry.Day, entry.ClassNumber, entry.EvenClassID, entry.OddClassID, entry.IsStatic, entry.ScheduleID).
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

func (r *EntryRepo) GetByID(id uint64) (domain.ScheduleEntry, error) {
	var entry domain.ScheduleEntry
	query, args, err := r.psql.Select("*").
		From("schedule_entries").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return domain.ScheduleEntry{}, err
	}

	if err := r.db.Get(&entry, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ScheduleEntry{}, apperror.ErrNotFound
		}
		return domain.ScheduleEntry{}, err
	}

	return entry, nil
}

type GetEntriesViewStruct struct {
	ID          uint64     `db:"schedule_entry_id"`
	ScheduleID  uint64     `db:"schedule_id"`
	Day         domain.Day `db:"day"`
	ClassNumber int        `db:"class_number"`
	IsStatic    bool       `db:"is_static"`

	// Even class information
	EvenClassID          *uint64           `db:"even_class_id"`
	EvenClassType        *domain.ClassType `db:"even_class_type"`
	EvenSubjectID        *uint64           `db:"even_subject_id"`
	EvenSubjectName      *string           `db:"even_subject_name"`
	EvenTeacherID        *uint64           `db:"even_teacher_id"`
	EvenTeacherFirstName *string           `db:"even_teacher_first_name"`
	EvenTeacherLastName  *string           `db:"even_teacher_last_name"`
	EvenTeacherSurname   *string           `db:"even_teacher_surname"`

	// Odd class information
	OddClassID          *uint64           `db:"odd_class_id"`
	OddClassType        *domain.ClassType `db:"odd_class_type"`
	OddSubjectID        *uint64           `db:"odd_subject_id"`
	OddSubjectName      *string           `db:"odd_subject_name"`
	OddTeacherID        *uint64           `db:"odd_teacher_id"`
	OddTeacherFirstName *string           `db:"odd_teacher_first_name"`
	OddTeacherLastName  *string           `db:"odd_teacher_last_name"`
	OddTeacherSurname   *string           `db:"odd_teacher_surname"`
}

func (g *GetEntriesViewStruct) ToView() domain.ScheduleEntryView {
	entry := domain.ScheduleEntryView{
		ID:          g.ID,
		Day:         g.Day,
		ClassNumber: g.ClassNumber,
		IsStatic:    g.IsStatic,
	}

	if g.EvenClassID != nil {
		entry.Even = &domain.ClassView{
			ID:        *g.EvenClassID,
			ClassType: *g.EvenClassType,
			Subject: domain.SubjectView{
				ID:   *g.EvenSubjectID,
				Name: *g.EvenSubjectName,
			},
			Teacher: domain.TeacherView{
				ID:        *g.EvenTeacherID,
				FirstName: *g.EvenTeacherFirstName,
				LastName:  *g.EvenTeacherLastName,
				Surname:   *g.EvenTeacherSurname,
			},
		}
	}

	if g.OddClassID != nil {
		entry.Odd = &domain.ClassView{
			ID:        *g.OddClassID,
			ClassType: *g.OddClassType,
			Subject: domain.SubjectView{
				ID:   *g.OddSubjectID,
				Name: *g.OddSubjectName,
			},
			Teacher: domain.TeacherView{
				ID:        *g.OddTeacherID,
				FirstName: *g.OddTeacherFirstName,
				LastName:  *g.OddTeacherLastName,
				Surname:   *g.OddTeacherSurname,
			},
		}
	}

	return entry
}

func (r *EntryRepo) GetEntriesView(scheduleID uint64) ([]domain.ScheduleEntryView, error) {
	var entries []GetEntriesViewStruct = make([]GetEntriesViewStruct, 0)

	query := `
		SELECT
		se.id AS schedule_entry_id,
		se.schedule_id,
		se.day,
		se.class_number,
		se.is_static,
	
		-- Even class information
		ec.id AS even_class_id,
		ec.class_type AS even_class_type,
		es.id AS even_subject_id,
		es.name AS even_subject_name,
		et.id AS even_teacher_id,
		et.first_name AS even_teacher_first_name,
		et.last_name AS even_teacher_last_name,
		et.surname AS even_teacher_surname,
	
		-- Odd class information
		oc.id AS odd_class_id,
		oc.class_type AS odd_class_type,
		os.id AS odd_subject_id,
		os.name AS odd_subject_name,
		ot.id AS odd_teacher_id,
		ot.first_name AS odd_teacher_first_name,
		ot.last_name AS odd_teacher_last_name,
		ot.surname AS odd_teacher_surname
	
	FROM schedule_entries se
	
	-- Join for even class
			 LEFT JOIN classes ec ON ec.id = se.even_class_id
			 LEFT JOIN subjects es ON es.id = ec.subject_id
			 LEFT JOIN teachers et ON et.id = ec.teacher_id
	
	-- Join for odd class
			 LEFT JOIN classes oc ON oc.id = se.odd_class_id
			 LEFT JOIN subjects os ON os.id = oc.subject_id
			 LEFT JOIN teachers ot ON ot.id = oc.teacher_id
	
	WHERE se.schedule_id = $1
	ORDER BY se.id DESC;
`
	if err := r.db.Select(&entries, query, scheduleID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.ScheduleEntryView{}, apperror.ErrNotFound
		}
		return []domain.ScheduleEntryView{}, err
	}

	entriesView := make([]domain.ScheduleEntryView, len(entries))
	for i := range entries {
		entriesView[i] = entries[i].ToView()
	}

	return entriesView, nil
}

func (r *EntryRepo) Update(id uint64, update domain.UpdateScheduleEntryDTO) error {
	q := r.psql.Update("schedule_entries").Where(squirrel.Eq{"id": id})

	if update.Day != nil {
		q = q.Set("day", *update.Day)
	}
	if update.ClassNumber != nil {
		q = q.Set("class_number", *update.ClassNumber)
	}
	if update.EvenClassID != nil {
		q = q.Set("even_class_id", *update.EvenClassID)
	}
	if update.OddClassID != nil {
		q = q.Set("odd_class_id", *update.OddClassID)
	}
	if update.IsStatic != nil {
		q = q.Set("is_static", *update.IsStatic)
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

func (r *EntryRepo) Delete(id uint64) error {
	query, args, err := r.psql.Delete("schedule_entries").
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
