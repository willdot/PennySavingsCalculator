import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import { HttpResponse, HttpHeaderResponse } from '@angular/common/http/src/response';
import { IRequest } from '../_models/request';

@Injectable()
export class CalculatorService {
    private url  = 'http://localhost:8080/calculate';

    constructor(private _http: HttpClient) {
    }


    calculate(request: IRequest): Observable<string> {

        const body = JSON.stringify(request);

        console.log(body);

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
          // The backend returned an unsuccessful response code.
          // The response body may contain clues as to what went wrong,
          console.error(
            `Backend returned code ${error.status}, ` +
            `body was: ${error.error}`);
        }
        // return an observable with a user-facing error message
        return throwError(
          'Something bad happened; please try again later.');
      }
}

