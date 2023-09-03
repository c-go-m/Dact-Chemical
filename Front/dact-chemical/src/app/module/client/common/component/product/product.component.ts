import { Component } from '@angular/core';
import { Product } from 'src/app/shared/entities/product.model';
import { ProductService } from 'src/app/shared/service/product.service';
import { ConfigService } from '../../../../../shared/config/config.service';


@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.scss']
})
export class ProductComponent {  
  products: Product[] = []
  activeIndex: number = 0
  urlBase:string = ""

  constructor(
    private readonly productService: ProductService,
    private readonly configService: ConfigService
    )
  {
    this.urlBase = this.configService.getAppConfig().urlStorage + "";
  }

  ngOnInit(): void {    
    this.getAllProduct()
  }

  getAllProduct(){
    this.productService.getAll().subscribe(result => {      
      this.products = result      
    })
  }

  validateSelectedPresentation(presentation?: number):boolean{
    return this.activeIndex == Number(presentation) - 1 
  }

  changePresentation(index: number){
    this.activeIndex = index
  }

}
