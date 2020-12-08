import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { NodeDetailPage } from './node-detail.page';

const routes: Routes = [
  {
    path: '',
    component: NodeDetailPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NodeDetailPageRoutingModule {}
