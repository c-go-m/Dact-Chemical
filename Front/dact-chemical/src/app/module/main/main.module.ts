import { NgModule, APP_INITIALIZER } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { ConfigService } from '../../shared/config/config.service';
import { InterceptorService } from '../../shared/interceptor/interceptor.service';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ReactiveFormsModule } from '@angular/forms';

import { MainRoutingModule } from './main-routing.module';

import { 
  MainComponent,
  HeaderComponent,
  FooterComponent,
} from './index';

import { ClientModule } from '../client/client.module';
import { AdminModule } from '../admin/admin.module';
import { SharedModule } from '../../shared/shared.module';


@NgModule({
  declarations: [
    MainComponent,
    HeaderComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    MainRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    
    AdminModule,
    SharedModule,
    ClientModule,
    NgbModule,
  ],
  providers: [
    ConfigService,
    {
      provide: APP_INITIALIZER,
      useFactory: (ConfigService: ConfigService) => async () =>
        await ConfigService.init('./assets/config',true),    
      deps: [ConfigService],
      multi: true,
    },
    { provide: HTTP_INTERCEPTORS, 
      useClass: InterceptorService, 
      multi: true 
    }
  ],
  bootstrap: [MainComponent]
})
export class MainModule { }
