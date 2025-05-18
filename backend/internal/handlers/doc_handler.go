package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/kerilOvs/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type DocHandler struct {
	userService *service.UserService
	docService  *service.DocService
}

func NewDocHandler(userService *service.UserService, docService *service.DocService) *DocHandler {
	return &DocHandler{
		userService: userService,
		docService:  docService,
	}
}

func (h *DocHandler) UploadDoc(c echo.Context) error {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("invalid user id"))
	}

	// Получаем файл из формы
	file, err := c.FormFile("document")
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("document is required"))
	}

	// Проверяем расширение файла (опционально)
	allowedExtensions := map[string]bool{
		".doc":  true,
		".docx": true,
	}

	ext := filepath.Ext(file.Filename)
	if !allowedExtensions[ext] {
		return c.JSON(http.StatusBadRequest, errorResponse("invalid file type"))
	}

	// Открываем файл
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse("failed to read file"))
	}
	defer src.Close()

	// Загружаем фото в MinIO
	objectName, err := h.docService.UploadDoc(c.Request().Context(), src, file.Size)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusInternalServerError)
	}

	// Получаем URL для доступа к фото
	docURL, err := h.docService.GetDocURL(c.Request().Context(), objectName)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusInternalServerError)
	}

	// Сохраняем информацию о фото в БД
	doc, err := h.userService.AddUserDoc(userID, docURL)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, doc)
}

func (h *DocHandler) GetDoc(c echo.Context) error {
	objectName := c.Param("id")
	url, err := h.docService.GetDocURL(c.Request().Context(), objectName)
	if err != nil {
		return errorResponseWithCode(c, err, http.StatusNotFound)
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}
