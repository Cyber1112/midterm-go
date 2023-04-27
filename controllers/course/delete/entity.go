package deleteCourseController

type InputDeleteCourse struct {
	CourseId  uint `json:"course_id" validate:"required"`
	StudentId uint `json:"student_id" validate:"required"`
}
