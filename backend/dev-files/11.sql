-- 11
SELECT *
FROM "Tickets" AS ti,
	 "Passengers" AS p,
	 "Luggage" AS lu,
	 "TicketsLuggage" AS tl
WHERE ti.ticket_id = tl.ticket_id
  AND tl.luggage_id = lu.luggage_id
  AND ti.passenger_id = p.passenger_id;


SELECT *
FROM "Tickets" AS ti,
	 "Passengers" AS p,
	 "Luggage" AS lu,
	 "TicketsLuggage" AS tl
WHERE ti.ticket_id = tl.ticket_id
  AND tl.luggage_id = lu.luggage_id
  AND ti.passenger_id = p.passenger_id
  AND p.sex = 'M'
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50;

SELECT DISTINCT ON (p.passenger_id) p.*
FROM "Tickets" AS ti,
	 "Passengers" AS p,
	 "Routes" AS r,
	 "RoutesStations" AS rs,
	 "ArrivalTime" AS art,
	 "Stations" AS s
WHERE ti.passenger_id = p.passenger_id
  AND r.route_id = ti.route_id
  AND rs.route_id = r.route_id
  AND s.station_id = rs.station_id
  AND rs.arrival_time_id = art.arrival_id
  AND DATE(art.real_arrival_time) = '2025-04-08'
  AND s.is_abroad = true;