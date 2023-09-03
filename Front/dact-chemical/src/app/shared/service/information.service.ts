import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Information } from '../entities/information.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class InformationService {
    url = 'information/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Information[]>{
        return this.http.get<Information[]>(this.url+"findall")
    }

    create(presentation: Information): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",presentation)
    }

    update(presentation: Information): Observable<boolean>{
        return this.http.put<boolean>(this.url+"update",presentation)
    }   
}