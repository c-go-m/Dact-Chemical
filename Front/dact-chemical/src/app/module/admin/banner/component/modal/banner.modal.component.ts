import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { Banner } from 'src/app/shared/entities/banner.model';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BannerService } from 'src/app/shared/service/banner.service';
import { Attached } from 'src/app/shared/entities/attached.model';

@Component({
  selector: 'app-banner-modal',
  templateUrl: './banner.modal.component.html',
  styleUrls: ['./banner.modal.component.scss']
})

export class BannerModalComponent implements OnInit {  
  @Input() public banner: Banner = {};
  
  formBanner: FormGroup       
  attached: Attached = {}

  constructor(
    private readonly activeModal: NgbActiveModal,
    private readonly fromGroup: FormBuilder,
    private readonly bannerService: BannerService 
    )
  {
    this.formBanner = this.initFormGroup(this.banner)
  }
  
  ngOnInit(): void {
    this.formBanner = this.initFormGroup(this.banner)
  }

  initFormGroup(banner: Banner): FormGroup {
    return this.fromGroup.group({      
      name: [banner ? banner.name : "" , Validators.required],
      position: [banner ? banner.position : 0 , Validators.required]
    })    
  }

  async submitAttached(attached: Attached){
    this.attached = attached;
  }
  
  mapObject(){
    if(this.banner){
      Object.assign(this.banner, this.formBanner.value)
      this.banner.attached = this.attached
    }else{      
      this.banner = Object.assign({}, this.formBanner.value)
      this.banner.attached = this.attached
    }
  }

  create(){
    this.mapObject()
   
    this.bannerService.create(this.banner).subscribe(() => {
      this.activeModal.close()
    })
      
  }

  update(){
    this.mapObject()
    
    /*this.bannerService.update(this.banner).subscribe(() => {
      this.activeModal.close()
    })*/
  }

  close(){
    this.activeModal.close()
  }

}
