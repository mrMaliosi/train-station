import { Employee } from './employees.types';
import { useEffect, useState } from 'react';

interface Props {
  isOpen: boolean;
  onClose: () => void;
  employee: Employee | null;
  onSave: () => void;
}

export default function EditEmployeeModal({ isOpen, onClose, employee, onSave }: Props) {
  const [form, setForm] = useState<Employee | null>(null);

  useEffect(() => {
    setForm(employee);
  }, [employee]);

  if (!isOpen || !form) return null;

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm((prev) => prev ? { ...prev, [name]: value } : null);
  };

  const handleSubmit = async () => {
    try {
      await fetch(`/api/employees/${form.ID}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(form),
      });
      onSave();
      onClose();
    } catch (error) {
      alert('Ошибка при сохранении');
      console.error(error);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-30 flex justify-center items-center z-50">
      <div className="bg-white p-6 rounded-xl w-[500px]">
        <h2 className="text-xl font-bold mb-4">Редактировать сотрудника</h2>

        <div className="flex flex-col gap-3">
          <input name="Surname" value={form.Surname} onChange={handleChange} placeholder="Фамилия" />
          <input name="Name" value={form.Name} onChange={handleChange} placeholder="Имя" />
          <input name="Patronymic" value={form.Patronymic ?? ''} onChange={handleChange} placeholder="Отчество" />
          <select name="Sex" value={form.Sex} onChange={handleChange}>
            <option value="m">Мужской</option>
            <option value="f">Женский</option>
          </select>
          {/* Добавь остальные поля аналогично */}
        </div>

        <div className="flex justify-end gap-2 mt-4">
          <button onClick={onClose}>Отмена</button>
          <button onClick={handleSubmit} className="bg-blue-600 text-white px-4 py-1 rounded">Сохранить</button>
        </div>
      </div>
    </div>
  );
}
