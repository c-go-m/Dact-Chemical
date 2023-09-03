import { Component, OnInit } from '@angular/core';
import { Menu } from 'src/app/shared/entities/menu.model';
import { MenuService } from 'src/app/shared/service/menu.service';
import { ConfigService } from '../../../../shared/config/config.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})


export class HeaderComponent implements OnInit {
  urlLogo = ""
  menus : Menu[] = []
  menusUser : Menu[] = []
  userAdmin : boolean = false 

  constructor(
      private readonly menuService:MenuService,
      private readonly configService: ConfigService
    ) 
  {        
    this.urlLogo = this.configService.getAppConfig().urlStorage + "assets/img/company/logo.svg";
  }

  ngOnInit(): void {  
    this.getMenu()
  }

  loginUser(){
    this.userAdmin = true
    this.validMenuUser()
  }

  logoutUser(){
    this.userAdmin = false
    this.validMenuUser()
  }

  getMenu(){
    this.menuService.getAll().subscribe(result => {      
      this.menus = result.sort(function(a,b){return Number(a.position) - Number(b.position) } )      
      this.validMenuUser()
    })
  }

  validMenuUser(){
    if(this.userAdmin){
      this.menusUser = this.menus.filter(m => m.type == "admin")
    }else{
      this.menusUser = this.menus.filter(m => m.type == "client")
    }    
  }
}

