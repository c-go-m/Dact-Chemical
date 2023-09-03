import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Category } from '../entities/category.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class CategoryService {
    url = 'category/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Category[]>{
        return this.http.get<Category[]>(this.url+"findall")
    }

    create(category: Category): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",category)
    }

    update(category: Category): Observable<boolean>{
        return this.http.put<boolean>(this.url+"update",category)
    }    
}