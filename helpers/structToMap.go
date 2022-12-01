package helpers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

func StructToMap(data interface{}) (m *echo.Map) {
	mapData, _ := json.Marshal(data)
	json.Unmarshal(mapData, &m)

	return m
}
