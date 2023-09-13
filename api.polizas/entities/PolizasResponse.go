package entities

type PolizasResponse struct {
	Poliza   Poliza
	Empleado Empleado
}

type PolizasResponse2 []PolizasResponse

type PolizaMessage struct {
	Message string
}
