import type { Ticket, TicketFilter, TicketAPI } from './tickets.types';
import type { AxiosResponse } from 'axios';
import api from '../../api';

/** Получить отфильтрованные билеты */
export async function getFilteredTickets(
  filters: TicketFilter
): Promise<Ticket[]> {
  const params: Record<string, any> = {};
  if (filters.routeID !== undefined) params.routeID = filters.routeID;
  if (filters.fromDate)             params.fromDate = filters.fromDate;
  if (filters.toDate)               params.toDate = filters.toDate;
  if (filters.status)               params.status = filters.status;

  const response: AxiosResponse<any> = await api.get('/tickets', { params });
  const raw: TicketAPI[] = Array.isArray(response.data) ? response.data : response.data.items;
  return raw.map((r): Ticket => ({
    ticketId:    r.TicketID,
    routeId:     r.RouteID,
    status:      r.TicketStatus,
    passengerId: r.PassengerID,
    boughtAt:    r.BoughtAt?.Valid ? r.BoughtAt.Time : undefined,
    price:       r.Price,
    trainNumber: r.TrainNumber,
  }));
}

/** Удаление билета */
export async function deleteTicket(id: number): Promise<void> {
  await api.delete(`/tickets/${id}`);
}