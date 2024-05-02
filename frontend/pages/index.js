// Archivo `pages/index.js`
import { useEffect, useState } from 'react';

export default function Home() {
    const [users, setUsers] = useState([]);
    const [products, setProducts] = useState([]);

    useEffect(() => {
        async function fetchData() {
            // Realiza llamadas a los microservicios
            const usersResponse = await fetch('/api/usuarios');
            const usersData = await usersResponse.json();
            setUsers(usersData);

            const productsResponse = await fetch('/api/products');
            const productsData = await productsResponse.json();
            setProducts(productsData);
        }

        fetchData();
    }, []);

    return (
        <div>
            <h1>Lista de Usuarios</h1>
            <ul>
                {users.map(user => (
                    <li key={user.id}>{user.nombre} - {user.email}</li>
                ))}
            </ul>

            <h1>Lista de Productos</h1>
            <ul>
                {products.map(product => (
                    <li key={product.id}>{product.name} - ${product.price}</li>
                ))}
            </ul>
        </div>
    );
}
