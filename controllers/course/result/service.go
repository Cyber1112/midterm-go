package resultCourseController

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

func (s *service) ResultPlaceService(input *InputDeleteCourse) (*models.Course, string) {
	course := models.Course{
		CourseID: input.CourseId,
	}

	resultCreatePlace, errCreatePlace := s.repository.ResultCourseRepository(&course)

	return resultCreatePlace, errCreatePlace
}
