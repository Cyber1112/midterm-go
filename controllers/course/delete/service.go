package deleteCourseController

import "github.com/Cyber1112/midterm-go/models"

type Service interface {
	CreatePlaceService(input *InputDeleteCourse) (*models.Course, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeletePlaceService(input *InputDeleteCourse) string {
	course := models.Course{
		CourseID:  input.CourseId,
		StudentId: input.StudentId,
	}

	errCreatePlace := s.repository.DeleteCourseRepository(&course)

	return errCreatePlace
}
