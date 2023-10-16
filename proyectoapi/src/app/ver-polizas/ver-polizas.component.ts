import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-ver-polizas',
  templateUrl: './ver-polizas.component.html',
  styleUrls: ['./ver-polizas.component.css'],
})
export class VerPolizasComponent implements OnInit {
  polizasData: any;
  idEmpleado: string = '';
  nombreEmpleado: string = '';
  token: any;

  constructor(
    private polizasApi: PolizasApiService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      const idempleado = params.get('idempleado');
      const nombre = params.get('nombre');
      if (nombre) {
        this.nombreEmpleado = nombre;
      }

      if (idempleado !== null) {
        this.idEmpleado = idempleado;
        this.obtenerPolizas(idempleado);
        this.token = localStorage.getItem('token');
      }
    });
  }

  obtenerPolizas(empleado: string) {
    this.polizasApi.obtenerPolizas(empleado).subscribe(
      (response) => {
        this.polizasData = response.Data.map(
          (item: { Poliza: any }) => item.Poliza
        );
        console.log(response);
      },
      (error) => {
        alert('No fue posible obtener las polizas');
        console.error('Error al obtener las polizas:', error);
      }
    );
  }

  regresar() {
    this.router.navigate([`/`]);
  }

  agregarPoliza() {
    this.router.navigate([
      `/agregarpoliza/${this.idEmpleado}/${this.nombreEmpleado}`,
    ]);
  }

  eliminarEmpleado(IdEmpleado: string, Opcion: string) {
    this.polizasApi.eliminarPoliza(IdEmpleado, Opcion, this.token).subscribe(
      (response) => {
        alert('Empleado eliminado con exito');
        this.router.navigate([`/`]);
        console.log(response);
      },
      (error) => {
        alert('No fue posible eliminar el empleado');
        console.error('Error al eliminar el empleado:', error);
      }
    );
  }
}
