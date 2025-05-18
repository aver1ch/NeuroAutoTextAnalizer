package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/models"
	"github.com/kerilOvs/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req struct {
		Id       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Surname  string    `json:"surname"` // `json:"about_myself,omitempty"`
		Email    string    `json:"email,omitempty"`
		Password string    `json:"password,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	user, err := h.service.CreateUser(req.Id, req.Name, req.Surname, req.Email, req.Password)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusBadRequest)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	if err := h.service.DeleteUser(id); err != nil {
		return errorResponseWithCode(c, err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) GetUserById(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	var req models.UserProfileUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := h.service.UpdateUserProfile(id, req); err != nil {
		return errorResponseWithCode(c, err, http.StatusBadRequest)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) GetUserDocs(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	photos, err := h.service.GetUserDocs(id)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, photos)
}

func (h *UserHandler) RemoveUserDoc(c echo.Context) error {

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	photoID, err := uuid.Parse(c.Param("photoId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid photo ID"))
	}

	if err := h.service.RemoveUserDoc(userID, photoID); err != nil {
		return errorResponseWithCode(c, err, http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

func errorResponse(msg string) map[string]interface{} {
	return map[string]interface{}{"error": msg}
}

func errorResponseWithCode(c echo.Context, err error, code int) error {
	return c.JSON(code, errorResponse(err.Error()))
}

// если потребуется проверка jwt токена

/*func getJWTUserID(r *http.Request) (uuid.UUID, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return uuid.UUID{}, fmt.Errorf("no authorization header provided")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to parse jwt: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("failed to parse jwt claims")
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to get subject from jwt: %w", err)
	}

	id, err := uuid.Parse(sub)

	return id, err
}*/

/*requestedID, err := uuid.Parse(c.Param("id"))
if err != nil {
	return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
}

tokenUserID, err := getJWTUserID(c.Request())
if err != nil {
	return c.JSON(http.StatusUnauthorized, errorResponse("Invalid or missing JWT token"))
}

if requestedID != tokenUserID {
	return c.JSON(http.StatusForbidden, errorResponse("You can only access your own data"))
}*/
