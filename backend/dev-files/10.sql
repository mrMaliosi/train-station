-- 10
SELECT *
FROM "Trains" AS t,
	 "Routes" AS r,
	 "RoutesStations" AS rs,
	 "Stations" AS s
WHERE t.train_type = 'пригородный'
  --AND r.route_id = 6
  AND r.route_id = rs.station_id
  AND t.train_number = r.train_number
  AND s.station_id = rs.station_id
  AND s.station_name = 'Грасиона'
  AND rs.station_number > 1;