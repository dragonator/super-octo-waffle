import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PinnedReposComponent } from './pinned-repos.component';

describe('PinnedReposComponent', () => {
  let component: PinnedReposComponent;
  let fixture: ComponentFixture<PinnedReposComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PinnedReposComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PinnedReposComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
