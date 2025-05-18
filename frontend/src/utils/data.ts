export function getSexLabel(sex: string): string {
  if (sex === 'M') return 'Мужской';
  if (sex === 'F') return 'Женский';
  return '-';
}
