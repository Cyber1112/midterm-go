package deleteCourseController

import (
	"github.com/Cyber1112/midterm-go/models"
	"github.com/go-redis/redis"
)

type Repository interface {
	DeleteCourseRepository(input *models.Course) string
}

type repository struct {
	db *redis.Client
}

func NewRepositoryCreate() *repository {
	return &repository{}
}

func (r *repository) DeleteCourseRepository(input *models.Course) string {

}
