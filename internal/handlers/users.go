package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrxacker/go-myapp/internal/dto"
	"github.com/mrxacker/go-myapp/internal/models"
	"github.com/mrxacker/go-myapp/internal/service"
	"github.com/mrxacker/go-myapp/pkg/logger"
)

type UserHandler struct {
	userService *service.UserService
	logger      logger.Logger
}

func NewUserHandler(userService *service.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	if userID == "" {
		h.logger.Warn("User ID is empty")
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	h.logger.Debug("Fetching user", logger.String("user_id", userID))

	user, err := h.userService.GetUser(models.UserID(userID))
	if err != nil {
		h.logger.Error("Failed to fetch user", logger.String("user_id", userID), logger.Error(err))
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	h.logger.Info("User fetched successfully", logger.String("user_id", userID))
	respondWithJSON(w, http.StatusOK, user)
}

// Create handles POST /api/users
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Warn("Invalid request body", logger.Error(err))
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Username == "" || req.Email == "" {
		h.logger.Warn("Missing required fields")
		respondWithError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	h.logger.Debug("Creating user", logger.String("name", req.Username), logger.String("email", req.Email))

	user, err := h.userService.CreateUser(req.Username, req.Email)
	if err != nil {
		h.logger.Error("Failed to create user", logger.Error(err))
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	h.logger.Infof("User created successfully, user_id=%v", user.ID)
	respondWithJSON(w, http.StatusCreated, user)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
