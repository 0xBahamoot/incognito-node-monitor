import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { IonicModule } from '@ionic/angular';

import { NodeConfigPage } from './node-config.page';

describe('NodeConfigPage', () => {
  let component: NodeConfigPage;
  let fixture: ComponentFixture<NodeConfigPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NodeConfigPage ],
      imports: [IonicModule.forRoot()]
    }).compileComponents();

    fixture = TestBed.createComponent(NodeConfigPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  }));

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
