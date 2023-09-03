import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Banner } from 'src/app/shared/entities/banner.model';
import { BannerService } from 'src/app/shared/service/banner.service';
import { BannerModalComponent } from '../modal/banner.modal.component';

@Component({
  selector: 'app-banner-list',
  templateUrl: './banner.list.component.html',
  styleUrls: ['./banner.list.component.scss']
})
export class BannerListComponent implements OnInit {  
  banners: Banner[] = []

  constructor(
    private readonly bannerService: BannerService,
    private readonly modalService: NgbModal
    ){}

  ngOnInit(): void {
    this.getAllBanner()
  }

  getAllBanner(){
    this.bannerService.getAll().subscribe(result => {
      this.banners = result
    })
  }

  openModal(banner?: Banner){
    const modalRef = this.modalService.open(BannerModalComponent, {size: 'lg'})
    
    modalRef.componentInstance.banner = banner;
    
    modalRef.result.then(() => {
      this.getAllBanner()
    });
  }

}
