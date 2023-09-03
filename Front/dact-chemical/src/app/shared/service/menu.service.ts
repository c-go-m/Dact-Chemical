import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Menu } from '../entities/menu.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class MenuService {
    url = 'menu/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Menu[]>{
        return this.http.get<Menu[]>(this.url+"findall")
    }

    create(menu: Menu): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",menu)
    }
}