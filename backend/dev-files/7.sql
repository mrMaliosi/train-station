-- 7
SELECT *
FROM "Routes" AS r,
	 "Trains" AS t
WHERE r.train_number = t.train_number
  AND r.route_id = 2
  AND r.status = 'отменён';

SELECT COUNT(*)
FROM (
	SELECT *
	FROM "Routes" AS r,
		 "Trains" AS t
	WHERE r.train_number = t.train_number
	  AND r.route_id = 2
	  AND r.status = 'отменён'
) AS sub;