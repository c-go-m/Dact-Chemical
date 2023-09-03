import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { ConfigService } from './config.service';
import { mockConfig, mockUrl, mockUrlBase, mockNotPath } from '../../mock/mockConfig';
import { Config } from './config.model';



describe('Test ConfigService', () => {

    let httpClient: HttpClient
    let httpTestingController: HttpTestingController
    let service: ConfigService

    beforeEach(() => {
      TestBed.configureTestingModule({
        imports: [HttpClientTestingModule],
        providers: [
          ConfigService
        ]
      })

      httpClient = TestBed.inject(HttpClient)
      service = TestBed.inject(ConfigService)
      httpTestingController = TestBed.inject(HttpTestingController)

    })
  
    afterEach(() => httpTestingController.verify())
  
    it('#Test Create Instance', () => {
      expect(service).toBeDefined()
    })

    it('#Test Method init config null is path', () => {   
      service.init(mockUrlBase,true).then(()=>{
        expect(ConfigService.appConfig).toEqual(mockConfig)    
      });

      let req = httpTestingController
              .expectOne(mockUrl);

      req.flush(mockConfig)              
    })

    it('#Test Method init config null not path', () => {   
      ConfigService.appConfig = new Config()
      service.init(mockNotPath,false)
      expect(ConfigService.appConfig).toEqual(mockConfig)          
    })

    it('#Test Method getAppConfig', () => {
      ConfigService.appConfig = mockConfig
      expect(service.getAppConfig()).toEqual(ConfigService.appConfig)    
    })

    it('#Test Method init exist config', () => {     
      ConfigService.appConfig = mockConfig 
      expect(service.init(mockUrlBase,true)).toBeDefined()    
    })
  })