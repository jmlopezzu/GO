import React, { useEffect, useState } from 'react';

const App = () => {
  const [items, setItems] = useState([]);

  // Establish WebSocket connection
  const socket = new WebSocket('ws://localhost:3000/ws');

  useEffect(() => {
    // Fetch items from the server when the component mounts
    const fetchItems = async () => {
      try {
        const response = await fetch('http://localhost:3000/api/items');
        const data = await response.json();
        setItems(data);
      } catch (error) {
        console.error('Error fetching items:', error);
      }
    };

    // Connection opened
    socket.addEventListener('open', (event) => {
      console.log('WebSocket connection opened:', event);
    });

    // Listen for messages
    socket.addEventListener('message', (event) => {
      console.log('Received message from server:', event.data);
    });

    // Connection closed
    socket.addEventListener('close', (event) => {
      console.log('WebSocket connection closed:', event);
    });

    // Fetch items from the server when the component mounts
    fetchItems();

    // Clean up on component unmount
    return () => {
      // Close the WebSocket connection when the component unmounts
      socket.close();
    };
  }, []); // Empty dependency array ensures the effect runs only once on mount

  return (
    <div>
      <h1>Azen WebSocket</h1>
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            {item.name} - ${item.price}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;
