package request

type RequestActivity struct {
	ActivityType      *string `json:"activityType"`
	DoneAt            *string `json:"doneAt"`
	DurationInMinutes *int    `json:"durationInMinutes"`
	UserId            *string `json:"userId"`
}
