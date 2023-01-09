import {Component, Input} from '@angular/core';
import {Chart} from "chart.js/auto";

@Component({
  selector: 'app-line-chart',
  templateUrl: './line-chart.component.html',
  styleUrls: ['./line-chart.component.css']
})
export class LineChartComponent {

  public chart: any

  ngOnInit() {
  }

  createChart(data: any[], dataName: string){
    this.chart?.destroy()

    let listLabels: any[] = []
    let listValues: any [] = []
    data.forEach((item) => {
      listLabels.push(item.Date)
      listValues.push(item.Value)
    })
    this.chart = new Chart("MyChart", {
      type: 'line', //this denotes tha type of chart

      data: {// values on X-Axis
        labels: listLabels,
        datasets: [
          {
            label: dataName,
            data: listValues,
            backgroundColor: 'blue'
          }
        ]
      },
      options: {
        aspectRatio:2.5
      }

    });
  }

}
