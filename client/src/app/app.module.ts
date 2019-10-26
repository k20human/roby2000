import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MaterialModule } from './material.module';
import { DashboardComponent } from './dashboard/dashboard.component';

@NgModule({
    declarations: [
        AppComponent,
        DashboardComponent
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        FontAwesomeModule,
        FlexLayoutModule,
        MaterialModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
