import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import {Observable} from "rxjs";
import { environment } from 'src/environments/environment';
import { Get100VisitorsResponse, GetOneVisitorResponse } from '../models/visitor';

@Injectable({
  providedIn: 'root'
})
export class VisitorService {
  private readonly apiEndpoint = environment.apiUrl.visitor;
  constructor(private httpClient: HttpClient) { }

  GetCurrentVisitor(): Observable<GetOneVisitorResponse> {
    console.log(this.apiEndpoint+`/current`);
    
    return this.httpClient.get<GetOneVisitorResponse>(this.apiEndpoint+`/current`);
  }

  GetLast100Visitors(): Observable<Get100VisitorsResponse> {
    console.log(this.apiEndpoint+`/last`);

    return this.httpClient.get<Get100VisitorsResponse>(this.apiEndpoint+`/last`);
  }

  GetTop100Visitors(): Observable<Get100VisitorsResponse> {
    console.log(this.apiEndpoint + `/top`);
    
    return this.httpClient.get<Get100VisitorsResponse>(this.apiEndpoint+`/top`);
  }
}
