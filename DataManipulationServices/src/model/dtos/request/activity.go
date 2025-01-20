package request

type RequestActivity struct {
	ActivityType      *string `json:"activityType"`
	DoneAt            *string `json:"doneAt"`
	DurationInMinutes *int    `json:"durationInMinutes"`
	UserId            *string `json:"userId"`
}

// To differentiate if property is null as present (e.g. { "name": null }) or null as not present (not sent at all)
type RequestActivityCustom struct {
	ActivityType      CustomString `json:"activityType"`
	DoneAt            CustomString `json:"doneAt"`
	DurationInMinutes CustomInt    `json:"durationInMinutes"`
	UserId            CustomString `json:"userId"`
}
