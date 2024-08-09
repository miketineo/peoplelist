import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
    const [records, setRecords] = useState([]);

    useEffect(() => {
        const apiUrl = process.env.REACT_APP_API_URL;
        axios.get(`${apiUrl || 'http://api-local.miketineo.com'}/api/records`)
            .then(response => {
                setRecords(response.data.users);
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }, []);

    return (
        <div>
            <h1>User Records</h1>
            <ul>
                {records.map((record, index) => (
                    <li key={index}>{record.name} - {record.role}</li>
                ))}
            </ul>
        </div>
    );
}

export default App;
