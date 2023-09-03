import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Presentation } from '../entities/presentation.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class PresentationService {
    url = 'presentation/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Presentation[]>{
        return this.http.get<Presentation[]>(this.url+"findall")
    }

    create(presentation: Presentation): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",presentation)
    }

    update(presentation: Presentation): Observable<boolean>{
        return this.http.put<boolean>(this.url+"update",presentation)
    }   
}