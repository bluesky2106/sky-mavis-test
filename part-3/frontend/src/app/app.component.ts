import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { map } from 'rxjs/operators';
import { Get100VisitorsResponse, GetOneVisitorResponse, Visitor } from './core/models/visitor';
import { VisitorService } from './core/services/visitor.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'skymavis';

  public currentVisitor: Visitor;
  public top100Visitors: Visitor[];
  public last100Visitors: Visitor[];

  headElements = ['#', 'IP Address', 'Location', 'Timezone', 'Last Visit', 'Number Of Visits'];
  constructor(
    private visitorService: VisitorService
  ) {
    this.currentVisitor = {
      ip_address: "",
      location: "",
      timezone: "",
      last_visit: "",
      visits: 0,
    };
    this.top100Visitors = [];
    this.last100Visitors = [];
  }

  ngOnInit(): void {
    this.visitorService.GetCurrentVisitor().
      subscribe((res: GetOneVisitorResponse) => {
        if (!!res.data) {
          this.currentVisitor = res.data;
        } else {
          alert("cannot get current visitor info")
        }
      },
      (error: HttpErrorResponse) => {
        alert(error.message)
      });
    
    this.visitorService.GetTop100Visitors().
      subscribe((res: Get100VisitorsResponse) => {
        if (!!res.data) {
          this.top100Visitors = res.data;
        } else {
          alert("cannot get top 100 visitors")
        }
      },
      (error: HttpErrorResponse) => {
        alert(error.message)
        });
    
      this.visitorService.GetLast100Visitors().
        subscribe((res: Get100VisitorsResponse) => {
          if (!!res.data) {
            this.last100Visitors = res.data;
          } else {
            alert("cannot get last 100 visitors")
          }
        },
        (error: HttpErrorResponse) => {
          alert(error.message)
        });
  }
}
