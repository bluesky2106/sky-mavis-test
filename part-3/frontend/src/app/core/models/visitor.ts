// Province : tỉnh / thành phố
export interface Visitor {
  ip_address: string;
  location: string;
  timezone: string;
  last_visit: string;
  visits: number;
}

export interface GetOneVisitorResponse {
  data: Visitor;
  error: string;
}

export interface Get100VisitorsResponse {
  data: Visitor[];
  error: string;
}