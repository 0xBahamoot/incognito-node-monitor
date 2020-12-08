import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { NodeConfigPageRoutingModule } from './node-config-routing.module';

import { NodeConfigPage } from './node-config.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    NodeConfigPageRoutingModule
  ],
  declarations: [NodeConfigPage]
})
export class NodeConfigPageModule {}
