package autoservice

import (
	"encoding/json"
	"log/slog"
	"main/internal/repositories/autorepository"
	"net/http"
)

type AutoService struct {
	autoRepository autorepository.AutoRepository
}

func New(autoRepository autorepository.AutoRepository) *AutoService {
	return &AutoService{
		autoRepository: autoRepository,
	}
}

func (as *AutoService) GetHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /autos", as.List)
	mux.HandleFunc("GET /autos/{id}", as.Get)
	mux.HandleFunc("POST /autos", as.Create)
	mux.HandleFunc("PUT /autos/{id}", as.Update)
	mux.HandleFunc("PATCH /autos/{id}", as.PartialUpdate)
	mux.HandleFunc("DELETE /autos/{id}", as.Delete)

	return mux
}

func (as *AutoService) List(w http.ResponseWriter, r *http.Request) {
	autos, err := as.autoRepository.List(r.Context())
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusOK,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(autos)
}

func (as *AutoService) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	a, err := as.autoRepository.Get(r.Context(), id)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusOK,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(a)
}

func (as *AutoService) Create(w http.ResponseWriter, r *http.Request) {
	var autoData autorepository.Auto
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = as.autoRepository.Create(r.Context(), &autoData)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusCreated,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(autoData)
}

func (as *AutoService) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var autoData autorepository.Auto
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = as.autoRepository.Update(r.Context(), &autoData, id)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusOK,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusOK)
}

func (as *AutoService) PartialUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var autoData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = as.autoRepository.PartialUpdate(r.Context(), autoData, id)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusNoContent,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusNoContent)
}

func (as *AutoService) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := as.autoRepository.Delete(r.Context(), id)
	if err != nil {
		defer slog.Error("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer slog.Info(
		"incoming request",
		"method", r.Method,
		"path", r.URL.Path,
		"status", http.StatusNoContent,
		"user_agent", r.UserAgent(),
	)
	w.WriteHeader(http.StatusNoContent)
}
