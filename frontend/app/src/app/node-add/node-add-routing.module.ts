import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { NodeAddPage } from './node-add.page';

const routes: Routes = [
  {
    path: '',
    component: NodeAddPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NodeAddPageRoutingModule {}
