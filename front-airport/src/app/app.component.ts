import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import Swal from "sweetalert2";
import * as moment from 'moment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'front-airport';

  debut = ""
  fin = ""
  datatypes = []
  dataTypeSelected = ""
  airportSelected = "";
  airports = []
  listMetrics: any[]
  isRangeForm = true
  isAverageForm = false
  dayDate = ""
  listAverages: any[]

  constructor(private http: HttpClient) {
    this.listMetrics = []
    this.listAverages = []
  }

  ngOnInit() {
    this.getDatatypes()
    this.getAirports()
  }


  getDatatypes() {
    this.http.get<[]>("http://localhost:8086/api/datatype/").subscribe((data) => {
      this.datatypes = data
    })
  }

  getAirports() {
    this.http.get<[]>("http://localhost:8086/api/airport/").subscribe((data) => {
      this.airports = data
    })
  }

  getAverages() {
    if (this.dayDate === "" || this.airportSelected === "") {
      Swal.fire({
        title: 'Error!',
        text: 'all input are required',
        icon: 'error',
        confirmButtonText: 'Cool'
      })
      return
    }
    let date = moment(new Date(this.dayDate))
    let dateString = date.format("DD/MM/YYYY")
    this.http.get<[]>("http://localhost:8086/api/airport/" + this.airportSelected + "/average/?date=" + dateString).subscribe((data) => {
      this.listAverages = data
    })
  }


  searchRange() {
    if (this.debut === "" || this.fin === "" || this.airportSelected === "" || this.dataTypeSelected == "") {
      Swal.fire({
        title: 'Error!',
        text: 'all input are required',
        icon: 'error',
        confirmButtonText: 'Cool'
      })
      return
    }
    let dateDebut = moment(new Date(this.debut))
    let dateFin = moment(new Date(this.fin))
    let dateDebutString = dateDebut.format("DD/MM/YYYY hh")
    console.log(dateDebutString)
    let dateFinString = dateFin.format("DD/MM/YYYY hh")

    this.http.get<[]>("http://localhost:8086/api/airport/" + this.airportSelected +
      "/datatype/" + this.dataTypeSelected + "/range/?dateDebut=" + dateDebutString + "&dateFin=" + dateFinString).subscribe((data) => {
      this.listMetrics = data
    })

  }

  setRangeForm() {
    this.isRangeForm = true
    this.isAverageForm = false
  }

  setAverageForm() {
    this.isAverageForm = true
    this.isRangeForm = false
  }
}
