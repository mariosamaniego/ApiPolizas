package entities

type Empleado struct {
	IdEmpleado int32  `json:"IdEmpleado"`
	Nombre     string `json:"Nombre"`
	Apellido   string `json:"Apellido"`
	Puesto     string `json:"Puesto"`
}

type Empleados []Empleado

type Mensaje struct {
	Respuesta       string
	StatusRespuesta int32
}
