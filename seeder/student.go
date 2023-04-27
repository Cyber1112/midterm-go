package seeder

import (
	"encoding/json"
	"github.com/Cyber1112/midterm-go/models"
	"github.com/go-redis/redis"
)

func Seed(client *redis.Client) {
	encoded, _ := json.Marshal(models.Student{ID: 1, Name: "Khakim"})

	client.LPush("students", encoded)

	var people [3]models.Course

	// Assign values to the elements of the array
	people[0] = models.Course{CourseID: 1, Name: "Math", Grade: 5, StudentId: 1}
	people[1] = models.Course{CourseID: 2, Name: "Physics", Grade: 4, StudentId: 1}
	people[2] = models.Course{CourseID: 3, Name: "Computing", Grade: 3, StudentId: 1}

	encoded, _ = json.Marshal(people)

	client.LPush("courses", encoded)
}
