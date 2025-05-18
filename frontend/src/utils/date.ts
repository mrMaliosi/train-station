/** Возвращает возраст (полных лет) по дате рождения */
export function calculateAge(birthDate: string | number | Date): number {
  const birth = new Date(birthDate);
  const now = new Date();
  let age = now.getFullYear() - birth.getFullYear();
  const m = now.getMonth() - birth.getMonth();
  if (m < 0 || (m === 0 && now.getDate() < birth.getDate())) {
    age--;
  }
  return age;
}

/** Возвращает стаж (полных лет) по дате приёма на работу */
export function getExperienceYears(hireDate: string | number | Date): number {
  const hiredAt = new Date(hireDate);
  const now = new Date();
  let years = now.getFullYear() - hiredAt.getFullYear();
  const m = now.getMonth() - hiredAt.getMonth();
  if (m < 0 || (m === 0 && now.getDate() < hiredAt.getDate())) {
    years--;
  }
  return years;
}
