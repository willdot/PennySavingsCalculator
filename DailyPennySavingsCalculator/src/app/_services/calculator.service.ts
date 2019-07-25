import { Injectable, OnInit } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import { IRequest } from '../_models/request';
import { environment } from '../../environments/environment.prod';

@Injectable()
export class CalculatorService {
    private url  = environment.apiUrl;

    constructor(private _http: HttpClient) {
    }

    calculate(request: IRequest): Observable<string> {

        const body = JSON.stringify(request);

        return this._http.post<string>(this.url, body)
        .pipe(
            catchError(this.handleError)
        );

    }

    private handleError(error: HttpErrorResponse) {
        if (error.error instanceof ErrorEvent) {
          // A client-side or network error occurred. Handle it accordingly.
          console.error('An error occurred:', error.error.message);
        } else {

          return throwError(
            error);
        }

        // return an observable with a user-facing error message
        return throwError(
          'Something bad happened; please try again later.');
      }
}

