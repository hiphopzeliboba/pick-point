package pickpoint

import (
	"encoding/json"
	"net/http"
	"pickpoint/internal/model"
	"pickpoint/internal/service"
	"strconv"
)

type PickPointHandler struct {
	pickPointService service.PickPointService
}

func NewPickPointHandler(pickPointService service.PickPointService) *PickPointHandler {
	return &PickPointHandler{
		pickPointService: pickPointService,
	}
}

func (h *PickPointHandler) Create(w http.ResponseWriter, r *http.Request) {
	var pickPoint model.PickPoint
	if err := json.NewDecoder(r.Body).Decode(pickPoint.City); err != nil {
		http.Error(w, "неверный формат запроса", http.StatusBadRequest)
		return
	}

	created, err := h.pickPointService.CreatePickPoint(r.Context(), &pickPoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *PickPointHandler) List(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	pickPoints, err := h.pickPointService.List(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pickPoints)
}
