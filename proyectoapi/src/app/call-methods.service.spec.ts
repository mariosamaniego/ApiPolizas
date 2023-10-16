import { TestBed } from '@angular/core/testing';

import { CallMethodsService } from './call-methods.service';

describe('CallMethodsService', () => {
  let service: CallMethodsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CallMethodsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
