package updateCourseController

type InputUpdateCourse struct {
	CourseId  uint `json:"course_id" validate:"required"`
	StudentId uint `json:"student_id" validate:"required"`
	Grade     uint `json:"grade" validate:"required"`
}
