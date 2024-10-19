package schedule

import (
	"errors"
	"time"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type ScheduleService struct {
	scheduleRepo     repository.ScheduleRepository
	scheduleSlotRepo repository.ScheduleSlotRepository
	lessonRepo       repository.LessonRepository
	teacherRepo      repository.TeacherRepository
}

func NewScheduleService(
	scheduleRepo repository.ScheduleRepository,
	scheduleSlotRepo repository.ScheduleSlotRepository,
	lessonRepo repository.LessonRepository,
	teacherRepo repository.TeacherRepository,
) *ScheduleService {
	return &ScheduleService{
		scheduleRepo:     scheduleRepo,
		scheduleSlotRepo: scheduleSlotRepo,
		lessonRepo:       lessonRepo,
		teacherRepo:      teacherRepo,
	}
}

func (s *ScheduleService) CreateSchedule(schedule domain.ScheduleCreate) (uint64, error) {
	if schedule.Slug == "" {
		return 0, apperror.ErrEmptyScheduleSlug
	}

	if schedule.Name == "" {
		return 0, apperror.ErrEmptyScheduleName
	}

	return s.scheduleRepo.Create(schedule)
}
func (s *ScheduleService) GetScheduleBySlug(slug string) (domain.ScheduleView, error) {
	schedule, err := s.scheduleRepo.GetBySlug(slug)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return domain.ScheduleView{}, apperror.ErrScheduleNotFound
		}

		return domain.ScheduleView{}, err
	}

	slots, err := s.scheduleSlotRepo.GetAllSlotsByScheduleID(schedule.ID)
	if err != nil {
		return domain.ScheduleView{}, err
	}

	groupedSlotsMap := make(map[time.Weekday][]domain.ScheduleSlotView)

	for _, slot := range slots {
		var slotView = domain.ScheduleSlotView{
			ID:             slot.ID,
			Number:         slot.Number,
			IsAlternating:  slot.IsAlternating,
			EvenWeekLesson: nil,
			OddWeekLesson:  nil,
		}

		if slot.EvenWeekLessonID != nil {
			lessonView, err := s.lessonRepo.GetWithRelationsByID(*slot.EvenWeekLessonID)
			if err != nil {
				return domain.ScheduleView{}, err
			}

			slotView.EvenWeekLesson = &lessonView
		}
		if slot.OddWeekLessonID != nil {
			lessonView, err := s.lessonRepo.GetWithRelationsByID(*slot.OddWeekLessonID)
			if err != nil {
				return domain.ScheduleView{}, err
			}

			slotView.OddWeekLesson = &lessonView
		}

		if _, ok := groupedSlotsMap[slot.Weekday]; !ok {
			groupedSlotsMap[slot.Weekday] = make([]domain.ScheduleSlotView, 0)
		}

		groupedSlotsMap[slot.Weekday] = append(groupedSlotsMap[slot.Weekday], slotView)
	}

	view := domain.ScheduleView{
		ID:       schedule.ID,
		Name:     schedule.Name,
		Slug:     schedule.Slug,
		Weekdays: make([]domain.ScheduleGroupedSlotView, 0),
	}

	for k, v := range groupedSlotsMap {
		view.Weekdays = append(view.Weekdays, domain.ScheduleGroupedSlotView{
			Day:   k,
			Slots: v,
		})
	}

	return view, nil
}
func (s *ScheduleService) UpdateSchedule(id uint64, update domain.ScheduleUpdate) error {
	if update.Slug != nil && *update.Slug == "" {
		return apperror.ErrEmptyScheduleSlug
	}

	if update.Name != nil && *update.Name == "" {
		return apperror.ErrEmptyScheduleName
	}
	return s.scheduleRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSchedule(id uint64) error {
	return s.scheduleRepo.Delete(id)
}
