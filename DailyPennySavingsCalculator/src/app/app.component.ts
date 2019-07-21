import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'DailyPennySavingsCalculator';

  public minDate: Date = new Date ("01/01/2019");
  public maxDate: Date = new Date ("01/01/2025");
  public startDateValue: Date = new Date ();
  public endDateValue: Date = new Date ();

  date: Date ;

  onGoClicked(): void {
    console.log(this.startDateValue);
  }

  onStartChange(args) {
    this.startDateValue = args.value;
  }
  onEndChange(args) {
    this.endDateValue = args.value;
  }
}
