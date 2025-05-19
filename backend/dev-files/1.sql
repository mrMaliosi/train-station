-- Получить перечень работников жд
SELECT *
FROM "Employees";

-- Получить число работников жд
SELECT COUNT(*) AS "Число работников" 
FROM "Employees";

-- Получить начальников отделов
SELECT d.department_id, d.department_name, e.* 
FROM "Departments" AS d,
	 "Employees" AS e
WHERE director_id = id;

-- Получить pаботников указанного отдела
/*
SELECT id, name, surname, patronymic, birth_date, child_number, hired_at, sex, salary, position_name, department_name
FROM "Employees" AS e,
	 "Positions" AS p,
	 "Departments" AS d
WHERE d.department_id = p.department_id
	  AND p.position_id = e.position_id
	  AND d.department_id = :department_number;
*/  
SELECT id, name, surname, patronymic, birth_date, child_number, hired_at, sex, salary, position_name, department_name
FROM "Employees" AS e,
	 "Positions" AS p,
	 "Departments" AS d
WHERE d.department_id = p.department_id
	  AND p.position_id = e.position_id
	  AND d.department_id = 1;

-- Получить работников по опыту на станции
/*
SELECT *
FROM "Employees"
WHERE EXTRACT(YEAR FROM age(current_date, e.hired_at)) >= :years_of_expirience_from
	  AND EXTRACT(YEAR FROM age(current_date, hired_at)) < :years_of_expirience_to
ORDER BY EXTRACT(YEAR FROM age(current_date, e.hired_at)) DESC;
*/
SELECT *
FROM "Employees"
WHERE EXTRACT(YEAR FROM age(current_date, hired_at)) >= 0
	  AND EXTRACT(YEAR FROM age(current_date, hired_at)) < 7
ORDER BY EXTRACT(YEAR FROM age(current_date, hired_at)) DESC;


-- Получить работников по половому пpизнаку
/*
SELECT *
FROM "Employees"
WHERE sex = :sex;
*/
SELECT *
FROM "Employees"
WHERE sex = 'M';

-- Получить работников по возрасту
/*
SELECT *
FROM "Employees"
WHERE EXTRACT(YEAR FROM age(current_date, birth_date)) >= :age_from
	  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < :age_to
ORDER BY EXTRACT(YEAR FROM age(current_date, hired_at)) ASC;
*/
SELECT *
FROM "Employees"
WHERE EXTRACT(YEAR FROM age(current_date, birth_date)) >= 0
	  AND EXTRACT(YEAR FROM age(current_date, birth_date)) < 50
ORDER BY EXTRACT(YEAR FROM age(current_date, hired_at)) ASC;

-- Получить работников по наличию и количества детей
/*
SELECT *
FROM "Employees"
WHERE child_number >= :child_number_from
	  AND child_number < :child_number_to
ORDER BY child_number ASC;
*/
SELECT *
FROM "Employees"
WHERE child_number >= 0
	  AND child_number < 3
ORDER BY child_number ASC;

-- Получить работников по размеру заработной платы
/*
SELECT *
FROM "Employees"
WHERE salary >= :salary_to
	  AND salary < :salary_from
ORDER BY salary ASC;
*/
SELECT *
FROM "Employees"
WHERE salary >= 0
	  AND salary < 50000
ORDER BY salary ASC;