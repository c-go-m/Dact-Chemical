import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Presentation } from 'src/app/shared/entities/presentation.model';
import { PresentationService } from 'src/app/shared/service/presentation.service';
import { PresentationModalComponent } from '../modal/presentation.modal.component';

@Component({
  selector: 'app-presentation-list',
  templateUrl: './presentation.list.component.html',
  styleUrls: ['./presentation.list.component.scss']
})
export class PresentationListComponent implements OnInit {  
  presentations: Presentation[] = []

  constructor(
    private readonly presentationService: PresentationService,
    private readonly modalService: NgbModal
    ){}

  ngOnInit(): void {
    this.getAllPresentation()
  }

  getAllPresentation(){
    this.presentationService.getAll().subscribe(result => {
      this.presentations = result
    })
  }

  deletePresentation(presentation: Presentation){ 
    this.presentationService.delete(presentation.Id ? presentation.Id : 0).subscribe(result =>{
      this.getAllPresentation()
    })
  }

  openModal(presentation?: Presentation){
    const modalRef = this.modalService.open(PresentationModalComponent, {size: 'lg'})
    
    modalRef.componentInstance.presentation = presentation;
    
    modalRef.result.then(() => {
      this.getAllPresentation()
    });
  }

}
