import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CategoryService } from 'src/app/shared/service/category.service';
import { Category } from 'src/app/shared/entities/category.model';
import { Attached } from 'src/app/shared/entities/attached.model';

@Component({
  selector: 'app-category-modal',
  templateUrl: './category.modal.component.html',
  styleUrls: ['./category.modal.component.scss']
})

export class CategoryModalComponent implements OnInit {  
  @Input() public category: Category = {}  
  
  formCategory: FormGroup
  attached: Attached = {}      

  constructor(
    private readonly activeModal: NgbActiveModal,
    private readonly fromGroup: FormBuilder,
    private readonly categoryService: CategoryService 
    )
  {
    this.formCategory = this.initFormGroup(this.category)
  }
  
  ngOnInit(): void {
    this.formCategory = this.initFormGroup(this.category)
  }

  initFormGroup(category: Category): FormGroup {
    return this.fromGroup.group({
      name: [category ? category.name : "" , Validators.required],
      position: [category ? category.position : 0 , Validators.required]
    })    
  }

  async submitAttached(attached: Attached){
    this.attached = attached
  }

  mapObject(){
    if(this.category){
      Object.assign(this.category, this.formCategory.value)
      this.category.attached = this.attached
    }else{      
      this.category = Object.assign({}, this.formCategory.value)
      this.category.attached = this.attached
    }  
  }  

  create(){
    this.mapObject()
        
    this.categoryService.create(this.category).subscribe(() => {
      this.activeModal.close();
    })
  }

  update(){
    this.mapObject()
    
    /*this.categoryService.update(this.category).subscribe(() => {
      this.activeModal.close()
    })*/
  }

  close(){
    this.activeModal.close()
  }

}
