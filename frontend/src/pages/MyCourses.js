import React, { useEffect, useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import AuthContext from "../context/AuthContext";
import './MyCourses.css';

const MyCourses = () => {
  const [courses, setCourses] = useState([]);
  const { userId, token } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (!userId || !token) {
      navigate('/login');
      return;
    }

    const fetchMyCourses = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/enrollments/${userId}`, {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('Error al obtener tus cursos');
        }

        const data = await response.json();
        setCourses(data || []);
      } catch (error) {
        console.error('Error al obtener tus cursos:', error);
        alert('Hubo un problema al obtener tus cursos');
      }
    };

    fetchMyCourses();
  }, [userId, token, navigate]);

  return (
    <div className="my-courses">
      <h2>Mis Cursos</h2>
      {courses.length > 0 ? (
        courses.map((course) => (
          <div key={course.id} className="course-item">
            <h3>{course.title}</h3>
            <p>{course.description}</p>
          </div>
        ))
      ) : (
        <p>No estás inscrito en ningún curso.</p>
      )}
    </div>
  );
};

export default MyCourses;