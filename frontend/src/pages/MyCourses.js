import React, { useEffect, useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import AuthContext from "../context/AuthContext";
import './MyCourses.css';

const MyCourses = () => {
  const [courses, setCourses] = useState([]);
  const { isAuthenticated, userId, token } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated || !userId || !token) {
      navigate('/login');
      return;
    }

    const fetchMyCourses = async () => {
      try {
        // Paso 1: Obtener las inscripciones del usuario
        const response = await fetch(`${API_BASE_URL}/enrollments/${userId}`, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Error al obtener tus cursos');
        }

        const enrollments = await response.json();

        // Paso 2: Para cada `course_id`, obtener la información completa del curso
        const courseDetailsPromises = enrollments.map(async (enrollment) => {
          const courseResponse = await fetch(`${API_BASE_URL}/courses/${enrollment.course_id}`, {
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${token}`,
            },
          });
          return courseResponse.ok ? courseResponse.json() : null;
        });

        const courseDetails = await Promise.all(courseDetailsPromises);

        // Filtrar cualquier valor `null` en caso de que alguna solicitud falle
        setCourses(courseDetails.filter(course => course !== null));
      } catch (error) {
        console.error('Error al obtener tus cursos:', error);
        alert('Hubo un problema al obtener tus cursos');
      }
    };

    fetchMyCourses();
  }, [isAuthenticated, userId, token, navigate]);

  return (
    <div className="my-courses">
      <h2>Mis Cursos</h2>
      {courses.length > 0 ? (
        courses.map((course) => (
          <div key={course.ID} className="course-item">
            <h3>{course.Title || "Sin título"}</h3>
            <p>{course.Description || "Sin descripción"}</p>
          </div>
        ))
      ) : (
        <p>No estás inscrito en ningún curso.</p>
      )}
    </div>
  );
};

export default MyCourses;