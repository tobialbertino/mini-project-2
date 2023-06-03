package exception

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Custom wrap error
type ClientError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (w *ClientError) Error() string {
	return fmt.Sprintf(`%v: %v `, w.Status, w.Message)
}

func NewClientError(code int, message string, c *gin.Context) {
	result := ClientError{
		Status:  http.StatusText(code),
		Message: message,
	}

	c.JSON(code, result)
}
