import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { NodeAddPageRoutingModule } from './node-add-routing.module';

import { NodeAddPage } from './node-add.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    NodeAddPageRoutingModule
  ],
  declarations: [NodeAddPage]
})
export class NodeAddPageModule {}
