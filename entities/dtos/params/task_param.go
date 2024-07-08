package params

import (
	"net/url"
	"task_mission/enums"
	"time"
)

type TaskParam struct {
	Title       string
	Description string
	Status      enums.TaskStatus
	ExpiredDate time.Time
}

func NewTaskParam(queryParam url.Values) *TaskParam {
	expiredDate, _ := time.Parse("2006-01-02", queryParam.Get("expiredDate"))
	return &TaskParam{
		Title:       queryParam.Get("title"),
		Description: queryParam.Get("description"),
		Status:      enums.TaskStatus(queryParam.Get("status")),
		ExpiredDate: expiredDate,
	}
}
