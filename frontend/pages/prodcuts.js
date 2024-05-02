// pages/products.js
import React, { useEffect, useState } from 'react';
import Table from '../components/Table';

const Products = () => {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        const fetchProducts = async () => {
            const response = await fetch('/api/products'); // Ajusta la URL según sea necesario
            const data = await response.json();
            setProducts(data);
        };
        fetchProducts();
    }, []);

    const columns = [
        { field: 'id', headerName: 'ID' },
        { field: 'name', headerName: 'Name' },
        { field: 'price', headerName: 'Price' },
        // Añade más columnas según sea necesario
    ];

    return (
        <div>
            <h1>Productos</h1>
            <Table data={products} columns={columns} />
        </div>
    );
};

export default Products;
