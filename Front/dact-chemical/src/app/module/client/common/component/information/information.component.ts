import { Component, Input } from '@angular/core';
import { Information } from 'src/app/shared/entities/information.model';
import { InformationService } from 'src/app/shared/service/information.service';
import { ConfigService } from '../../../../../shared/config/config.service';

@Component({
  selector: 'app-information',
  templateUrl: './information.component.html',
  styleUrls: ['./information.component.scss']
})
export class InformationComponent {  
  @Input() rute:string = ''
  informations: Information[] = []
  urlBase:string = ""
  
  constructor(
    private readonly informationService: InformationService,
    private readonly configService: ConfigService
    )
  {
    this.urlBase = this.configService.getAppConfig().urlStorage + "";
  }

  ngOnInit(): void {    
    this.getAllInformation();
  }

  getAllInformation(){
    this.informationService.getAll().subscribe(result => {           
      result = result.filter(info => info.menu?.rute == this.rute)
      this.informations = result.sort(function(a,b){return (a.position?a.position:1) - (b.position?b.position:2)})
    })
  }

}
