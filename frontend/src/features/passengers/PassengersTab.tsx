import React, { useEffect, useState } from 'react';
import type { PassengerWithInfo, PassengerFilter } from './passengers.types';
import { getFilteredPassengers, deletePassenger } from './passengers.api';
import PassengersTable from './PassengersTable';

export default function PassengersTab() {
  const [passengers, setPassengers] = useState<PassengerWithInfo[]>([]);
  const [filters, setFilters] = useState<PassengerFilter>({});

  useEffect(() => {
    fetchPassengers();
  }, [filters]);

  const fetchPassengers = () => {
    getFilteredPassengers(filters).then(setPassengers).catch(console.error);
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm('Удалить пассажира?')) return;
    await deletePassenger(id);
    fetchPassengers();
  };

  return (
    <>      
      <PassengersTable
        passengers={passengers}
        onFilterChange={setFilters}
        onDelete={handleDelete}
      />
    </>
  );
}
