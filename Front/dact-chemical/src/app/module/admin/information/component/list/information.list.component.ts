import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Information } from 'src/app/shared/entities/information.model';
import { InformationService } from 'src/app/shared/service/information.service';
import { InformationModalComponent } from '../modal/information.modal.component';

@Component({
  selector: 'app-information-list',
  templateUrl: './information.list.component.html',
  styleUrls: ['./information.list.component.scss']
})
export class InformationListComponent implements OnInit {  
  informations: Information[] = []

  constructor(
    private readonly informationService: InformationService,
    private readonly modalService: NgbModal
    ){}

  ngOnInit(): void {
    this.getAllInformation()
  }

  getAllInformation(){
    this.informationService.getAll().subscribe(result => {
      this.informations = result
    })
  }

  deleteInformation(information: Information){ 
    this.informationService.delete(information.Id ? information.Id : 0).subscribe(result =>{
      this.getAllInformation()
    })
  }

  openModal(information?: Information){
    const modalRef = this.modalService.open(InformationModalComponent, {size: 'lg'})
    
    modalRef.componentInstance.information = information;
    
    modalRef.result.then(() => {
      this.getAllInformation()
    });
  }

}
