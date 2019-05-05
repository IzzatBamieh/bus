package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	bus    *Bus
	logger *Logger
	forge  *LambdaNodeForgeService
}

type HTTPCreateServiceCommand struct {
	Encoder
	Name string
}

func newHTTPCreateServiceCommand(reader io.Reader) (*HTTPCreateServiceCommand, error) {
	command := &HTTPCreateServiceCommand{}
	err := command.decode(reader, command)

	return command, err
}

func (command *HTTPCreateServiceCommand) GetName() string {
	return command.Name
}

type httpProblem struct {
	Encoder
	Message string
	Error   string
}

func newHTTPProblem(response http.ResponseWriter, err error, message string) {
	problem := httpProblem{
		Message: message,
		Error:   err.Error(),
	}

	response.WriteHeader(400)
	err = problem.encode(response, problem)
	if err != nil {
		return
	}
}

type listServicesResponse struct {
	Encoder
	Services []string
}

func NewHandler(bus *Bus, logger *Logger, forge *LambdaNodeForgeService) *Handler {
	return &Handler{
		bus:    bus,
		logger: logger,
		forge:  forge,
	}
}

func setupResponse(response *http.ResponseWriter, req *http.Request) {
	(*response).Header().Set("Access-Control-Allow-Origin", "*")
	(*response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*response).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
}

func (handler *Handler) PostServices(
	response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	setupResponse(&response, request)
	if request.Method == "OPTIONS" {
		response.WriteHeader(200)
		return
	}

	command, err := newHTTPCreateServiceCommand(request.Body)
	if err != nil {
		newHTTPProblem(response, err, "could not understand the request")
		return
	}
	err = handler.forge.CreateService(command)
	if err != nil {
		newHTTPProblem(response, err, "problem creating the service")
		return
	}
	response.WriteHeader(201)
}

func (handler *Handler) GetServices(
	response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	setupResponse(&response, request)
	services := handler.forge.ListServices()
	view := &listServicesResponse{
		Services: services,
	}
	response.WriteHeader(200)
	err := view.encode(response, view)
	if err != nil {
		handler.logger.Error(err)
	}
}

func (handler *Handler) DeleteService(
	writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	setupResponse(&writer, request)
	err := handler.forge.DeleteService(params.ByName("name"))
	if err != nil {
		newHTTPProblem(writer, err, "problem deleting the service")
		return
	}
	writer.WriteHeader(204)
}
