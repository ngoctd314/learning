package register

import (
	"encoding/json"
	"fmt"
	"go-learn/acme/internal/modules/data"
	"net/http"
)

type RegisterModel interface {
	Do(in *data.Person) (int, error)
}

// RegisterHandler is the HTTP handler for the "Register" endpoint
// In this simplified example we are assuming all possible errors
// are user errors and returining "bad request" HTTP 400.
type RegisterHandler struct {
	registerer RegisterModel
}

// NewRegisterHandler with constructor injection applied, our RegisterHandler is less coupled to the model
// layer and our external resources (database and upstream service).
func NewRegisterHandler(model RegisterModel) *RegisterHandler {
	return &RegisterHandler{
		registerer: model,
	}
}

func (h *RegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// extract payload from request
	requestPayload, err := h.extractPayload(request)
	if err != nil {
		// output error
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// register person
	id, err := h.register(requestPayload)
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// happy path
	response.Header().Add("Location", fmt.Sprintf("/person/%d/", id))
	response.WriteHeader(http.StatusCreated)
}

type registerRequest struct{}

func (h *RegisterHandler) extractPayload(request *http.Request) (*registerRequest, error) {
	requestPayload := &registerRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(requestPayload)
	if err != nil {
		return nil, err
	}

	return requestPayload, nil
}

func (h *RegisterHandler) register(requestPayload *registerRequest) (int, error) {
	// call the logic

	return h.registerer.Do(&data.Person{})
}
