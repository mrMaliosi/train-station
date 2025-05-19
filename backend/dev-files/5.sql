-- 5
SELECT DISTINCT ON (l.id) l.*
FROM "Locomotives" AS l,
	 "Repairs" AS r
WHERE r.repair_end_date >= '2022-01-01'
  AND r.repair_end_date < '2025-01-01'
  AND l.id = r.locomotive_id
  AND r.repair_type = 'плановый';


SELECT DISTINCT ON (l.id) l.*
FROM "Locomotives" AS l,
	 "Repairs" AS r
WHERE r.repair_start_date = '2022-01-01'
  AND l.id = r.locomotive_id;


SELECT *
FROM (
	SELECT l.id, l.model, l.status, COUNT(l.id) AS repair_count
	FROM "Locomotives" AS l,
		 "Repairs" AS r
	WHERE l.id = r.locomotive_id
	GROUP BY l.id, l.model, l.status
) AS sub
WHERE sub.repair_count = 2
ORDER BY repair_count DESC;