-- 6
SELECT t.*, r.end_time - r.start_time AS route_time
FROM "Trains" AS t,
	 "Routes" AS r
WHERE r.train_number = t.train_number
ORDER BY route_time;

SELECT t.*, ti.price
FROM "Trains" AS t,
	 "Routes" AS r,
	 "Tickets" AS ti
WHERE r.train_number = t.train_number
  AND ti.route_id = r.route_id 
ORDER BY ti.price;

SELECT t.*, ti.price, r.end_time - r.start_time AS route_time
FROM "Trains" AS t,
	 "Routes" AS r,
	 "Tickets" AS ti
WHERE r.train_number = t.train_number
  AND ti.route_id = r.route_id 
ORDER BY ti.price, route_time;