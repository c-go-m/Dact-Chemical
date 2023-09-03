import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { Presentation } from 'src/app/shared/entities/presentation.model';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PresentationService } from 'src/app/shared/service/presentation.service';
import { Attached } from 'src/app/shared/entities/attached.model';
import { ProductService } from 'src/app/shared/service/product.service';
import { Product } from 'src/app/shared/entities/product.model';

@Component({
  selector: 'app-presentation-modal',
  templateUrl: './presentation.modal.component.html',
  styleUrls: ['./presentation.modal.component.scss']
})

export class PresentationModalComponent implements OnInit {  
  @Input() public presentation: Presentation = {}
  
  formPresentation: FormGroup    
  products: Product[] = []
  attached: Attached = {}

  constructor(
    private readonly activeModal: NgbActiveModal,
    private readonly fromGroup: FormBuilder,
    private readonly presentationService: PresentationService,
    private readonly productService: ProductService  
    )
  {
    this.formPresentation = this.initFormGroup(this.presentation)
  }
  
  ngOnInit(): void {
    this.getAllProduct()
    this.formPresentation = this.initFormGroup(this.presentation)
  }

  getAllProduct(){
    this.productService.getAll().subscribe(result => {
      this.products = result
    })
  }  

  initFormGroup(presentation: Presentation): FormGroup {
    return this.fromGroup.group({
      name: [presentation ? presentation.name : "" , Validators.required],
      position: [presentation ? presentation.position : 0 , Validators.required],
      productId: [presentation ? presentation.productId : 0 , Validators.required]
    })    
  }

  async submitAttached(attached: Attached){
    this.attached = attached
  }

  mapObject(){
    if(this.presentation){
      Object.assign(this.presentation, this.formPresentation.value)
      this.presentation.attached = this.attached
      this.presentation.product = {}
    }else{      
      this.presentation = Object.assign({}, this.formPresentation.value)
      this.presentation.attached = this.attached
    }    
  }
  
  create(){
    this.mapObject()        
    this.presentationService.create(this.presentation).subscribe(() => {
      this.activeModal.close();
    })
  }

  update(){
    this.mapObject()
    /*this.presentationService.update(this.presentation).subscribe(() => {
      this.activeModal.close()
    })*/
  }

  close(){
    this.activeModal.close()
  }

}
