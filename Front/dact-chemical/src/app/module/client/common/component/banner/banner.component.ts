import { Component } from '@angular/core';
import { Banner } from 'src/app/shared/entities/banner.model';
import { BannerService } from 'src/app/shared/service/banner.service';
import { ConfigService } from '../../../../../shared/config/config.service';

@Component({
  selector: 'app-banner',
  templateUrl: './banner.component.html',
  styleUrls: ['./banner.component.scss']
})
export class BannerComponent {  
  banners: Banner[] = []
  activeIndex: number = 1
  urlBase:string = ""
  
  constructor(
      private readonly bannerService: BannerService,
      private readonly configService: ConfigService
    )
  {    
    this.urlBase = this.configService.getAppConfig().urlStorage + "";
  }

  ngOnInit(): void {    
    this.getAllBanner();
  }

  getAllBanner(){
    this.bannerService.getAll().subscribe(result => {
      this.banners = result.sort( function(a,b){ return Number(a.position) - Number(b.position) })           
    })
  }

}
