package utils

import (
	"fmt"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

func HandleError(err error, c *gin.Context, gValid galidator.Validator) interface{} {
	if c.Request.ContentLength == 0 {
		return "provide body"
	}

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		if e.Field == "" {
			return "provide a json body"
		}
		msg := fmt.Sprintf("Invalid value for field '%s'. Expected a value of type '%s'", e.Field, e.Type)
		return msg
	}

	return gValid.DecryptErrors(err)
}