package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"regexp"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 32),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

func RegisterValidator() gin.HandlerFunc {
	return func(context *gin.Context) {

		var registerRequest RegisterRequest
		_ = context.ShouldBindBodyWith(&registerRequest, binding.JSON)

		if err := registerRequest.Validate(); err != nil {
			models.SendErrorResponse(context, http.StatusBadRequest, err.Error())
			return
		}

		context.Next()
	}
}

func LoginValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}

func RefreshValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
