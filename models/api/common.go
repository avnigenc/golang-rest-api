package models

type HealthResponse struct {
	Message string
	Time    int64
}

type GenericResponse struct {
	Result interface{}
	Error  interface{}
}
