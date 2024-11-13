// src/components/Login.js
import React, { useState, useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import AuthContext from "../context/AuthContext";
import { USERS_API_URL } from '../config';
import './Login.css';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const { login, isAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${USERS_API_URL}/users/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();
      if (response.ok && data.token) {
        console.log("Login response:", data);
        login(data.token);
        setMessage('Inicio de sesión exitoso');
      } else {
        setMessage('Credenciales incorrectas');
      }
    } catch (error) {
      console.error('Error al iniciar sesión:', error);
      setMessage('Error al conectar con el servidor');
    }
  };

  // Redirige a "Mis Cursos" cuando la autenticación es exitosa
  useEffect(() => {
    if (isAuthenticated) {
      navigate('/my-courses');
    }
  }, [isAuthenticated, navigate]);

  return (
    <div className="login-form">
      <h2>Iniciar Sesión</h2>
      <p>{message}</p>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Iniciar sesión</button>
      </form>
    </div>
  );
}

export default Login;