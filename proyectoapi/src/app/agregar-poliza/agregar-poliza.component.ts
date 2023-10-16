import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-agregar-poliza',
  templateUrl: './agregar-poliza.component.html',
  styleUrls: ['./agregar-poliza.component.css'],
})
export class AgregarPolizaComponent implements OnInit {
  articulosData: any;
  body: any;
  formData: formData = new formData();
  Nombre: string = '';
  EmpleadoGenero: string = '';
  Sku: string = '';
  Cantidad: string = '';
  IdPoliza: string = '';
  token: any;

  constructor(
    private polizasApi: PolizasApiService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.consumirArticulos();
    this.route.paramMap.subscribe((params) => {
      const idEmpleado = params.get('idempleado');
      const nombreEmpleado = params.get('nombre');

      if (idEmpleado && nombreEmpleado) {
        this.token = localStorage.getItem('token');
        this.EmpleadoGenero = idEmpleado;
        this.Nombre = nombreEmpleado;
      }
    });
  }

  guardarPoliza() {
    this.formData.IdPoliza = parseInt(this.IdPoliza);
    this.formData.EmpleadoGenero = parseInt(this.EmpleadoGenero);
    this.formData.Sku = parseInt(this.Sku);
    this.formData.Cantidad = parseInt(this.Cantidad);

    const data = this.formData;

    this.body = {
      IdPoliza: data.IdPoliza,
      EmpleadoGenero: data.EmpleadoGenero,
      Sku: data.Sku,
      Cantidad: data.Cantidad,
    };

    this.polizasApi.guardarPoliza(this.body, this.token).subscribe(
      (response) => {
        console.log(response);
        if (response.Meta.Status === 'FAIL') {
          alert(response.Data.Respuesta);
        } else {
          alert(response.Data.Respuesta);
          this.router.navigate([
            `polizas/${this.EmpleadoGenero}/${this.Nombre}`,
          ]);
        }
      },
      (error) => {
        console.error('Error al guardar la poliza', error);
      }
    );
  }

  consumirArticulos() {
    this.polizasApi.obtenerArticulos().subscribe(
      (response) => {
        this.articulosData = response.Data;
        console.log(this.articulosData);
      },
      (error) => {}
    );
  }

  regresar() {
    this.router.navigate([`polizas/${this.EmpleadoGenero}/${this.Nombre}`]);
  }
}

export class formData {
  IdPoliza: number = 0;
  EmpleadoGenero: number = 0;
  Sku: number = 0;
  Cantidad: number = 0;
}
