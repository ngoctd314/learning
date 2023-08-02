package csrf

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type Tweet struct {
	Title   string
	Content string
}

// Run ...
func Run(db *sqlx.DB) {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		db.NamedExec("INSERT INTO tweets (:title, :content)", Tweet{
			Title:   fmt.Sprintf("title_%d", time.Now().UnixMilli()),
			Content: fmt.Sprintf("content_%d", time.Now().UnixMilli()),
		})
		w.Write([]byte("create success"))
	})

	http.ListenAndServe(":8080", nil)
}
