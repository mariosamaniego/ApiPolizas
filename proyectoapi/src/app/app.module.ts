import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { EmpleadosComponent } from './empleados/empleados.component';
import { HttpClientModule } from '@angular/common/http';

import { FormsModule } from '@angular/forms';
import { AnadirEmpleadoComponent } from './anadir-empleado/anadir-empleado.component';
import { VerArticulosComponent } from './ver-articulos/ver-articulos.component';
import { VerPolizasComponent } from './ver-polizas/ver-polizas.component';
import { AgregarPolizaComponent } from './agregar-poliza/agregar-poliza.component';
import { ActualizarEmpleadoComponent } from './actualizar-empleado/actualizar-empleado.component';
import { VerPolizaComponent } from './ver-poliza/ver-poliza.component';

@NgModule({
  declarations: [
    AppComponent,
    EmpleadosComponent,
    AnadirEmpleadoComponent,
    VerArticulosComponent,
    VerPolizasComponent,
    AgregarPolizaComponent,
    ActualizarEmpleadoComponent,
    VerPolizaComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
