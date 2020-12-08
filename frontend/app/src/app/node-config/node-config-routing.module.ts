import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { NodeConfigPage } from './node-config.page';

const routes: Routes = [
  {
    path: '',
    component: NodeConfigPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NodeConfigPageRoutingModule {}
