-- 8
SELECT *
FROM "Routes" AS r,
	 "DelayReason" AS c
WHERE r.route_id = c.route_id
  AND r.status = 'задерживается'
  AND r.route_id = 16
  AND c.reason = 'Восстание коммунистов';

SELECT COUNT(*) AS delayed_num
FROM "Routes" AS r,
	 "DelayReason" AS c
WHERE r.route_id = c.route_id
  AND r.status = 'задерживается'
  AND r.route_id = 16
  AND c.reason = 'Восстание коммунистов';


SELECT r.route_id, r.train_number, COUNT(r.route_id) AS return_num
FROM "Routes" AS r,
	 "Tickets" AS ti,
	 "DelayReason" AS c
WHERE r.route_id = ti.route_id
  AND r.route_id = 16
  AND c.reason = 'Восстание коммунистов'
  AND ti.ticket_status = 'возвращён'
GROUP BY r.route_id, r.train_number;