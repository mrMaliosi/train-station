import type { DepartmentInfo, DepartmentInfoAPI } from './departments.types';
import type { AxiosResponse } from 'axios';
import api from '../../api';

/** Получить информацию по департаментам */
export async function getDepartmentsInfo(): Promise<DepartmentInfo[]> {
  const response: AxiosResponse<any> = await api.get('/departments/info');
  const raw: DepartmentInfoAPI[] = Array.isArray(response.data) ? response.data : response.data.items;
  return raw.map((r): DepartmentInfo => ({
    name:               r.department_name,
    directorName:       r.director_name || undefined,
    directorSurname:    r.director_surname || undefined,
    directorPatronymic: r.director_patronymic || undefined,
    birthDate:          r.birth_date?.Valid ? r.birth_date.Time : undefined,
  }));
}

/** Удаление департамента */
export async function deleteDepartment(id: number): Promise<void> {
  await api.delete(`/departments/${id}`);
}