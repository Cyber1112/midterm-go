package createCourseController

import (
	"github.com/Cyber1112/midterm-go/models"
	"github.com/go-redis/redis"
)

type Repository interface {
	CreateCourseRepository(input *models.Course) (*models.Course, string)
}

type repository struct {
	db *redis.Client
}

func (r *repository) CreateCourseRepository(input *models.Course) (*models.Course, string) {
	values, _ := r.db.Get("courses").Result()

	print(values)

	//key := -1;
	//course := models.Course{}

	//for key, value := range values {
	//	if input.CourseID == value["course_id"] && value["student_id"] == input.StudentId {
	//		prin
	//		break
	//	}
	//}
}
