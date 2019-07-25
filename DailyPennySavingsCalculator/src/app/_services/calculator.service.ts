import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import { HttpResponse, HttpHeaderResponse } from '@angular/common/http/src/response';
import { IRequest } from '../_models/request';
import { IResponse } from '../_models/response';

@Injectable()
export class CalculatorService {
    private url  = 'http://localhost:8080/calculator';

    constructor(private _http: HttpClient) {
    }


    calculate(request: IRequest): Observable<IResponse> {

        const body = JSON.stringify(request);

        console.log(body);

        return this._http.post<IResponse>(this.url, body)
        .pipe(
            catchError(this.handleError)
        );

    }

    private handleError(err: HttpErrorResponse) {
        console.log(err.error);
        // tslint:disable-next-line: deprecation
        return Observable.throw(err);
    }
}

