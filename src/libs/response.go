package libs

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"-"`
}

func (res *Resp) Send(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(res.Code)
	return json.NewEncoder(c.Response()).Encode(res)
}

func Response(data interface{}, status string, message string, code int) *Resp {
	return &Resp{
		Status:  status,
		Message: message,
		Data:    data,
		Code:    code,
	}
}
