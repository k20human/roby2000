import { NgModule } from '@angular/core';
import {
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatSidenavModule,
    MatToolbarModule
} from '@angular/material';

@NgModule({
    imports: [
        MatCardModule,
        MatIconModule,
        MatToolbarModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatSidenavModule,
        MatListModule
    ],
    exports: [
        MatCardModule,
        MatIconModule,
        MatToolbarModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatSidenavModule,
        MatListModule
    ]
})
export class MaterialModule {}
