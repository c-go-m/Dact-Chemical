import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { config, lastValueFrom, Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Config } from './config.model';

@Injectable()

export class ConfigService {    
    public static appConfig: Config;


    constructor(private readonly http: HttpClient){    }

    public async init(pathDeployJson: string, isPath: boolean)
    {         
      if ( (!ConfigService.appConfig) || (!ConfigService.appConfig.urlBase) ) {                  
          await this.loadConfig(pathDeployJson, isPath);       
        }
    }

    public getAppConfig():Config{
      return ConfigService.appConfig
    }
    
    private async loadConfig(pathConfigData: string, isPath: boolean) {      
      if(isPath){
        let loading = this.getJSON(pathConfigData)   
        ConfigService.appConfig = await lastValueFrom(loading);        
      }else{               
        ConfigService.appConfig = JSON.parse(pathConfigData);
      }
    }
    
    private getJSON(pathConfigData: string): Observable<Config> {
      const jsonFile = `${pathConfigData}/config.${environment.appConfig}.json`;
      return this.http.get<Config>(jsonFile);
    }  
}