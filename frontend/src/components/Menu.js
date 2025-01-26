import React, { useEffect, useState } from "react";

function Menu() {
  const [menuItems, setMenuItems] = useState([]);

  useEffect(() => {
    // Здесь делаем запрос к backend'у
    fetch("http://localhost:8080/menu")
      .then((response) => response.json())
      .then((data) => {
        setMenuItems(data);
      })
      .catch((error) => {
        console.error("Error fetching menu:", error);
      });
  }, []); // Пустой массив зависимостей означает, что запрос будет сделан только один раз при монтировании компонента

  return (
    <div>
      <h1>Menu</h1>
      <ul>
        {menuItems.map((item) => (
          <li key={item.id}>
            <h3>{item.name}</h3>
            <p>{item.description}</p>
            <p>{item.price} USD</p>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Menu;
