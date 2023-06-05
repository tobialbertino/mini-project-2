package exception

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorTranslation(err error, c *gin.Context) {
	messages := make(map[string]string)
	english := en.New()
	uni := ut.New(english, english)

	// translate into bahasa
	trans, _ := uni.GetTranslator("en")

	if err != nil {
		object, _ := err.(validator.ValidationErrors)

		for _, key := range object {
			messages[key.Field()] = key.Translate(trans)
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"errors":  messages,
		"message": "Bad Request",
	})
}

func BindJSONError(err error, c *gin.Context) {
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			ValidationErrorTranslation(ve, c)
			return
		}
		NewClientError(http.StatusBadRequest, "Bad Request", c)
		return
	}
}
