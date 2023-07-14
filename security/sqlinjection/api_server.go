package sqlinjection

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// APIServer ...
func APIServer(db *sqlx.DB) {

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var request Request
		json.Unmarshal(data, &request)

		if _, err := db.Exec("UPDATE users SET age = ? WHERE name = ?", request.Age, request.Name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("UPDATE success"))
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		var listUser []User
		if err := db.Select(&listUser, "SELECT * FROM users"); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		data, _ := json.Marshal(listUser)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	})

	http.ListenAndServe(":8080", nil)
}

// User model
type User struct {
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

// Request ...
type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
