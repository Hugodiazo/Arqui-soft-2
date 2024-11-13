import React, { useEffect, useState, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import AuthContext from "../context/AuthContext";
import './CourseDetail.css';

const CourseDetail = () => {
  const { id } = useParams(); // ID del curso desde la URL
  const [course, setCourse] = useState(null);
  const { token, isAuthenticated, userId } = useContext(AuthContext);
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
        console.log("Detalles del curso recibidos:", data); // Confirmamos los datos en consola
        setCourse(data); // Asignamos los datos directamente
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
      console.log(`Intentando inscribirse en el curso con ID: ${id}`);
      const response = await fetch(`${API_BASE_URL}/courses/enroll`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ user_id: userId, course_id: id }),
      });

      if (response.ok) {
        alert('Inscripción exitosa');
        navigate('/my-courses');
      } else {
        alert('Error al inscribirse en el curso');
      }
    } catch (error) {
      console.error('Error al inscribirse:', error);
      alert('Hubo un problema al inscribirse en el curso');
    }
  };

  // Confirmación de que `course` contiene los datos correctos antes de renderizar
  if (!course) return <p>Cargando detalles del curso...</p>;

  return (
    <div className="course-detail">
      <h2>{course.Title || "Sin título"}</h2>
      <p>{course.Description || "Sin descripción"}</p>
      <p>Instructor: {course.Instructor || "Desconocido"}</p>
      <p>Duración: {course.Duration ? `${course.Duration} horas` : "Desconocida"}</p>
      <p>Nivel: {course.Level || "Desconocido"}</p>
      <button onClick={handleEnrollment} className="enroll-button">Inscribirme</button>
    </div>
  );
};

export default CourseDetail;