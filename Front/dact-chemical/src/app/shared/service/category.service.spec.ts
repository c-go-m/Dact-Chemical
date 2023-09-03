import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { CategoryService } from './category.service';
import { mockListCategory } from 'src/app/mock/mockCategory';
import { mockNotPath, mockUrlBase } from 'src/app/mock/mockConfig';
import { ConfigService } from '../config/config.service';
import { APP_INITIALIZER } from '@angular/core';

describe('Test ConfigService', () => {

    let httpClient: HttpClient
    let httpTestingController: HttpTestingController
    let service: CategoryService

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
          CategoryService
        ]
      })

      httpClient = TestBed.inject(HttpClient)
      service = TestBed.inject(CategoryService)
      httpTestingController = TestBed.inject(HttpTestingController)

    })
  
    afterEach(() => httpTestingController.verify())
  
    it('#Test Create Instance', () => {
      expect(service).toBeDefined()
    })

    it('#Test Method getAll', () => {   
      service.getAll().subscribe( resutl => {
        expect(resutl).toEqual(mockListCategory)           
      });

      let req = httpTestingController
              .expectOne(mockUrlBase+"/category/findall");
      req.flush(mockListCategory)              
    })

    it('#Test Method create', () => {   
      service.create(mockListCategory[0]).subscribe( resutl => {
        expect(resutl).toEqual(true)    
      });

      let req = httpTestingController
              .expectOne(mockUrlBase+"/category/create");
      req.flush(true)              
    })

  })

