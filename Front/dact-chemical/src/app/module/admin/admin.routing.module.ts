import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BannerListComponent } from './banner/index';
import { CategoryListComponent } from './category/index';
import { PresentationListComponent } from './presentation/index';
import { ProductListComponent } from './product/index';
import { InformationListComponent } from './information/component/list/information.list.component';

const routes: Routes = [
  { path: '', redirectTo: '/admin/banner', pathMatch: 'full' },
  { path: 'banner', component: BannerListComponent },
  { path: 'category', component: CategoryListComponent },
  { path: 'product', component: ProductListComponent },
  { path: 'presentation', component: PresentationListComponent },
  { path: 'information', component: InformationListComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
