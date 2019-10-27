import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { MercureService } from './mercure.service';
import { TOPICS, MOVEMENT_DIRECTIONS } from '../shared/shared.const';

@Injectable({
    providedIn: 'root'
})
export class MovementService extends MercureService {
    goUp() {
        this.sendData(TOPICS.movement, {
            direction: MOVEMENT_DIRECTIONS.up,
        });
    }

    goDown() {
        this.sendData(TOPICS.movement, {
            direction: MOVEMENT_DIRECTIONS.down,
        });
    }

    goRight() {
        this.sendData(TOPICS.movement, {
            direction: MOVEMENT_DIRECTIONS.right,
        });
    }

    goLeft() {
        this.sendData(TOPICS.movement, {
            direction: MOVEMENT_DIRECTIONS.left,
        });
    }
}
