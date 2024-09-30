package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RequestData struct {
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

type Point struct {
	X int
	Y int
}

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/get-path", GetPath)
	r.Run(":8080")
}

func CreateResponse(code int, message string, model interface{}) *Response {
	return &Response{
		Code:  code,
		Msg:   message,
		Model: model,
	}
}

func GetPath(c *gin.Context) {
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	start := Point{
		X: requestData.X1,
		Y: requestData.Y1,
	}
	end := Point{
		X: requestData.X2,
		Y: requestData.Y2,
	}

	paths := shortestPath(start, end)
	var response [][]int

	for _, path := range paths {
		single := []int{path.X, path.Y}
		response = append(response, single)
	}

	resp := CreateResponse(200, "Success", response)
	c.JSON(http.StatusOK, resp)
}

func shortestPath(start, end Point) []Point {
	path := []Point{start}
	current := start

	for current != end {
		if current.X < end.X {
			current.X++
		} else if current.X > end.X {
			current.X--
		} else if current.Y < end.Y {
			current.Y++
		} else if current.Y > end.Y {
			current.Y--
		}
		path = append(path, current)
	}

	return path
}
