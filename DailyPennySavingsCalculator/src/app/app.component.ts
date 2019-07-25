import { Component } from '@angular/core';
import { CalculatorService } from './_services/calculator.service';
import { IRequest } from './_models/request';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(private calculatorService: CalculatorService) {

  }

  title = 'DailyPennySavingsCalculator';

  total = '';
  errorMessage = '';

  public minDate: Date = new Date ('01/01/2019');
  public maxDate: Date = new Date ('01/01/2025');
  public startDateValue: Date = new Date ();
  public endDateValue: Date = new Date ();

  date: Date ;

  onGoClicked(): void {

    const req: IRequest = {
      start: this.startDateValue,
      end: this.endDateValue
    };


    this.calculatorService.calculate(req).subscribe(response => {
        this.total = response;
        this.errorMessage = '';
      },
      error => {
        const message: HttpErrorResponse = <any>error;

        this.errorMessage = message.error;
        this.total = '';
      }
    );

  }

  onStartChange(args) {
    this.startDateValue = args.value;
  }
  onEndChange(args) {
    this.endDateValue = args.value;
  }
}
