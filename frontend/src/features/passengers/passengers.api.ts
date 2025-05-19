import type { PassengerWithInfo, PassengerFilter } from './passengers.types';
import type { AxiosResponse } from 'axios';
import api from '../../api';

/** Получить отфильтрованных пассажиров */
export async function getFilteredPassengers(
  filters: PassengerFilter
): Promise<PassengerWithInfo[]> {
  const params: Record<string, any> = {};
  if (filters.routeID !== undefined)   params.routeID    = filters.routeID;
  if (filters.sex)                     params.sex        = filters.sex;
  if (filters.minAge !== undefined)    params.minAge     = filters.minAge;
  if (filters.maxAge !== undefined)    params.maxAge     = filters.maxAge;
  if (filters.hasLuggage !== undefined) params.hasLuggage = filters.hasLuggage;
  if (filters.abroad !== undefined)    params.abroad     = filters.abroad;
  if (filters.travelDate)              params.travelDate = filters.travelDate;

  const response: AxiosResponse<any> = await api.get('/passengers/filter', { params });
  const raw = Array.isArray(response.data) ? response.data : response.data.items;
  return raw.map((r: any): PassengerWithInfo => ({
    passenger_id: r.passenger_id,
    name:         r.name,
    surname:      r.surname,
    patronimic:   r.patronimic,
    sex:          r.sex,
    birth_date:   r.birth_date,
    age:          r.age,
    route_id:     r.route_id,
    travel_date:  r.travel_date,
    has_luggage:  r.has_luggage,
  }));
}
/** Удаление пассажира */
export async function deletePassenger(id: number): Promise<void> {
  await api.delete(`/passengers/${id}`);
}