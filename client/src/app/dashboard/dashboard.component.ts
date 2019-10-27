import { Component, OnInit } from '@angular/core';
import { faArrowLeft, faArrowRight, faArrowUp, faArrowDown } from '@fortawesome/free-solid-svg-icons';
import { MovementService } from '../services/movement.service';

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

    constructor(
        private movementService: MovementService
    ) {
    }

    ngOnInit() {
    }

    moveUp() {
        this.movementService.goUp();
    }

    moveDown() {
        this.movementService.goDown();
    }

    moveLeft() {
        this.movementService.goLeft();
    }

    moveRight() {
        this.movementService.goRight();
    }
}
