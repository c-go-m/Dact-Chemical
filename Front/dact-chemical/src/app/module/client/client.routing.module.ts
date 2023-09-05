import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { ServiceComponent } from './service/service.component';
import { ProductListComponent } from './product-list/product-list.component';

const routes: Routes = [  
  { path: 'initial', component: HomeComponent },
  { path: 'product', component: ProductListComponent },
  { path: 'service', component: HomeComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ClientRoutingModule { }
