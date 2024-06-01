package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" validate:"gte=0,lte=130"`
	CheckOUt time.Time `form:"check_out" binding:"required"`
}

type Person struct {
	Age  int    `form:"age" binding:"required,gte=18,lte=130"`
	Role string `form:"role" binding:"required,oneof=admin user"`
	Name string `form:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var nw Person
		if err := c.ShouldBindQuery(&nw); err != nil {
			var vErr validator.ValidationErrors
			var lists []ValidatorError
			if errors.As(err, &vErr) {
				for _, e := range vErr {
					lists = append(lists, NewValidatorError(e))
				}
			}
			c.JSON(400, gin.H{
				"errors": lists,
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": nw,
		})
	})

	r.Run()
}

func getErrorMsg(fe validator.FieldError) []byte {
	m := map[string]any{
		"tag":              fe.Tag(),
		"actual_tag":       fe.ActualTag(),
		"namespace":        fe.Namespace(),
		"struct_namespace": fe.StructNamespace(),
		"field":            fe.Field(),
		"struct_field":     fe.StructField(),
		"value":            fe.Value(),
		"param":            fe.Param(),
		"kind":             fe.Kind(),
		"type":             fe.Type(),
		// "translate":        fe.Translate(),
		"error": fe.Error(),
	}

	data, _ := json.Marshal(m)
	return data
}
