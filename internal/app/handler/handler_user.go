package handler

import (
    "encoding/json"
    "net/http"
)

// UserHandler handle requests related with users
type UserHandler struct {
    AppHandler *AppHandler
}

// Create new instance of UserHandler
func NewUserHandler(appHandler *AppHandler) *UserHandler {
    return &UserHandler{
        AppHandler: appHandler,
    }
}

// Return user by ID
func (ah *ArticleHandler) GetUser(w http.ResponseWriter, r *http.Request) {

    userID := 1
    user, err := ah.AppHandler.userService.GetUserByID(userID)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Send user like a JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
