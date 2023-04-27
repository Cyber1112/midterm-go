package updateCourseController

import "github.com/Cyber1112/midterm-go/models"

type Service interface {
	CreatePlaceService(input *InputUpdateCourse) (*models.Course, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdatePlaceService(input *InputUpdateCourse) (*models.Course, string) {
	course := models.Course{
		CourseID:  input.CourseId,
		StudentId: input.StudentId,
		Grade:     input.Grade,
	}

	resultCreatePlace, errCreatePlace := s.repository.UpdateCourseRepository(&course)

	return resultCreatePlace, errCreatePlace
}
