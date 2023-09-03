import { NgModule } from '@angular/core';
import { BannerService } from './service/banner.service';
import { CategoryService } from './service/category.service';
import { PresentationService } from './service/presentation.service';
import { ProductService } from './service/product.service';
import { MenuService } from './service/menu.service';
import { AttachedComponent } from './component/index'
import { InformationService } from './service/information.service';


@NgModule({
  declarations: [
    AttachedComponent
  ],
  exports:[
    AttachedComponent
  ],
  imports: [    
  ],
  providers: [
    BannerService,
    CategoryService,
    MenuService,
    PresentationService,
    InformationService,
    ProductService
  ],  
})
export class SharedModule { }