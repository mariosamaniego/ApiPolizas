import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'template-empleados',
  templateUrl: './empleados.component.html',
  styleUrls: ['./empleados.component.css'],
})
export class EmpleadosComponent implements OnInit {
  empleadosData: any;
  tablaVisible: boolean = true;
  seccionOculta: boolean = true;
  textoVisible: boolean = false;
  mostrarFormulario: boolean = false;
  mostrarEncabezado: boolean = true;
  mostrarParrafo: boolean = true;

  constructor(private polizasApi: PolizasApiService, private router: Router) {}

  ngOnInit(): void {
    this.obtenerEmpleados();
  }

  obtenerEmpleados() {
    this.polizasApi.getEmpleados().subscribe(
      (response) => {
        this.empleadosData = response.Data;
        console.log(response);
      },
      (error) => {
        console.error('Error al obtener datos:', error);
      }
    );
  }

  ocultarSeccion() {
    this.seccionOculta = false;
  }

  regresarTabla() {
    this.mostrarParrafo = true;
    this.mostrarEncabezado = true;
    this.tablaVisible = true;
    this.mostrarFormulario = false;
  }
}
