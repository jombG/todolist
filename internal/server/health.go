package server

import (
	"net/http"
	"todolist/internal/server/generated"
	"todolist/pkg/pointer"
	"todolist/pkg/response"
)

func (s *Server) GetHealth(w http.ResponseWriter, r *http.Request) {
	res, err := s.useCases.GetHealth(r.Context())
	if err != nil {
		response.JSON(
			w,
			http.StatusInternalServerError,
			generated.Error{
				Code:    generated.INTERNALSERVERERROR,
				Message: pointer.Pointer(err.Error()),
			},
		)
	}
	response.JSON(
		w, http.StatusOK, generated.RespSuccess{Ok: res.Ok},
	)
}
