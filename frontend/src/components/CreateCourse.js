// src/components/CreateCourse.js
import React, { useState, useContext } from 'react';
import { API_BASE_URL } from '../config';
import AuthContext from "../context/AuthContext";
import './CreateCourse.css';

const CreateCourse = () => {
  const { token } = useContext(AuthContext); // Obtener el token
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [duration, setDuration] = useState('');
  const [level, setLevel] = useState('');
  const [availability, setAvailability] = useState(true);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(`${API_BASE_URL}/courses`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ title, description, duration: parseInt(duration, 10), level, availability }),
      });

      if (response.ok) {
        alert('Curso creado con éxito');
        setTitle('');
        setDescription('');
        setDuration('');
        setLevel('');
        setAvailability(true);
      } else {
        alert('Error al crear el curso');
      }
    } catch (error) {
      console.error('Error al crear el curso:', error);
    }
  };

  return (
    <div className="create-course">
      <h2>Crear Curso</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Título del curso" value={title} onChange={(e) => setTitle(e.target.value)} />
        <textarea placeholder="Descripción del curso" value={description} onChange={(e) => setDescription(e.target.value)}></textarea>
        <input type="number" placeholder="Duración (en horas)" value={duration} onChange={(e) => setDuration(e.target.value)} />
        <select value={level} onChange={(e) => setLevel(e.target.value)}>
          <option value="">Seleccionar nivel</option>
          <option value="beginner">Principiante</option>
          <option value="intermediate">Intermedio</option>
          <option value="advanced">Avanzado</option>
        </select>
        <label>
          <input type="checkbox" checked={availability} onChange={(e) => setAvailability(e.target.checked)} />
          Disponible
        </label>
        <button type="submit">Crear Curso</button>
      </form>
    </div>
  );
};

export default CreateCourse;