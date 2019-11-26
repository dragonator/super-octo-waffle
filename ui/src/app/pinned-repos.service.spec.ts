import { TestBed } from '@angular/core/testing';

import { PinnedReposService } from './pinned-repos.service';

describe('PinnedReposService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: PinnedReposService = TestBed.get(PinnedReposService);
    expect(service).toBeTruthy();
  });
});
