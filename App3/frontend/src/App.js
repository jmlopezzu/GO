// src/App.js
import React, { useState, useEffect } from 'react';

const App = () => {
  const [items, setItems] = useState([]);
  const [newItem, setNewItem] = useState({ name: '', price: 0 });

  useEffect(() => {
    fetch('/api/items')
      .then((response) => response.json())
      .then((data) => setItems(data));
  }, []);

  const handleAddItem = () => {
    fetch('/api/items', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newItem),
    })
      .then((response) => response.json())
      .then((data) => {
        setItems([...items, data]);
        setNewItem({ name: '', price: 0 });
      });
  };

  return (
    <div>
      <h1>Simple Item Manager</h1>
      <ul>
        {items.map((item) => (
          <li key={item.ID}>
            {item.Name} - ${item.Price}
          </li>
        ))}
      </ul>
      <div>
        <input
          type="text"
          placeholder="Name"
          value={newItem.name}
          onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
        />
        <input
          type="number"
          placeholder="Price"
          value={newItem.price}
          onChange={(e) => setNewItem({ ...newItem, price: parseFloat(e.target.value) })}
        />
        <button onClick={handleAddItem}>Add Item</button>
      </div>
    </div>
  );
};

export default App;
