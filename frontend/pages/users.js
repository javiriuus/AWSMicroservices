// pages/users.js
import React, { useEffect, useState } from 'react';
import Table from '../components/Table';

const Users = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch('/api/users'); // Ajusta la URL según sea necesario
            const data = await response.json();
            setUsers(data);
        };
        fetchUsers();
    }, []);

    const columns = [
        { field: 'id', headerName: 'ID' },
        { field: 'name', headerName: 'Name' },
        { field: 'email', headerName: 'Email' },
        // Añade más columnas según sea necesario
    ];

    return (
        <div>
            <h1>Usuarios</h1>
            <Table data={users} columns={columns} />
        </div>
    );
};

export default Users;
