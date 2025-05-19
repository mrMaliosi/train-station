-- 13
SELECT *
FROM "Tickets" AS ti,
	 "Routes" AS r
WHERE ti.ticket_status = 'возвращён'
  AND r.route_id = ti.route_id
  AND DATE(r.start_time) = '2025-04-02'
  AND r.route_id = 6;
  