package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	num1Str := c.Query("num1")
	num2Str := c.Query("num2")

	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)

	response := map[string]int{
		"result": num1 + num2,
	}

	c.JSON(http.StatusOK, response)
}

func Subtract(c *gin.Context) {
	num1Str := c.Query("num1")
	num2Str := c.Query("num2")

	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)

	response := map[string]int{
		"result": num1 - num2,
	}

	c.JSON(http.StatusOK, response)
}

func Divide(c *gin.Context) {
	num1str := c.Query("num1")
	num2str := c.Query("num2")

	num1, _ := strconv.Atoi(num1str)
	num2, _ := strconv.Atoi(num2str)

	response := map[string]int{
		"result": num1 / num2,
	}

	c.JSON(http.StatusOK, response)
}

func Multiply(c *gin.Context) {
	num1Str := c.Query("num1")
	num2str := c.Query("num2")

	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2str)

	response := map[string]int{
		"result": num1 * num2,
	}

	c.JSON(http.StatusOK, response)
}
