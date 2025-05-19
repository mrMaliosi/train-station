-- Получить перечень работниколв в бригаде
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1;

-- Получить число работников в бригаде
/*
SELECT COUNT(*) AS "Число работников в бригаде"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id;
*/

SELECT COUNT(*) AS "Число работников в бригаде"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1;

-- Получить перечень и общее число работников по всем отделам
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id;


SELECT COUNT(*) AS "Число работников по всем отделам"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id;


-- Получить перечень и общее число работников в указанном отделе
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Departments" AS d
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.department_id = d.department_id
  AND d.department_id = :department_id;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Departments" AS d
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.department_id = d.department_id
  AND d.department_id = 1;

/*
SELECT COUNT(*) AS "Число работников в отделе"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Departments" AS d
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id;
*/

SELECT COUNT(*) AS "Число работников в отделе"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Departments" AS d
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.department_id = d.department_id
  AND d.department_id = 1;


-- Получить перечень и общее число работников, обслуживающих некоторый локомотив
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Locomotives" AS l
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = l.locomotive_brigade_id
  AND l.id = :locomotive_id;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Locomotives" AS l
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = l.locomotive_brigade_id
  AND l.id = 1;

/*
SELECT COUNT(*) AS "Число работников, обслуживающих локомотив"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Locomotives" AS l
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = l.locomotive_brigade_id
  AND l.id = :locomotive_id;
*/

SELECT COUNT(*) AS "Число работников, обслуживающих локомотив"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm,
	 "Locomotives" AS l
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = l.locomotive_brigade_id
  AND l.id = 1;

-- По возрасту
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= :age_from
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < :age_to
ORDER BY EXTRACT(YEAR FROM age(current_date, hired_at)) ASC;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50
ORDER BY EXTRACT(YEAR FROM age(current_date, birth_date)) ASC;

/*
SELECT COUNT(*) AS "Число работников"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= :age_from
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < :age_to;
*/

SELECT COUNT(*) AS "Число работников"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50;

-- По средней зарплате 
/*
SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id
  AND salary >= :salary_to
  AND salary < :salary_from;
*/

SELECT *
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1
  AND e.salary >= 0
  AND e.salary < 52000;

/*
SELECT COUNT(*) AS "Число работников"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = :brigade_id
  AND salary >= :salary_to
  AND salary < :salary_from;
*/

SELECT COUNT(*) AS "Число работников"
FROM "Brigades" AS b,
	 "Employees" AS e,
	 "BrigadeMembers" AS bm
WHERE b.brigade_id = bm.brigade_id
  AND bm.employee_id = e.id
  AND b.brigade_id = 1
  AND e.salary >= 0
  AND e.salary < 52000;