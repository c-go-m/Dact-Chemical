import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';

import { AdminRoutingModule } from './admin.routing.module';
import { BannerListComponent, BannerModalComponent } from './banner/index';
import { CategoryListComponent, CategoryModalComponent } from './category/index';
import { PresentationListComponent, PresentationModalComponent } from './presentation/index';
import { ProductListComponent, ProductModalComponent  } from './product/index';
import { InformationListComponent, InformationModalComponent  } from './information/index';


import { SharedModule } from 'src/app/shared/shared.module';


@NgModule({
  declarations: [
    BannerListComponent,
    BannerModalComponent,
    CategoryListComponent,
    CategoryModalComponent,
    PresentationListComponent,
    PresentationModalComponent,
    ProductListComponent,
    ProductModalComponent,
    InformationListComponent,
    InformationModalComponent
  ],
  imports: [    
    ReactiveFormsModule,
    AdminRoutingModule,
    SharedModule,
    CommonModule
  ],
  providers: [],  
})
export class AdminModule { }
