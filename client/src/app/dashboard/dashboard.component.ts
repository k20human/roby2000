import { Component, OnInit } from '@angular/core';
import { faArrowLeft, faArrowRight, faArrowUp, faArrowDown } from '@fortawesome/free-solid-svg-icons';

@Component({
    selector: 'app-dashboard',
    templateUrl: './dashboard.component.html',
    styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
    iconUp = faArrowUp;
    iconLeft = faArrowLeft;
    iconRight = faArrowRight;
    iconDown = faArrowDown;

    constructor() {
    }

    ngOnInit() {
    }

}
