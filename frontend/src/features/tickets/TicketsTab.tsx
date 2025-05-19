import React, { useEffect, useState } from 'react';
import type { Ticket, TicketFilter } from './tickets.types';
import { getFilteredTickets, deleteTicket } from './tickets.api';
import TicketsTable from './TicketsTable';

export default function TicketsTab() {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [filters, setFilters] = useState<TicketFilter>({});

  useEffect(() => {
    load();
  }, [filters]);

  const load = () => {
    getFilteredTickets(filters).then(setTickets).catch(console.error);
  };

  const handleDelete = async (id: number) => {
    if (!window.confirm('Удалить билет?')) return;
    await deleteTicket(id);
    load();
  };

  return (
    <>
      <TicketsTable
        tickets={tickets}
        onFilterChange={setFilters}
        onDelete={handleDelete}
      />
    </>
  );
}