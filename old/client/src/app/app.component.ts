import { Component } from '@angular/core';
import { faRobot } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'roby2000';
  icon = faRobot;
}
