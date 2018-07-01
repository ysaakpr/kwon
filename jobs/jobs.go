package jobs

import (
	"time"

	"github.com/ysaakpr/kwon/base"
)

// Status defines the status of a Job
type Status int

func (s Status) String() string {
	names := [...]string{
		"INITIATED",
		"QUEUED",
		"SCHEDULED",
		"STARTED",
		"VALIDATING",
		"RUNNING",
		"DONE",
		"FAILED",
		"CANCELLED",
		"VALIDATION_SUCCESS",
		"VALIDATION_FAILED",
	}

	if s < Initiated || s > ValidationFailed {
		return "Unknown"
	}

	return names[s]
}

// Job Statuses
const (
	Initiated         Status = 0
	Queued            Status = 1
	Scheduled         Status = 2
	Started           Status = 3
	Validating        Status = 4
	Running           Status = 5
	Done              Status = 6
	Failed            Status = 7
	Cancelled         Status = 8
	ValidationSuccess Status = 9
	ValidationFailed  Status = 10
)

// Type defines the job type
type Type struct {
	ID          string
	Description string
	System      string
	IconClass   string
	Active      string
	CreateTime  time.Time
	UpdateTime  time.Time
}

//TableName for job_type
func (Type) TableName() string {
	return "job_type"
}

// Action defines the actions possible on a Job
type Action int

// Log defines the logs that user can be updated on a job
type Log struct {
	base.Model
	JobID uint
	Log   string
	Level string
}

//TableName for jobLog
func (Log) TableName() string {
	return "job_log"
}

// Job edfines the job descriptor
type Job struct {
	base.Model
	Status        string     `json:"status"`
	JobTypeID     string     `json:"-" gorm:"column:job_type"`
	JobType       *Type      `json:"job_type" gorm:"foreignkey:JobTypeID"`
	RunLength     int64      `json:"run_length" gorm:"column:run_length"`
	Progress      int64      `json:"progress" gorm:"column:progress"`
	ErrorCount    int64      `json:"error_count"`
	CreatedBy     int        `json:"created_by"`
	Payload       string     `json:"payload"`
	PayloadSchema string     `json:"payload_schema"`
	ScheduledAt   *time.Time `json:"scheduled_at"`
	StartTime     *time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	Action        *string    `json:"-" gorm:"-"`
	Logs          *[]Log     `json:"-" gorm:"-"`
}

//TableName overrides
func (Job) TableName() string {
	return "job"
}
