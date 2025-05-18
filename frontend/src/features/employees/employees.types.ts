// employees.types.ts
export type Sex = 'M' | 'F' | string;

export interface Employee {
  ID: number;
  Name: string;
  Surname: string;
  Patronymic?: string;
  Sex: string;
  PositionName: string;
  DepartmentName: string;
  Salary: number;
  BirthDate: { Time: string | number };
  HiredAt?: { Time: string | number };
  Experience?: number;
  ChildNumber?: number;
}

// Тип для формы фильтров (строки для input)
export interface EmployeeFilterForm {
  departmentID: string;
  sex: string;
  ageFrom: string;
  ageTo: string;
  experienceFrom: string;
  experienceTo: string;
  childrenFrom: string;
  childrenTo: string;
  minSalary: string;
  maxSalary: string;
}

export interface EmployeeCreate {
  name: string;
  surname: string;
  patronymic?: string;
  sex: 'M' | 'F';
  position_id: number;
  salary: number;
  birth_date: string; // 'YYYY-MM-DD'
  hired_at: string;   // 'YYYY-MM-DD'
  child_number: number;
}