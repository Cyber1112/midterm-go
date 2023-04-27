package resultCourseController

import (
	"github.com/Cyber1112/midterm-go/models"
	"github.com/go-redis/redis"
)

type Repository interface {
	ResultCourseRepository(input *models.Course) (*models.Course, string)
}

type repository struct {
	db *redis.Client
}

func NewRepositoryCreate() *repository {
	return &repository{}
}

func (r *repository) ResultCourseRepository(input *models.Course) (*models.Course, string) {

}
