-- 3
SELECT DISTINCT ON (e.id) e.*
FROM "Employees" AS e,
	 "MedicalCheckups" AS m,
	 "Positions" AS p
WHERE p.position_name = 'Водитель локомотива'
  AND p.position_id = e.position_id
  AND e.id = m.employee_id
  AND EXTRACT(YEAR FROM m.medical_checkup_date) = 2022
  AND e.sex = 'M'
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50
  AND e.salary >= 0
  AND e.salary < 50000;


SELECT COUNT(*) AS "Число водитилей"
FROM (
	SELECT DISTINCT ON (e.id) e.*
	FROM "Employees" AS e,
		 "MedicalCheckups" AS m,
		 "Positions" AS p
	WHERE p.position_name = 'Водитель локомотива'
	  AND p.position_id = e.position_id
	  AND e.id = m.employee_id
	  AND EXTRACT(YEAR FROM m.medical_checkup_date) = 2022
	  AND e.sex = 'M'
	  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
	  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50
	  AND e.salary >= 0
	  AND e.salary < 50000
) AS sub;


