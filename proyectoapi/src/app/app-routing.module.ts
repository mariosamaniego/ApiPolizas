import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EmpleadosComponent } from './empleados/empleados.component';
import { AnadirEmpleadoComponent } from './anadir-empleado/anadir-empleado.component';
import { VerArticulosComponent } from './ver-articulos/ver-articulos.component';
import { VerPolizasComponent } from './ver-polizas/ver-polizas.component';
import { VerPolizaComponent } from './ver-poliza/ver-poliza.component';
import { AgregarPolizaComponent } from './agregar-poliza/agregar-poliza.component';

const routes: Routes = [
  { path: 'polizas/:idempleado/:nombre', component: VerPolizasComponent },
  { path: '', redirectTo: '/empleados', pathMatch: 'full' },
  {
    path: 'empleados',
    component: EmpleadosComponent,
  },
  {
    path: 'agregarpoliza/:idempleado/:nombre',
    component: AgregarPolizaComponent,
  },
  { path: 'poliza/:idpoliza', component: VerPolizaComponent },
  { path: 'agregarempleado', component: AnadirEmpleadoComponent },
  { path: 'verarticulos', component: VerArticulosComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
