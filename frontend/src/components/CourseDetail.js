// src/pages/CourseDetail.js
import React, { useEffect, useState, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import AuthContext from "../context/AuthContext";
import './CourseDetail.css';

const CourseDetail = () => {
  const { id } = useParams(); // ID del curso desde la URL
  const [course, setCourse] = useState(null);
  const { token, isAuthenticated } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/courses/${id}`, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Error al obtener detalles del curso');
        }

        const data = await response.json();
        setCourse(data);
      } catch (error) {
        console.error('Error al obtener detalles del curso:', error);
        alert('Hubo un problema al cargar el curso');
      }
    };

    fetchCourse();
  }, [id, token]);

  const handleEnrollment = async () => {
    if (!isAuthenticated) {
      alert('Debes iniciar sesión para inscribirte en un curso');
      navigate('/login');
      return;
    }

    try {
      const response = await fetch(`${API_BASE_URL}/courses/enroll`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ course_id: id }),
      });

      if (response.ok) {
        alert('¡Inscripción exitosa!');
        navigate('/my-courses');
      } else {
        alert('Error al inscribirse en el curso');
      }
    } catch (error) {
      console.error('Error al inscribirse:', error);
      alert('Hubo un problema al inscribirse en el curso');
    }
  };

  if (!course) return <p>Cargando detalles del curso...</p>;

  return (
    <div className="course-detail">
      <h2>{course.title}</h2>
      <p>{course.description}</p>
      <p>Instructor: {course.instructor}</p>
      <p>Duración: {course.duration} horas</p>
      <p>Nivel: {course.level}</p>
      <button onClick={handleEnrollment}>Inscribirme</button>
    </div>
  );
};

export default CourseDetail;