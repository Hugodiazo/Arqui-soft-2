// src/components/Register.js
import React, { useState } from 'react';
import { USERS_API_URL } from '../config';
import './Register.css';

function Register() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('user'); // Valor por defecto para el rol
  const [message, setMessage] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${USERS_API_URL}/users`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, email, password, role }),
      });

      if (!response.ok) throw new Error(`Error: ${response.statusText}`);

      setMessage('Usuario registrado con éxito');
      setName('');
      setEmail('');
      setPassword('');
      setRole('user');
    } catch (error) {
      console.error('Error al registrar usuario:', error);
      setMessage('Error al conectar con el servidor');
    }
  };

  return (
    <div className="register-form">
      <h2>Registrarse</h2>
      <p className="register-message">{message}</p>
      <form onSubmit={handleRegister}>
        <input
          type="text"
          placeholder="Nombre"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
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
        <select value={role} onChange={(e) => setRole(e.target.value)}>
          <option value="user">Usuario</option>
          <option value="admin">Administrador</option>
        </select>
        <button type="submit">Registrarse</button>
      </form>
    </div>
  );
}

export default Register;