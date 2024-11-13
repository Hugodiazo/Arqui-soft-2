// src/components/Navbar.js
import React, { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import AuthContext from "../context/AuthContext";
import './Navbar.css';

const Navbar = () => {
  const { isAuthenticated, logout, userRole } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <nav className="navbar">
      <ul className="navbar-list">
        <li className="navbar-item">
          <Link to="/" className="navbar-link">Home</Link>
        </li>
        <li className="navbar-item">
          <Link to="/courses" className="navbar-link">Cursos</Link>
        </li>
        {isAuthenticated && (
          <>
            <li className="navbar-item">
              <Link to="/my-courses" className="navbar-link">Mis Cursos</Link>
            </li>
            {userRole === 'admin' && ( // Verifica si el usuario es administrador
              <li className="navbar-item">
                <Link to="/create-course" className="navbar-link">Crear Curso</Link>
              </li>
            )}
            <li className="navbar-item">
              <button className="logout-button" onClick={handleLogout}>
                Cerrar Sesión
              </button>
            </li>
          </>
        )}
        {!isAuthenticated && (
          <>
            <li className="navbar-item">
              <Link to="/login" className="navbar-link">Iniciar Sesión</Link>
            </li>
            <li className="navbar-item">
              <Link to="/register" className="navbar-link">Registro</Link>
            </li>
          </>
        )}
      </ul>
    </nav>
  );
};

export default Navbar;