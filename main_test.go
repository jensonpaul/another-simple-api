package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test para la función sumar
func TestSumar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Caso de prueba: Suma válida
	a, b := 5, 3
	esperado := a + b

	req, _ := http.NewRequest("GET", fmt.Sprintf("/sumar?a=%d&b=%d", a, b), nil)
	resp := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/sumar", sumar)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(body), fmt.Sprintf(`"resultado":%d`, esperado))

	// Caso de prueba: Parámetros inválidos
	req, _ = http.NewRequest("GET", "/sumar?a=invalid&b=3", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Test para la función restar
func TestRestar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Caso de prueba: Resta válida
	a, b := 7, 2
	esperado := a - b

	req, _ := http.NewRequest("GET", fmt.Sprintf("/restar?a=%d&b=%d", a, b), nil)
	resp := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/restar", restar)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(body), fmt.Sprintf(`"resultado":%d`, esperado))

	// Caso de prueba: Parámetros inválidos
	req, _ = http.NewRequest("GET", "/restar?a=invalid&b=2", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Test para la función multiplicar
func TestMultiplicar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Caso de prueba: Resta válida
	a, b := 5, 9
	esperado := a * b

	req, _ := http.NewRequest("GET", fmt.Sprintf("/multiplicar?a=%d&b=%d", a, b), nil)
	resp := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/multiplicar", multiplicar)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(body), fmt.Sprintf(`"resultado":%d`, esperado))

	// Caso de prueba: Parámetros inválidos
	req, _ = http.NewRequest("GET", "/multiplicar?a=invalid&b=9", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Test para la función dividir
func TestDividir(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Caso de prueba: Resta válida
	a, b := 25, 5
	esperado := a / b

	req, _ := http.NewRequest("GET", fmt.Sprintf("/dividir?a=%d&b=%d", a, b), nil)
	resp := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/dividir", dividir)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(body), fmt.Sprintf(`"resultado":%d`, esperado))

	// Caso de prueba: Parámetros inválidos
	req, _ = http.NewRequest("GET", "/dividir?a=invalid&b=5", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// Prueba de integración del servidor
func TestServidor(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest("GET", "/sumar?a=2&b=3", nil)
	resp := httptest.NewRecorder()

	router := crearRouter()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
