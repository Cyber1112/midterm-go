package createCourseController

import "github.com/Cyber1112/midterm-go/models"

type Service interface {
	CreateCourseService(input *InputCreateCourse) (*models.Course, string)
}

type service struct {
	repository Repository
}

func (s *service) CreateCourseService(input *InputCreateCourse) (*models.Course, string) {
	course := models.Course{
		CourseID:  input.CourseId,
		Grade:     input.Grade,
		StudentId: input.StudentId,
	}

	resultCreatePlace, errCreatePlace := s.repository.CreateCourseRepository(&course)

	return resultCreatePlace, errCreatePlace
}
