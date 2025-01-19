package Entity

import (
	"strings"
	"time"
)

type Activity struct {
	ActivityId        string
	UserId            string
	DoneAt            time.Time
	DurationInMinutes int64
	CaloriesBurned    float64
	ActivityType      ActivityType
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CaloriesFactor struct {
	ActivityType      *string
	DurationInMinutes *int
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

func CountCalories(minutes int64, activityType ActivityType) float64 {
	caloriesPerMinute := map[ActivityType]float64{
		Walking:    4.0,
		Yoga:       4.0,
		Stretching: 4.0,
		Cycling:    8.0,
		Swimming:   8.0,
		Dancing:    8.0,
		Hiking:     10.0,
		Running:    10.0,
		HIIT:       10.0,
		JumpRope:   10.0,
	}

	if calories, exists := caloriesPerMinute[activityType]; exists {
		return calories * float64(minutes)
	}
	return 0.0
}

func IsValidActivityType(activityType string) bool {
	activityType = strings.ToLower(activityType)
	switch ActivityType(activityType) {
	case
		"walking",
		"yoga",
		"stretching",
		"cycling",
		"swimming",
		"dancing",
		"hiking",
		"running",
		"hiit",
		"jumprope":
		return true
	default:
		return false
	}
}
