package entities

type Poliza struct {
	IdPoliza       int32 `json:"IdPoliza"`
	EmpleadoGenero int32 `json:"EmpleadoGenero"`
	Sku            int32 `json:"Sku"`
	Cantidad       int32 `json:"Cantidad"`
}

type Polizas []Poliza
