import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { NodeDetailPageRoutingModule } from './node-detail-routing.module';

import { NodeDetailPage } from './node-detail.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    NodeDetailPageRoutingModule
  ],
  declarations: [NodeDetailPage]
})
export class NodeDetailPageModule {}
