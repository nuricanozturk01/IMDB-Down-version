import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CelebDetailsComponent } from './celeb-details.component';

describe('CelebDetailsComponent', () => {
  let component: CelebDetailsComponent;
  let fixture: ComponentFixture<CelebDetailsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CelebDetailsComponent]
    });
    fixture = TestBed.createComponent(CelebDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
