package cmdinjection

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

func APIServer() {
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		sh, lookErr := exec.LookPath("sh")
		if lookErr != nil {
			w.Write([]byte("execute command fail" + lookErr.Error()))
			return
		}
		args := []string{"sh", "-c", name}
		err := syscall.Exec(sh, args, os.Environ())
		if err != nil {
			w.Write([]byte("execute command fail" + err.Error()))
			return
		}

		w.Write([]byte("execute command success"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
