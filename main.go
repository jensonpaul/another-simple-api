package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creamos un router y lo configuramos
	router := crearRouter()

	// Iniciamos el servidor en el puerto 8080
	router.Run(":8080")
}

// Función para crear un router y configurar las rutas
func crearRouter() *gin.Engine {
	// Creamos una nueva instancia de Gin
	router := gin.Default()

	// Definimos los endpoints para cada operación
	router.GET("/sumar", sumar)
	router.GET("/restar", restar)
	router.GET("/multiplicar", multiplicar)
	router.GET("/dividir", dividir)

	return router
}

// Función para sumar dos números
func sumar(c *gin.Context) {
	// Obtenemos los parámetros 'a' y 'b' de la URL
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	// Realizamos la suma
	resultado := a + b

	// Retornamos el resultado en formato JSON
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// Función para restar dos números
func restar(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a - b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// Función para multiplicar dos números
func multiplicar(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a * b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}

// Función para dividir dos números
func dividir(c *gin.Context) {
	a, errA := strconv.ParseFloat(c.Query("a"), 64)
	b, errB := strconv.ParseFloat(c.Query("b"), 64)

	if errA != nil || errB != nil || b == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	resultado := a / b
	c.JSON(http.StatusOK, gin.H{"resultado": resultado})
}
