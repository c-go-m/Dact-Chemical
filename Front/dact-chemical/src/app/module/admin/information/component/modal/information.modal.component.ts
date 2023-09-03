import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { Information } from 'src/app/shared/entities/information.model';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InformationService } from 'src/app/shared/service/information.service';
import { MenuService } from 'src/app/shared/service/menu.service';
import { Attached } from 'src/app/shared/entities/attached.model';
import { Menu } from 'src/app/shared/entities/menu.model';

@Component({
  selector: 'app-information-modal',
  templateUrl: './information.modal.component.html',
  styleUrls: ['./information.modal.component.scss']
})

export class InformationModalComponent implements OnInit {  
  @Input() public information: Information = {};
  
  formInformation: FormGroup    
  menus: Menu[] = []
  attached: Attached = {}

  constructor(
    private readonly activeModal: NgbActiveModal,
    private readonly fromGroup: FormBuilder,
    private readonly informationService: InformationService,
    private readonly menuService: MenuService
    )
  {
    this.formInformation = this.initFormGroup(this.information)
  }
  
  ngOnInit(): void {
    this.getAllMenu()
    this.formInformation = this.initFormGroup(this.information)
  }
  
  getAllMenu(){
    this.menuService.getAll().subscribe(result => {
      this.menus = result
    })
  }  

  initFormGroup(information: Information): FormGroup {
    return this.fromGroup.group({      
      name: [information ? information.name : "" , Validators.required],
      value: [information ? information.value : "" , Validators.required],
      position: [information ? information.position : 0 , Validators.required],
      menuId: [information ? information.menuId : 0 , Validators.required],
    })    
  }

  async submitAttached(attached: Attached){
    this.attached = attached;
  }
  
  mapObject(){
    if(this.information){
      Object.assign(this.information, this.formInformation.value)
      this.information.attached = this.attached
    }else{      
      this.information = Object.assign({}, this.formInformation.value)
      this.information.attached = this.attached
    }
  }

  create(){
    this.mapObject()
   
    this.informationService.create(this.information).subscribe(() => {
      this.activeModal.close()
    })
      
  }

  update(){
    this.mapObject()
    
    /*this.informationService.update(this.information).subscribe(() => {
      this.activeModal.close()
    })*/
  }

  close(){
    this.activeModal.close()
  }

}
