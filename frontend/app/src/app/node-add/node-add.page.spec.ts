import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { NodeAddPage } from './node-add.page';

describe('NodeAddPage', () => {
  let component: NodeAddPage;
  let fixture: ComponentFixture<NodeAddPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NodeAddPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(NodeAddPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
