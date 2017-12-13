package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Render content
func Render(c *gin.Context, data gin.H) {
	code := data["code"]
	payload := data["payload"]
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(code.(int), payload)
	case "application/xml":
		// Respond with XML
		c.XML(code.(int), payload)
	default:
		// Respond with JSON
		c.JSON(code.(int), payload)
	}
}

func string2Int(value string) int {
	parse, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}
	return int(parse)
}

// LimitAndOffset is a function to parse limit and offset
func LimitAndOffset(limit string, offset string) (int, int) {
	reLimit := string2Int(limit)
	if reLimit == 0 {
		reLimit = -1
	}
	reOffset := string2Int(offset)
	return reLimit, reOffset
}

// ParseParam2Int is a function to parse value to int
func ParseParam2Int(value string) uint {
	parse, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0
	}
	return uint(parse)
}
