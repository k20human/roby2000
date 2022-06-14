import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
    providedIn: 'root'
})
export class MercureService {

    protected sendData(topicName: string, data: any) {
        const body = new URLSearchParams({
            data: JSON.stringify(data),
            topic: topicName
        });

        fetch(environment.hub.url, {method: 'POST', body, headers: {
            Authorization: 'Bearer ' + environment.hub.credentials
        }});
    }
}
