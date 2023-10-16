import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-anadir-empleado',
  templateUrl: './anadir-empleado.component.html',
  styleUrls: ['./anadir-empleado.component.css'],
})
export class AnadirEmpleadoComponent implements OnInit {
  body: any;
  token: any;
  formData = { idempleado: '', nombre: '', apellido: '', puesto: '' };

  constructor(private polizasApi: PolizasApiService, private router: Router) {}

  ngOnInit(): void {
    this.token = localStorage.getItem('token');
  }

  guardarEmpleado() {
    this.body = {
      idempleado: this.formData.idempleado,
      nombre: this.formData.nombre,
      apellido: this.formData.apellido,
      puesto: this.formData.puesto,
    };
    this.polizasApi.guardarEmpleado(this.body, this.token.toString()).subscribe(
      (response) => {
        if (response.Meta.Status === 'FAIL') {
          alert(response.Data.Respuesta);
        } else {
          alert(response.Data.Respuesta);
          this.router.navigate(['/']);
        }
      },
      (error) => {
        console.error('Error al guardar empleado', error);
      }
    );
  }
}
