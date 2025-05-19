export interface DepartmentInfoAPI {
  department_name: string;
  director_name?: string;
  director_surname?: string;
  director_patronymic?: string;
  birth_date?: { Time: string; Valid: boolean };
}

export interface DepartmentInfo {
  name: string;
  directorName?: string;
  directorSurname?: string;
  directorPatronymic?: string;
  birthDate?: string;
}