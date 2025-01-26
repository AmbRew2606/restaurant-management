import React from "react";
import { BrowserRouter as Router, Route, Routes, Navigate } from "react-router-dom";
import Login from "./components/Login";
import AdminMenu from "./components/AdminMenu";
import AdminSettings from "./components/AdminSettings";
import Menu from "./components/Menu";  // Импорт компонента Menu

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<h1>Welcome to the Restaurant App</h1>} />
        <Route path="/menu" element={<Menu />} />
      </Routes>
    </Router>
  );


  //для админки
  // const isAuthenticated = localStorage.getItem("authenticated") === "true";

  // return (
  //   <Router>
  //     <Routes>
  //       <Route path="/admin/" element={<Login />} />
  //       <Route
  //         path="/admin/menu"
  //         element={isAuthenticated ? <AdminMenu /> : <Navigate to="/admin/" />}
  //       />
  //       <Route
  //         path="/admin/settings"
  //         element={isAuthenticated ? <AdminSettings /> : <Navigate to="/admin/" />}
  //       />
  //     </Routes>
  //   </Router>
  // );
}

export default App;


// import logo from './logo.svg';
// import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }

// export default App;
