import { Component } from '@angular/core';
import { Category } from 'src/app/shared/entities/category.model';
import { CategoryService } from 'src/app/shared/service/category.service';
import { ConfigService } from '../../../../../shared/config/config.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.scss']
})
export class CategoryComponent {  
  categories: Category[] = []
  urlBase:string = ""
  
  constructor(
    private readonly categoryService: CategoryService,
    private readonly configService: ConfigService
    )
  {
    this.urlBase = this.configService.getAppConfig().urlStorage + "";
  }

  ngOnInit(): void {    
    this.getAllCategory();
  }

  getAllCategory(){
    this.categoryService.getAll().subscribe(result => {      
      this.categories = result.sort(function(a,b){return (a.position?a.position:1) - (b.position?b.position:2)})
    })
  }

}
