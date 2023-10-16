import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, RouterOutlet } from '@angular/router';
import { PolizasApiService } from './polizas-api.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'proyectoapi';

  constructor(private router: Router, private polizasApi: PolizasApiService) {}
  ngOnInit(): void {
    if (localStorage.getItem('token') == null) {
      this.obtenerToken();
    }
  }

  agregarEmpleado() {
    this.router.navigate(['/agregarempleado']);
  }

  verEmpleados() {
    this.router.navigate(['/']);
  }

  verArticulos() {
    this.router.navigate(['/verarticulos']);
  }

  obtenerToken() {
    this.polizasApi.postToken().subscribe(
      (response) => {
        localStorage.setItem('token', response.data.token);
        console.log(response);
      },
      (error) => {
        console.error('Error al obtener datos:', error);
      }
    );
  }
}
