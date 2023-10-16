import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VerArticulosComponent } from './ver-articulos.component';

describe('VerArticulosComponent', () => {
  let component: VerArticulosComponent;
  let fixture: ComponentFixture<VerArticulosComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [VerArticulosComponent]
    });
    fixture = TestBed.createComponent(VerArticulosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
