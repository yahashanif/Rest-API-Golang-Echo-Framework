package controllers

import (
	"net/http"
	"rest-api-go-echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {

	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegawai(idInt, nama, alamat, telepon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)
}
