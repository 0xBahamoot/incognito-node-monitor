import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { NodeDetailPage } from './node-detail.page';

describe('NodeDetailPage', () => {
  let component: NodeDetailPage;
  let fixture: ComponentFixture<NodeDetailPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NodeDetailPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(NodeDetailPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
