import { NgModule } from '@angular/core';
import { ClientRoutingModule } from './client.routing.module';
import { GalleriaModule } from 'primeng/galleria';
import { HomeComponent } from './home/home.component';
import { ProductListComponent } from './product-list/product-list.component';
import { CategoryComponent, ProductComponent, BannerComponent, InformationComponent } from './common/index';

@NgModule({
  declarations: [
    BannerComponent,
    HomeComponent,
    CategoryComponent,
    ProductComponent,
    InformationComponent,
    ProductListComponent
  ],
  imports: [    
    ClientRoutingModule,
    GalleriaModule
  ],
  providers: [],  
})
export class ClientModule { }
