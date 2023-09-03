import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Attached } from '../../entities/attached.model';

@Component({
  selector: 'app-attached',
  templateUrl: './attached.component.html',
  styleUrls: ['./attached.component.scss']
})

export class AttachedComponent implements OnInit {  
  
  @Input() attached?: Attached = {}; 
  @Output() submitAttached: EventEmitter<Attached> = new EventEmitter(); 
  
  constructor()
  {
    

  }
  
  ngOnInit(): void {

  }

  async getAttached(event: any){
    if(event.files.length == 1){
      var nameFile = event.files[0].name.split(".")
      var promise = this.getBase64(event.files[0])
      var data:any = await promise

      this.setAttached(nameFile,data)
    }
  }
  
  setAttached(nameFile:string[], data:any){
    if(this.attached) {
      this.attached.name = nameFile[0]
      this.attached.extension = nameFile[1]
      this.attached.data = data.split(",")[1]  
      this.attached.url = ""
  
      this.submitAttached.emit(this.attached)     
    }
  }

  getBase64(file:any) {
    return new Promise(function(resolve, reject) {
        var reader = new FileReader();
        reader.onload = function() { resolve(reader.result); };
        reader.onerror = reject;
        reader.readAsDataURL(file);
    });    
  }

}
