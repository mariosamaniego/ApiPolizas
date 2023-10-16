import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CallMethodsService } from './call-methods.service';
import { HttpHeaders, HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class PolizasApiService {
  private jsonToken = {
    appId: '58ccba34-6382-481d-b87f-3fe7d95d430e',
    appKey: '53d16da74313af15c29d5a486390a572e6255d4855fd7405d7b017a4de06bf76',
  };
  private apiUrl = 'http://10.59.21.109:3000/api/v1';
  HEADER: any;
  constructor(
    private callMethods: CallMethodsService,
    private http: HttpClient
  ) {}

  getEmpleados(): Observable<any> {
    return this.callMethods.get(`${this.apiUrl}/Empleados`);
  }

  obtenerPolizas(empleado: string): Observable<any> {
    return this.callMethods.get(
      `${this.apiUrl}/PolizaEmpelado?idempleado=${empleado}`
    );
  }

  obtenerPoliza(idPoliza: string): Observable<any> {
    return this.callMethods.get(
      `${this.apiUrl}/ConsultarPoliza?idpoliza=${idPoliza}`
    );
  }

  eliminarPoliza(
    idPoliza: string,
    opcion: string,
    token: string
  ): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        Authorization: token,
      }),
    };
    const url = `${this.apiUrl}/Eliminar?opcion=${opcion}&eliminar=${idPoliza}`;
    return this.callMethods.post(url, '', this.HEADER);
  }

  eliminarEmpleado(
    idEmpleado: string,
    opcion: string,
    token: string
  ): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        Authorization: token,
      }),
    };
    const url = `${this.apiUrl}/Eliminar?opcion=${opcion}&eliminar=${idEmpleado}`;
    return this.callMethods.post(url, '', this.HEADER);
  }

  guardarPoliza(body: string, token: string): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        Authorization: token,
      }),
    };
    const url = `${this.apiUrl}/AgregarPolizas`;
    return this.callMethods.post(url, body, this.HEADER);
  }

  guardarEmpleado(body: string, token: string): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        Authorization: token,
      }),
    };
    const url = `${this.apiUrl}/AgregarEmpleado`;
    return this.callMethods.post(url, body, this.HEADER);
  }

  obtenerArticulos(): Observable<any> {
    const url = `${this.apiUrl}/Articulos`;
    return this.callMethods.get(url);
  }

  actualizarPoliza(body: string, token: string): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        Authorization: token,
      }),
    };

    const url = `${this.apiUrl}/ActualizarPoliza`;
    return this.callMethods.post(url, body, this.HEADER);
  }

  postToken(): Observable<any> {
    this.HEADER = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };

    return this.http.post(
      'https://apigateway.coppel.com:58443/sso-dev/api/v1/app/authenticate',
      this.jsonToken,
      this.HEADER
    );
  }
}
