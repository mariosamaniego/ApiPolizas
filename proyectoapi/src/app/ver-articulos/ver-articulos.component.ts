import { Component, OnInit } from '@angular/core';
import { PolizasApiService } from '../polizas-api.service';

@Component({
  selector: 'app-ver-articulos',
  templateUrl: './ver-articulos.component.html',
  styleUrls: ['./ver-articulos.component.css'],
})
export class VerArticulosComponent implements OnInit {
  articulosData: any;

  constructor(private polizasApi: PolizasApiService) {}

  ngOnInit(): void {
    this.obtenerArticulos();
  }

  obtenerArticulos() {
    this.polizasApi.obtenerArticulos().subscribe(
      (response) => {
        this.articulosData = response.Data;
        console.log(response);

        if (this.articulosData == 0 || this.articulosData == null) {
          alert('No hay articulos disponibles');
        }
      },
      (error) => {
        alert('No fue posible obtener los articulos');
        console.error('Error al obtener articulos:', error);
      }
    );
  }
}
