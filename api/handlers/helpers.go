package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, code int, message error) {
	JSON(w, code, map[string]string{"error": message.Error()})
}

func InternalServerError(w http.ResponseWriter, err error) {
	Error(w, http.StatusInternalServerError, err)
}

func BadRequest(w http.ResponseWriter, err error) {
	Error(w, http.StatusBadRequest, err)
}

func NotFound(w http.ResponseWriter, err error) {
	Error(w, http.StatusNotFound, err)
}

func Unimplemented(w http.ResponseWriter) {
	Error(w, http.StatusNotImplemented, errors.New("not implemented"))
}

func GetBody[T any](r *http.Request) (T, error) {
	var body T

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}

	return body, nil
}

func GetParam(r *http.Request, param string) (string, error) {
	value := r.PathValue(param)
	if value == "" {
		return value, errors.New("missing parameter")
	}

	return value, nil
}

func GetID(r *http.Request) (string, error) {
	return GetParam(r, "id")
}
