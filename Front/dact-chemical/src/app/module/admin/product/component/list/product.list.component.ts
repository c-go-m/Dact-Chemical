import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Product } from 'src/app/shared/entities/product.model';
import { ProductService } from 'src/app/shared/service/product.service';
import { ProductModalComponent } from '../modal/product.modal.component';

@Component({
  selector: 'app-product-list',
  templateUrl: './product.list.component.html',
  styleUrls: ['./product.list.component.scss']
})
export class ProductListComponent implements OnInit {  
  products: Product[] = []

  constructor(
    private readonly productService: ProductService,
    private readonly modalService: NgbModal
    ){}

  ngOnInit(): void {
    this.getAllProduct()
  }

  getAllProduct(){
    this.productService.getAll().subscribe(result => {
      this.products = result
    })
  }

  deleteBanner(product: Product){ 
    this.productService.delete(product.Id ? product.Id : 0).subscribe(result =>{
      this.getAllProduct()
    })
  }

  openModal(product?: Product){
    const modalRef = this.modalService.open(ProductModalComponent, {size: 'lg'})
    
    modalRef.componentInstance.product = product;
    
    modalRef.result.then(() => {
      this.getAllProduct()
    });
  }

}
