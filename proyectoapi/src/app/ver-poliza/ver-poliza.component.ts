import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';
import { Location } from '@angular/common';

@Component({
  selector: 'app-ver-poliza',
  templateUrl: './ver-poliza.component.html',
  styleUrls: ['./ver-poliza.component.css'],
})
export class VerPolizaComponent implements OnInit {
  articulosData: any;
  body: any;
  polizaDataPoliza: any;
  polizaDataEmpleado: any;
  idPoliza: string = '';
  sku: string = '';
  cantidad: string = '';
  token: any;

  constructor(
    private polizasApi: PolizasApiService,
    private route: ActivatedRoute,
    private router: Router,
    private location: Location
  ) {}

  ngOnInit(): void {
    this.consumirArticulos();
    this.route.paramMap.subscribe((params) => {
      const idPoliza = params.get('idpoliza');
      if (idPoliza) {
        this.idPoliza = idPoliza;
        this.obtenerPoliza(idPoliza);
        this.token = localStorage.getItem('token');
      }
    });
  }

  obtenerPoliza(idPoliza: string) {
    this.polizasApi.obtenerPoliza(idPoliza).subscribe(
      (response) => {
        this.polizaDataPoliza = response.Data.Poliza;
        this.polizaDataEmpleado = response.Data.Empleado;
        this.sku = this.polizaDataPoliza.Sku;
        this.cantidad = this.polizaDataPoliza.Cantidad;
        console.log(response);
      },
      (error) => {
        alert('No fue posible obtener la poliza');
        console.error('Error al obtener la poliza:', error);
      }
    );
  }

  actualizarPoliza() {
    const dataPoliza = this.polizaDataPoliza;
    const dataEmpleado = this.polizaDataEmpleado;

    this.body = {
      IdPoliza: dataPoliza.IdPoliza,
      IdEmpleado: dataEmpleado.IdEmpleado,
      Sku: parseInt(dataPoliza.Sku),
      Cantidad: dataPoliza.Cantidad,
      Nombre: dataEmpleado.Nombre,
      Apellido: dataEmpleado.Apellido,
      Puesto: dataEmpleado.Puesto,
    };

    this.polizasApi.actualizarPoliza(this.body, this.token).subscribe(
      (response) => {
        console.log(response);
        if (response.Meta.Status === 'FAIL') {
          alert(response.Data.Respuesta);
          console.log(response);
        } else {
          alert(response.Data.Respuesta);
          this.router.navigate(['/']);
        }
      },
      (error) => {
        alert('Error al actualizar la poliza');
        console.log('Error al actualiza la poliza', error);
      }
    );
  }

  eliminarPoliza(IdPoliza: string, Opcion: string) {
    this.polizasApi.eliminarPoliza(IdPoliza, Opcion, this.token).subscribe(
      (response) => {
        alert('Poliza eliminada con exito');
        const IdEmpleado = this.polizaDataPoliza.EmpleadoGenero;
        const Nombre = this.polizaDataEmpleado.Nombre.trim();
        this.router.navigate([`polizas/${IdEmpleado}/${Nombre}`]);
        console.log(response);
      },
      (error) => {
        alert('No fue posible eliminar la poliza');
        console.error('Error al eliminar la poliza:', error);
      }
    );
  }

  regresar() {
    const IdEmpleado = this.polizaDataPoliza.EmpleadoGenero;
    const Nombre = this.polizaDataEmpleado.Nombre.trim();
    this.router.navigate([`polizas/${IdEmpleado}/${Nombre}`]);
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
}
