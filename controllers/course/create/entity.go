package createCourseController

type InputCreateCourse struct {
	StudentId uint `json:"student_id" validate:"required"`
	CourseId  uint `json:"course_id" validate:"required"`
	Grade     uint `json:"grade" validate:"required"`
}
