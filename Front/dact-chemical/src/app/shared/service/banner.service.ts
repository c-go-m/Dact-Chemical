import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

import { Observable } from 'rxjs';
import { Banner } from '../entities/banner.model';
import { ConfigService } from '../config/config.service';

@Injectable()
export class BannerService {
    url = 'banner/'
    constructor(
        private readonly http: HttpClient,
        private readonly configService: ConfigService
        ){
        this.url = this.configService.getAppConfig().urlBase + this.url
    }

    getAll(): Observable<Banner[]>{
        return this.http.get<Banner[]>(this.url+"findall")
    }

    create(banner: Banner): Observable<boolean>{
        return this.http.post<boolean>(this.url+"create",banner)
    }

    update(banner: Banner): Observable<boolean>{
        return this.http.put<boolean>(this.url+"update",banner)
    }
}