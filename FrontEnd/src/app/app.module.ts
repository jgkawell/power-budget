import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { ServiceWorkerModule } from '@angular/service-worker';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { AboutComponent } from './components/pages/about/about.component';
import { AddDebitComponent } from './components/add-debit/add-debit.component';
import { HeaderComponent } from './components/layout/header/header.component';
import { DebitsComponent } from './components/debits/debits.component';
import { DebitItemComponent } from './components/debit-item/debit-item.component';
import { SpeedDialFabComponent } from './components/speed-dial-fab/speed-dial-fab.component';
import { environment } from '../environments/environment';


@NgModule({
  declarations: [
    AppComponent,
    DebitsComponent,
    DebitItemComponent,
    HeaderComponent,
    AddDebitComponent,
    AboutComponent,
    SpeedDialFabComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
    MatButtonModule,
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: environment.production,
    }),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
