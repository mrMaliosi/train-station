export interface TicketAPI {
  TicketID: number;
  RouteID: number;
  TicketStatus: string;
  PassengerID: number;
  BoughtAt: { Time: string; Valid: boolean };
  Price: number;
  TrainNumber: string;
}

export interface Ticket {
  ticketId: number;
  routeId: number;
  status: string;
  passengerId: number;
  boughtAt?: string;
  price: number;
  trainNumber: string;
}

export interface TicketFilter {
  routeID?: number;
  fromDate?: string;
  toDate?: string;
  status?: string;
}