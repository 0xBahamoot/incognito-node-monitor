import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'folder/Inbox',
    pathMatch: 'full'
  },
  {
    path: 'folder/:id',
    loadChildren: () => import('./folder/folder.module').then( m => m.FolderPageModule)
  },
  {
    path: 'node-detail',
    loadChildren: () => import('./node-detail/node-detail.module').then( m => m.NodeDetailPageModule)
  },
  {
    path: 'node-add',
    loadChildren: () => import('./node-add/node-add.module').then( m => m.NodeAddPageModule)
  },
  {
    path: 'node-config',
    loadChildren: () => import('./node-config/node-config.module').then( m => m.NodeConfigPageModule)
  },
  {
    path: 'setting',
    loadChildren: () => import('./setting/setting.module').then( m => m.SettingPageModule)
  }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule {}
