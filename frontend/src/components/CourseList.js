// src/components/CourseList.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import './CourseList.css';

const CourseList = () => {
  const [courses, setCourses] = useState([]);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/courses`);
        const data = await response.json();
        setCourses(data);
      } catch (error) {
        console.error('Error al obtener los cursos:', error);
      }
    };

    fetchCourses();
  }, []);

  return (
    <div className="course-list">
      {courses.length > 0 ? (
        courses.map((course) => (
          <div key={course.id} className="course-item">
            <h3 className="course-title">
              <Link to={`/courses/${course.id}`}>{course.title}</Link>
            </h3>
            <p className="course-description">{course.description}</p>
            <p className="course-instructor">Instructor: {course.instructor}</p>
            <p className="course-duration">Duración: {course.duration} horas</p>
          </div>
        ))
      ) : (
        <p>No hay cursos disponibles.</p>
      )}
    </div>
  );
};

export default CourseList;