// components/Sidebar.js
import React, {useEffect, useState} from 'react';
import { Link } from 'react-router-dom';
import './Sidebar.css';

const Sidebar = () => {
    const [expandedChapters, setExpandedChapters] = useState({});
    const [expandedQuizzes, setExpandedQuizzes] = useState({});
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchCourses = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/courses`;

            try {
                const response = await fetch(apiUrl, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                    credentials: 'include',
                });
                const result = await response.json();
                // TODO if 401 show unauthorized
                // TODO if 500 show server error
                if (response.status === 500) {
                    throw Error(result.message);
                }

                if (response.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (result.courses === undefined) {
                    throw Error("error getting courses list");
                }
                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchCourses();
    }, []);

    const toggleChapters = (courseId) => {
        setExpandedChapters(prev => ({
            ...prev,
            [courseId]: !prev[courseId]
        }));
    };

    const toggleQuizzes = (courseId) => {
        setExpandedQuizzes(prev => ({
            ...prev,
            [courseId]: !prev[courseId]
        }));
    };

    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <Link to={`/`} className="sidebar-link">
                    <h1 className="app-name">Nursing AcademIQ</h1>
                </Link>
            </div>
            <div className="sidebar-content">
                <h2 className="sidebar-title">Μαθημάτα</h2>
                <nav className="sidebar-nav">
                    {courses.map(course => (
                        <div key={course.id} className="sidebar-course">
                            <div className="sidebar-course-title">{course.title}</div>

                            {course.chapters?.length > 0 && (
                                <div className="sidebar-section">
                                    <button
                                        className="toggle-section-btn"
                                        onClick={() => toggleChapters(course.id)}
                                    >
                                        {expandedChapters[course.id] ? '▲' : '▼'} Θεματικές Ενότητες
                                    </button>
                                    {expandedChapters[course.id] && (
                                        <ul className="sidebar-list">
                                            {course.chapters.map(chapter => (
                                                <li key={chapter.ID}>
                                                    <Link to={`/courses/${course.id}/chapters/${chapter.ID}/quizzes`} className="sidebar-link">
                                                        {chapter.Title}
                                                    </Link>
                                                </li>
                                            ))}
                                        </ul>
                                    )}
                                </div>
                            )}

                            {course.quizzes?.length > 0 && (
                                <div className="sidebar-section">
                                    <button
                                        className="toggle-section-btn"
                                        onClick={() => toggleQuizzes(course.id)}
                                    >
                                        {expandedQuizzes[course.id] ? '▲' : '▼'} Quizzes
                                    </button>
                                    {expandedQuizzes[course.id] && (
                                        <ul className="sidebar-list">
                                            {course.quizzes.map(quiz => (
                                                <li key={quiz.ID}>
                                                    <Link to={`/courses/${course.id}/quizzes/${quiz.ID}`} className="sidebar-link">
                                                        {quiz.Title}
                                                    </Link>
                                                </li>
                                            ))}
                                        </ul>
                                    )}
                                </div>
                            )}
                        </div>
                    ))}
                </nav>
            </div>
        </div>
    );
};

export default Sidebar;
