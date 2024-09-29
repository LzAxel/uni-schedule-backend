package schedule

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"uni-schedule-backend/internal/domain"
)

type ScheduleSlotRepo struct {
	db   *sqlx.DB
	psql squirrel.StatementBuilderType
}

func NewScheduleSlotRepo(db *sqlx.DB) *ScheduleSlotRepo {
	return &ScheduleSlotRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *ScheduleSlotRepo) Create(slot domain.ScheduleSlot) (domain.ID, error) {
	query, args, err := r.psql.Insert("schedule_slots").
		Columns("schedule_id", "weekday", "number", "is_alternating", "even_week_lesson_id", "odd_week_lesson_id").
		Values(slot.ScheduleID, slot.Weekday, slot.Number, slot.IsAlternating, slot.EvenWeekLessonID, slot.OddWeekLessonID).
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

func (r *ScheduleSlotRepo) GetByID(id domain.ID) (domain.ScheduleSlot, error) {
	query, args, err := r.psql.Select("*").From("schedule_slots").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return domain.ScheduleSlot{}, err
	}

	var slot domain.ScheduleSlot
	err = r.db.QueryRow(query, args...).Scan(&slot.ID, &slot.ScheduleID, &slot.Weekday, &slot.Number, &slot.IsAlternating, &slot.EvenWeekLessonID, &slot.OddWeekLessonID)
	if err != nil {
		return domain.ScheduleSlot{}, err
	}
	return slot, nil
}

func (r *ScheduleSlotRepo) Update(id domain.ID, update domain.ScheduleSlotUpdate) error {
	q := r.psql.Update("schedule_slots").Where(squirrel.Eq{"id": id})

	if update.ScheduleID != nil {
		q = q.Set("schedule_id", *update.ScheduleID)
	}
	if update.Weekday != nil {
		q = q.Set("weekday", *update.Weekday)
	}
	if update.Number != nil {
		q = q.Set("number", *update.Number)
	}
	if update.IsAlternating != nil {
		q = q.Set("is_alternating", *update.IsAlternating)
	}
	if update.EvenWeekLessonID != nil {
		q = q.Set("even_week_lesson_id", *update.EvenWeekLessonID)
	}
	if update.OddWeekLessonID != nil {
		q = q.Set("odd_week_lesson_id", *update.OddWeekLessonID)
	}

	query, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *ScheduleSlotRepo) Delete(id domain.ID) error {
	query, args, err := r.psql.Delete("schedule_slots").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}
