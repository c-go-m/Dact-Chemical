import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { BannerService } from './banner.service';
import { mockListBanner } from 'src/app/mock/mockBanner';
import { mockNotPath, mockUrlBase } from 'src/app/mock/mockConfig';
import { ConfigService } from '../config/config.service';
import { APP_INITIALIZER } from '@angular/core';

describe('Test ConfigService', () => {

    let httpClient: HttpClient
    let httpTestingController: HttpTestingController
    let service: BannerService

    beforeEach(() => {
      TestBed.configureTestingModule({
        imports: [
          HttpClientTestingModule          
        ],        
        providers: [
          ConfigService,
          {
            provide: APP_INITIALIZER,
            useFactory: (ConfigService: ConfigService) => async () =>
              await ConfigService.init(mockNotPath,false),    
            deps: [ConfigService],
            multi: true,
          },
          BannerService
        ]
      })

      httpClient = TestBed.inject(HttpClient)
      service = TestBed.inject(BannerService)
      httpTestingController = TestBed.inject(HttpTestingController)

    })
  
    afterEach(() => httpTestingController.verify())
  
    it('#Test Create Instance', () => {
      expect(service).toBeDefined()
    })

    it('#Test Method getAll', () => {   
      service.getAll().subscribe( resutl => {
        expect(resutl).toEqual(mockListBanner)           
      });

      let req = httpTestingController
              .expectOne(mockUrlBase+"/banner/findall");
      req.flush(mockListBanner)              
    })

    it('#Test Method create', () => {   
      service.create(mockListBanner[0]).subscribe( resutl => {
        expect(resutl).toEqual(true)    
      });

      let req = httpTestingController
              .expectOne(mockUrlBase+"/banner/create");
      req.flush(true)              
    })

  })

