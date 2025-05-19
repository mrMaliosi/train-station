-- 4
SELECT *
FROM "Locomotives" AS l,
	 "Stations" AS s
WHERE l.base_station_id = s.station_id
  AND s.station_id = 3;

SELECT DISTINCT ON (l.id) l.*, art.real_arrival_time
FROM "Locomotives" AS l,
	 "Stations" AS s,
	 "RoutesStations" AS rs,
	 "ArrivalTime" AS art,
	 "Trains" AS t,
	 "Routes" AS r
WHERE l.base_station_id = s.station_id
  AND art.arrival_id = rs.arrival_time_id
  AND s.station_id = rs.station_id
  AND l.id = t.locomotive_id
  AND r.train_number = t.train_number
  AND r.route_id = rs.route_id
  AND s.station_id = 3
  AND DATE(art.real_arrival_time) = '2025-04-04';


SELECT l.id, l.model, COUNT(l.id) AS ended_count
FROM "Locomotives" AS l,
	 "Trains" AS t,
	 "Routes" AS r
WHERE l.id = t.locomotive_id
  AND t.train_number = r.train_number
  AND r.status = 'закончил'
GROUP BY l.id, l.model;
	
SELECT *
FROM (
	SELECT l.id, l.model, COUNT(l.id) AS ended_count
	FROM "Locomotives" AS l,
		 "Trains" AS t,
		 "Routes" AS r
	WHERE l.id = t.locomotive_id
	  AND t.train_number = r.train_number
	  AND r.status = 'закончил'
	GROUP BY l.id, l.model
) AS sub
WHERE sub.ended_count = 1
ORDER BY ended_count DESC;

