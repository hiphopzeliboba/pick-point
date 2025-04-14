package intake

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pickpoint/internal/model"
	"pickpoint/internal/service"
	"strconv"
)

type IntakeHandler struct {
	intakeService service.IntakeService
}

func NewIntakeHandler(intakeService service.IntakeService) *IntakeHandler {
	return &IntakeHandler{
		intakeService: intakeService,
	}
}

type createIntakeRequest struct {
	PickPointId int `json:"pvzId"`
}

func (h *IntakeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createIntakeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "неверный формат запроса", http.StatusBadRequest)
		return
	}

	intake := &model.Intake{
		PickPointId: req.PickPointId,
	}

	created, err := h.intakeService.CreateIntake(r.Context(), intake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *IntakeHandler) CloseLastReception(w http.ResponseWriter, r *http.Request) {
	pvzId := chi.URLParam(r, "pvzId")
	id, err := strconv.Atoi(pvzId)
	if err != nil {
		http.Error(w, "неверный формат id", http.StatusBadRequest)
		return
	}

	err = h.intakeService.CloseIntake(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
