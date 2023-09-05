import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Product } from '../entities/product.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class ProductService {
    url = 'product/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Product[]>{
        return this.http.get<Product[]>(this.url+"findall")
    }

    create(product: Product): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",product)
    }

    update(product: Product): Observable<boolean>{
        return this.http.put<boolean>(this.url+"update",product)
    }   

    delete(id: number): Observable<boolean>{
        return this.http.delete<boolean>(this.url+"delete/" + id)
    }
}