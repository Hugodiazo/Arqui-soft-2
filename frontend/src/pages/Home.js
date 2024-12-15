import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './Home.css';

function Home() {
  const [query, setQuery] = useState('');
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCourses = async () => {
      setLoading(true);
      setError(null);
      try {
        const response = await fetch('http://localhost:8081/courses');
        if (!response.ok) {
          throw new Error('Error al obtener los cursos');
        }
        const data = await response.json();
        setCourses(Array.isArray(data) ? data : []);
      } catch (error) {
        console.error('Error al obtener los cursos:', error);
        setError('Hubo un problema al obtener los cursos');
        setCourses([]);
      } finally {
        setLoading(false);
      }
    };

    fetchCourses();
  }, []);

  const handleSearch = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const response = query.trim()
        ? await fetch(`http://localhost:8083/search?q=${encodeURIComponent(query)}`)
        : await fetch('http://localhost:8081/courses');

      if (!response.ok) {
        throw new Error('Error al realizar la búsqueda');
      }

      const data = await response.json();

      // Asegúrate de que los datos sean un array
      const formattedData = Array.isArray(data)
        ? data.map(course => ({
            ID: course.ID || course.id,
            Title: course.Title || course.title,
            Description: course.Description || course.description,
            Instructor: course.Instructor || course.instructor,
            Duration: course.Duration || course.duration,
            Level: course.Level || course.level,
          }))
        : [];

      setCourses(formattedData);
    } catch (error) {
      console.error('Error al buscar cursos:', error);
      setError('Hubo un problema al realizar la búsqueda');
      setCourses([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="home">
      <h2>Bienvenido a la Plataforma de Cursos</h2>
      <form onSubmit={handleSearch} className="search-form">
        <input
          type="text"
          placeholder="Buscar cursos..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        <button type="submit">Buscar</button>
      </form>

      {loading && <p>Cargando...</p>}
      {error && <p className="error">{error}</p>}

      <div className="courses">
        {Array.isArray(courses) && courses.length > 0 ? (
          courses.map((course) => (
            <div key={course.ID} className="course-item">
              <Link to={`/courses/${course.ID}`}>
                <h3>{course.Title}</h3>
              </Link>
              <p>{course.Description}</p>
              <p>Instructor: {course.Instructor}</p>
              <p>Duración: {course.Duration} horas</p>
              <p>Nivel: {course.Level}</p>
            </div>
          ))
        ) : (
          !loading && <p>No se encontraron cursos.</p>
        )}
      </div>
    </div>
  );
}

export default Home;