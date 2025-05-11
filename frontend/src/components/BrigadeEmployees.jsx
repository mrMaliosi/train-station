import { useState, useEffect } from "react";
import EmployeesTable from "./EmployeesTable";

export default function EmployeesFilter() {
  const [filters, setFilters] = useState({
    gender: "",
    minExperience: "",
    maxExperience: "",
    position: ""
  });

  const [sortBy, setSortBy] = useState("");
  const [order, setOrder] = useState("asc");
  const [employees, setEmployees] = useState([]);

  const fetchEmployees = async () => {
    const query = {
      ...filters,
      sortBy,
      order
    };

    const params = new URLSearchParams(
      Object.fromEntries(
        Object.entries(query).filter(([_, v]) => v !== "")
      )
    );

    const res = await fetch(`http://localhost:8080/brigades/employees?${params}`);
    const data = await res.json();
    setEmployees(data);
  };

  useEffect(() => {
    fetchEmployees();
  }, [filters, sortBy, order]);

  const handleChange = (e) => {
    setFilters({ ...filters, [e.target.name]: e.target.value });
  };

  return (
    <div>
      <h2>Рабочие</h2>
      <EmployeesTable employees={employees} />
    </div>
  );
}
