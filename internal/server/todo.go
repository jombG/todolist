package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"todolist/internal/metrics"
	"todolist/internal/model"
	"todolist/internal/server/generated"
	"todolist/pkg/pointer"
	"todolist/pkg/response"
)

func (s *Server) GetTasks(w http.ResponseWriter, r *http.Request) {
	res, err := s.useCases.GetTasks(r.Context())
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusInternalServerError,
			generated.Error{
				Code:    generated.INTERNALSERVERERROR,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}

	tasks := make([]generated.Task, 0, len(res))
	for _, elem := range res {
		tasks = append(tasks, generated.Task{
			Description: elem.Description,
			Id:          elem.ID.String(),
			Title:       elem.Title,
		})
	}

	response.JSON(
		w, http.StatusOK, tasks,
	)
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req generated.RequestCreateTask

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusBadRequest,
			generated.Error{
				Code:    generated.BADREQUEST,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}

	res, err := s.useCases.CreateTask(r.Context(), &model.Task{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusInternalServerError,
			generated.Error{
				Code:    generated.INTERNALSERVERERROR,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}

	metrics.MetricCreateTaskCounter.Inc()
	response.JSON(w, http.StatusCreated, generated.ResponseCreateTask{
		Description: res.Description,
		Id:          res.ID.String(),
		Title:       res.Title,
	})
}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request, id generated.ID) {
	taskId, err := uuid.Parse(id)
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusBadRequest,
			generated.Error{
				Code:    generated.BADREQUEST,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}
	err = s.useCases.DeleteTask(r.Context(), taskId)
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusInternalServerError,
			generated.Error{
				Code:    generated.INTERNALSERVERERROR,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}

	metrics.MetricDeleteTaskCounter.Inc()
	response.JSON(
		w, http.StatusNoContent, struct{}{},
	)
}

func (s *Server) FinishTask(w http.ResponseWriter, r *http.Request, id generated.ID) {
	taskId, err := uuid.Parse(id)
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusBadRequest,
			generated.Error{
				Code:    generated.BADREQUEST,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}
	err = s.useCases.FinishTask(r.Context(), taskId)
	if err != nil {
		metrics.MetricErrorCounter.Inc()
		response.JSON(
			w,
			http.StatusInternalServerError,
			generated.Error{
				Code:    generated.INTERNALSERVERERROR,
				Message: pointer.Pointer(err.Error()),
			},
		)
		return
	}

	metrics.MetricFinishTaskCounter.Inc()
	response.JSON(
		w, http.StatusNoContent, struct{}{},
	)
}
