import { Component } from '@angular/core';


@Component({
  selector: 'app-service',
  templateUrl: './service.component.html',
  styleUrls: ['./service.component.scss']
})
export class ServiceComponent {  
  rute: string =""
  constructor()
  {}

  ngOnInit(): void {    
    this.rute = window.location.pathname
  }  
}