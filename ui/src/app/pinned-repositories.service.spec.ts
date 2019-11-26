import { TestBed } from '@angular/core/testing';

import { PinnedRepositoriesService } from './pinned-repositories.service';

describe('PinnedRepositoriesService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: PinnedRepositoriesService = TestBed.get(PinnedRepositoriesService);
    expect(service).toBeTruthy();
  });
});
