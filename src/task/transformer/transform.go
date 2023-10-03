package transformer

import (
	"database/sql"

	"github.com/fwidjaya20/symphonic-example/src/task/application/public"
	"github.com/fwidjaya20/symphonic-example/src/task/domain/entity"
	"github.com/golang-module/carbon"
)

func TransformRequestToEntity(request public.CreateTaskRequest) entity.Task {
	task := entity.Task{
		Title:       request.Title,
		Description: sql.NullString{},
		IsCompleted: false,
		IsPriority:  request.IsPriority,
		DueDate:     request.DueDate,
	}

	if request.Description != nil {
		task.Description = sql.NullString{
			String: *request.Description,
			Valid:  true,
		}
	}

	return task
}

func TransformTaskEntityAsResponse(task entity.Task) public.TaskResponse {
	response := public.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
		IsPriority:  task.IsPriority,
		DueDate:     task.DueDate,
		FmtDueDate:  carbon.Parse(task.DueDate.String()).Format("M, d Y H:i:s O"),
	}

	if task.Description.Valid {
		response.Description = &task.Description.String
	}

	return response
}

func TransformTaskEntitiesAsResponses(tasks []entity.Task) []public.TaskResponse {
	var responses = make([]public.TaskResponse, len(tasks))

	for i, it := range tasks {
		responses[i] = TransformTaskEntityAsResponse(it)
	}

	return responses
}
