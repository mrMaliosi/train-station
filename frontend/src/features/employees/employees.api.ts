import type { Employee, EmployeeCreate } from './employees.types';
import type { AxiosResponse } from 'axios';
import api from '../../api';

// Новые поля фильтров для API
export interface EmployeeFilters {
  departmentID?: number;
  sex?: string;
  ageFrom?: number;
  ageTo?: number;
  experienceFrom?: number;
  experienceTo?: number;
  childrenFrom?: number;
  childrenTo?: number;
  minSalary?: number;
  maxSalary?: number;
}

// Добавление сотрудника
export async function addEmployee(employee: EmployeeCreate): Promise<Employee> {
  const response = await api.post('/employees', employee);
  return response.data;
}

// Обновление сотрудника
export async function updateEmployee(employee: Employee): Promise<void> {
  await api.put(`/employees/${employee.ID}`, employee);
}

// Удаление сотрудника
export async function deleteEmployee(id: number): Promise<void> {
  await api.delete(`/employees/${id}`);
}

export async function getFilteredEmployees(
  filters: EmployeeFilters
): Promise<Employee[]> {
  const params: Record<string, any> = {};
  if (filters.departmentID)   params.department_id    = filters.departmentID;
  if (filters.sex)            params.sex              = filters.sex;
  if (filters.ageFrom)        params.age_from         = filters.ageFrom;
  if (filters.ageTo)          params.age_to           = filters.ageTo;
  if (filters.experienceFrom) params.experience_from  = filters.experienceFrom;
  if (filters.experienceTo)   params.experience_to    = filters.experienceTo;
  if (filters.childrenFrom)   params.children_from    = filters.childrenFrom;
  if (filters.childrenTo)     params.children_to      = filters.childrenTo;
  if (filters.minSalary)      params.salary_from      = filters.minSalary;
  if (filters.maxSalary)      params.salary_to        = filters.maxSalary;

  const response: AxiosResponse<any> = await api.get('/employees', { params });
  const rawSource = response.data ?? [];
  const raw = Array.isArray(rawSource)
    ? rawSource
    : (rawSource.items ?? []);
  const cleaned = Array.isArray(raw) ? raw : [];

  return cleaned.map((r: any): Employee => ({
    ID:             r.id,
    Name:           r.name,
    Surname:        r.surname,
    Patronymic:     r.patronymic || '',
    Sex:            r.sex,
    PositionName:   r.position_name,
    DepartmentName: r.department_name,
    Salary:         r.salary,
    BirthDate:      { Time: r.birth_date?.Time },
    HiredAt:        r.hired_at?.Time ? { Time: r.hired_at.Time } : undefined,
    Experience:     r.experience,
    ChildNumber:    r.child_number,
  }));
}