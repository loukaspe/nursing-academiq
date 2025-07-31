// components/Sidebar.js
import React, {useEffect, useState} from 'react';
import {Link} from 'react-router-dom';
import './sidebar.css';

const Sidebar = () => {
    const itemsToShow = 2

    const [expandedQuizzes, setExpandedQuizzes] = useState({});
    const [expandedChapters, setExpandedChapters] = useState({});

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

    const toggleQuizzes = (courseId) => {
        setExpandedQuizzes(prev => ({
            ...prev,
            [courseId]: !prev[courseId]
        }));
    };

    const toggleChapters = (courseId) => {
        setExpandedChapters(prev => ({
            ...prev,
            [courseId]: !prev[courseId]
        }));
    };

    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <Link className="app-name" to={`/`}>
                    <h1 className="app-name">
                        Nursing Academiq
                    </h1>
                </Link>
            </div>
            <div className="sidebar-content">
                <h2 className="sidebar-title">Mαθήματα</h2>
                <nav className="sidebar-nav">
                    {courses.map(course => {
                        const quizzesToShow = expandedQuizzes[course.id]
                            ? course.quizzes
                            : course.quizzes?.slice(0, itemsToShow) || [];

                        const chaptersToShow = expandedChapters[course.id]
                            ? course.chapters
                            : course.chapters?.slice(0, itemsToShow) || [];

                        return (
                            <div key={course.id} className="sidebar-course">
                                <div className="sidebar-course-title">{course.title}</div>

                                {course.chapters?.length > 0 && (
                                    <div className="sidebar-section">
                                        <div className="sidebar-subtitle">Θεματικές Ενότητες</div>
                                        <ul className="sidebar-list">
                                            {chaptersToShow.map(chapter => (
                                                <li key={chapter.ID}>
                                                    <Link to={`/courses/${course.id}/chapter/${chapter.ID}`}
                                                          className="sidebar-link">
                                                        {chapter.Title}
                                                    </Link>
                                                </li>
                                            ))}
                                        </ul>

                                        {course.chapters.length > 2 && (
                                            <button
                                                className="toggle-quizzes-btn"
                                                onClick={() => toggleChapters(course.id)}
                                            >
                                                {expandedChapters[course.id] ? '▲ Λιγότερα' : '▼ Περισσότερα'}
                                            </button>
                                        )}
                                    </div>
                                )}

                                {course.quizzes?.length > 0 && (
                                    <div className="sidebar-section">
                                        <div className="sidebar-subtitle">Quizzes</div>
                                        <ul className="sidebar-list">
                                            {quizzesToShow.map(quiz => (
                                                <li key={quiz.ID}>
                                                    <Link to={`/courses/${course.id}/quizzes/${quiz.ID}`}
                                                          className="sidebar-link">
                                                        {quiz.Title}
                                                    </Link>
                                                </li>
                                            ))}
                                        </ul>

                                        {course.quizzes.length > 2 && (
                                            <button
                                                className="toggle-quizzes-btn"
                                                onClick={() => toggleQuizzes(course.id)}
                                            >
                                                {expandedQuizzes[course.id] ? '▲ Λιγότερα' : '▼ Περισσότερα'}
                                            </button>
                                        )}
                                    </div>
                                )}
                            </div>
                        );
                    })}
                </nav>
            </div>
        </div>
    );
};

export default Sidebar;
