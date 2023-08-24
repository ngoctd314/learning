package main





import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name string
}

func main() {
	sqlx.Open("", "")
	gin.Default()
}
