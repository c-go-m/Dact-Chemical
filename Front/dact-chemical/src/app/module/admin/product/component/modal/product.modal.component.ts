import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { Product } from 'src/app/shared/entities/product.model';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductService } from 'src/app/shared/service/product.service';
import { CategoryService } from 'src/app/shared/service/category.service';
import { Category } from 'src/app/shared/entities/category.model';
import { Attached } from 'src/app/shared/entities/attached.model';

@Component({
  selector: 'app-product-modal',
  templateUrl: './product.modal.component.html',
  styleUrls: ['./product.modal.component.scss']
})

export class ProductModalComponent implements OnInit {  
  @Input() public product: Product = {}
  
  formProduct: FormGroup    
  categories: Category[] = []
  attached: Attached = {}

  constructor(
    private readonly activeModal: NgbActiveModal,
    private readonly fromGroup: FormBuilder,
    private readonly productService: ProductService,
    private readonly categoryService: CategoryService 
    )
  {
    this.formProduct = this.initFormGroup(this.product)
    
  }
  
  ngOnInit(): void {
    this.getAllCategory()
    this.formProduct = this.initFormGroup(this.product)
  }

  getAllCategory(){
    this.categoryService.getAll().subscribe(result => {
      this.categories = result
    })
  }

  initFormGroup(product: Product): FormGroup {
    return this.fromGroup.group({
      name: [product ? product.name : "" , Validators.required], 
      categoryId: [product ? product.categoryId : 0 , Validators.required]
    })    
  }

  async submitAttached(attached: Attached){
    this.attached = attached;
  }

  mapObject(){
    if(this.product){
      Object.assign(this.product, this.formProduct.value)
      this.product.attached = this.attached
      this.product.category = {}
    }else{      
      this.product = Object.assign({}, this.formProduct.value)
      this.product.attached = this.attached
    }
  } 

  create(){
    this.mapObject()
    this.productService.create(this.product).subscribe(() => {
      this.activeModal.close();
    })
  }

  update(){
    this.mapObject()
    /*this.productService.update(this.product).subscribe(() => {
      this.activeModal.close()
    })*/
  }

  close(){
    this.activeModal.close()
  }

}
