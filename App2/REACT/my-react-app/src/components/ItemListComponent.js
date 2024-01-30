// ItemListComponent.js

import React from 'react';

const ItemListComponent = ({ items }) => {
  return (
    <div>
      <h2>Item List</h2>
      <ul>
        {items.map(item => (
          <li key={item.id}>
            {item.name} - Price: {item.price}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ItemListComponent;
