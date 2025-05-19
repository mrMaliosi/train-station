export interface PassengerWithInfo {
  passenger_id: number;
  name: string;
  surname: string;
  patronimic?: string;
  sex: string;
  birth_date: string;
  age: number;
  route_id: number;
  travel_date: string;
  has_luggage: boolean;
}

export interface PassengerFilter {
  routeID?: number;
  sex?: string;
  minAge?: number;
  maxAge?: number;
  hasLuggage?: boolean;
  abroad?: boolean;
  travelDate?: string;
}
