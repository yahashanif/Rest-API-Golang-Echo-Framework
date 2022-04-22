package controllers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func HandleUpload(c echo.Context) error {
	if c.Request().Method != "POST" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Only Method POST"})
	}

	basePath, _ := os.Getwd()
	reader, err := c.Request().MultipartReader()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	part, err := reader.NextPart()
	if err == io.EOF {
		return err
	}

	fileLocation := filepath.Join(basePath, "files", part.FileName())
	dst, err := os.Create(fileLocation)
	if dst != nil {
		defer dst.Close()
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	if _, err := io.Copy(dst, part); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]string{"message": part.FileName()})
}
