package Entity

import "time"

type Activity struct {
	ActivityId        *string
	ActivityType      *string
	DoneAt            *time.Time
	DurationInMinutes *int64
	CaloriesBurned    *float64
	CreatedAt         *time.Time
}

type ActivityType string

const (
	Walking    ActivityType = "Walking"
	Yoga       ActivityType = "Yoga"
	Stretching ActivityType = "Stretching"
	Cycling    ActivityType = "Cycling"
	Swimming   ActivityType = "Swimming"
	Dancing    ActivityType = "Dancing"
	Hiking     ActivityType = "Hiking"
	Running    ActivityType = "Running"
	HIIT       ActivityType = "HIIT"
	JumpRope   ActivityType = "JumpRope"
)
