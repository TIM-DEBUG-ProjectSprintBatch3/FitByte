package request

type ActivityUriParam struct {
	Limit             string  `uri:"limit" validate:""`
	Offset            string  `uri:"offset" validate:""`
	ActivityType      string  `uri:"activityType" validate:""`
	DoneAtFrom        float64 `uri:"doneAtFrom" validate:""`
	DoneAtTo          float64 `uri:"doneAtTo" validate:""`
	CaloriesBurnedMin float64 `uri:"caloriesBurnedMin" validate:""`
	CaloriesBurnedMax float64 `uri:"caloriesBurnedMax" validate:""`
}
