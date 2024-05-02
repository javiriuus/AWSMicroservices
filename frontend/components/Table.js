// components/Table.js
import React from 'react';

const Table = ({ data, columns }) => {
    return (
        <table>
            <thead>
                <tr>
                    {columns.map((col) => (
                        <th key={col.field}>{col.headerName}</th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {data.map((row, index) => (
                    <tr key={index}>
                        {columns.map((col) => (
                            <td key={col.field}>{row[col.field]}</td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    );
};

export default Table;
