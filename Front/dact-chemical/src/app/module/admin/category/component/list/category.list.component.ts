import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Category } from 'src/app/shared/entities/category.model';
import { CategoryService } from 'src/app/shared/service/category.service';
import { CategoryModalComponent } from '../modal/category.modal.component';

@Component({
  selector: 'app-category-list',
  templateUrl: './category.list.component.html',
  styleUrls: ['./category.list.component.scss']
})
export class CategoryListComponent implements OnInit {  
  categories: Category[] = []

  constructor(
    private readonly categoryService: CategoryService,
    private readonly modalService: NgbModal
    ){}

  ngOnInit(): void {
    this.getAllCategory()
  }

  getAllCategory(){
    this.categoryService.getAll().subscribe(result => {
      this.categories = result
    })
  }

  deleteCaregory(category:Category){ 
    this.categoryService.delete(category.Id ? category.Id : 0).subscribe(result =>{
      this.getAllCategory()
    })
  }

  openModal(category?: Category){
    const modalRef = this.modalService.open(CategoryModalComponent, {size: 'lg'})
    
    modalRef.componentInstance.category = category;
    
    modalRef.result.then(() => {
      this.getAllCategory()
    });
  }

}
