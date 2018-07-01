package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

//Service defines the service logics of Job
type Service struct {
	db *gorm.DB
}

//NewJobService Create new JobService instance
func NewJobService(db *gorm.DB) *Service {
	service := &Service{}
	service.db = db
	return service
}

//Search will search all the jobs in the store
func (s *Service) Search(ctx *iris.Context) []Job {
	var jobs []Job
	s.db.Limit(10).Preload("JobType").Find(&jobs)
	return jobs
}
