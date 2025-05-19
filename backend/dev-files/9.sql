-- 9
SELECT *
FROM "Tickets" AS ti
WHERE ti.bought_at >= '2022-01-01'
  AND ti.bought_at < '2025-01-01'
  AND ti.ticket_status = 'куплен';

SELECT COUNT(*)
FROM "Tickets" AS ti
WHERE ti.bought_at >= '2022-01-01'
  AND ti.bought_at < '2025-01-01'
  AND ti.ticket_status = 'куплен';

SELECT *
FROM "Tickets" AS ti,
	 "Routes" AS r
WHERE ti.bought_at >= '2022-01-01'
  AND ti.bought_at < '2025-01-01'
  AND r.route_id = ti.route_id
  AND r.route_id = 1
  AND ti.ticket_status = 'куплен';


SELECT r.*, ti.price, r.end_time - r.start_time AS route_time
FROM "Tickets" AS ti,
	 "Routes" AS r
WHERE ti.bought_at >= '2022-01-01'
  AND ti.bought_at < '2025-01-01'
  AND r.route_id = ti.route_id
  AND r.route_id = 1
  AND ti.ticket_status = 'куплен'
ORDER BY ti.price, route_time;
