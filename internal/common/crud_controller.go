package common

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CrudController[P, Q, T any] interface {
	Read(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
}

func NewCrudController[P, Q, T any](
	reader ReaderService[Q, T],
	inserter WriterService[P],
	updater WriterService[T],
	remover WriterService[T],
	queryCreator QueryCreator[Q],
	deleteCriteriaCreator QueryCreator[T],
) CrudController[P, Q, T] {
	return &crudController[P, Q, T]{
		reader:                reader,
		inserter:              inserter,
		updater:               updater,
		remover:               remover,
		queryCreator:          queryCreator,
		deleteCriteriaCreator: deleteCriteriaCreator,
	}
}

type QueryCreator[Q any] interface {
	Create(*http.Request) *Q
}

type crudController[P, Q, T any] struct {
	reader                ReaderService[Q, T]
	inserter              WriterService[P]
	updater               WriterService[T]
	remover               WriterService[T]
	queryCreator          QueryCreator[Q]
	deleteCriteriaCreator QueryCreator[T]
}

func (controller *crudController[P, Q, T]) sendResult(writer http.ResponseWriter, result any) {
	err := json.NewEncoder(writer).Encode(result)

	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (controller *crudController[P, Q, T]) Read(writer http.ResponseWriter, request *http.Request) {

	list, err := controller.reader.Execute(request.Context(), controller.queryCreator.Create(request))

	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(NewErrorToJson(err))
		return
	}

	controller.sendResult(writer, list)
}

func (controller *crudController[P, Q, T]) Create(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	payload := new(P)
	err = json.Unmarshal(bodyBytes, payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	log.Println(payload)
	result, err := controller.inserter.Execute(request.Context(), payload)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(NewErrorToJson(err))
		return
	}

	writer.WriteHeader(http.StatusCreated)
	controller.sendResult(writer, result)
}

func (controller *crudController[P, Q, T]) Delete(writer http.ResponseWriter, request *http.Request) {

	_, err := controller.remover.Execute(request.Context(), controller.deleteCriteriaCreator.Create(request))
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(NewErrorToJson(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *crudController[P, Q, T]) Update(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	payload := new(T)
	err = json.Unmarshal(bodyBytes, payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	log.Println(payload)
	result, err := controller.updater.Execute(request.Context(), payload)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(NewErrorToJson(err))
		return
	}

	writer.WriteHeader(http.StatusOK)
	controller.sendResult(writer, result)
}
