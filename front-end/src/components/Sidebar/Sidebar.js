// components/Sidebar.js
import React, {useEffect, useState} from 'react';
import {Link, useLocation} from 'react-router-dom';
import './Sidebar.css';
import Cookies from "universal-cookie";

const cookies = new Cookies();

const Sidebar = () => {
    const location = useLocation();
    const token = cookies.get("access_token");

    const [expandedChapters, setExpandedChapters] = useState({});
    const [expandedQuizzes, setExpandedQuizzes] = useState({});
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        const fetchCourses = async () => {
            try {
                const response = await fetch(`${process.env.REACT_APP_API_URL}/courses`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                    credentials: 'include',
                });
                const result = await response.json();

                if (response.status === 401) throw new Error('Unauthorized: 401');
                if (response.status === 500) throw new Error(result.message || 'Server error');
                if (!result.courses) throw new Error('Error getting courses list');

                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchCourses();
    }, []);

    // Expand if current location matches a chapter in course
    useEffect(() => {
        courses.forEach((course) => {
            // Expand chapters if any chapter link active
            if (
                course.chapters?.some(
                    (ch) =>
                        location.pathname.startsWith(`/courses/${course.id}/chapters/${ch.ID}`) &&
                        location.pathname.includes('/edit') === false // ignore edit here for chapter expand
                )
            ) {
                setExpandedChapters((prev) => ({...prev, [course.id]: true}));
            }

            // Expand quizzes if any quiz link active or quiz edit page active
            if (
                course.quizzes?.some(
                    (q) =>
                        location.pathname.startsWith(`/courses/${course.id}/quizzes/${q.ID}`) ||
                        location.pathname.startsWith(`/courses/${course.id}/quizzes/${q.ID}/edit`) ||
                        location.pathname.startsWith(`/courses/${course.id}/quizzes/create`) ||
                        location.pathname.startsWith(`/quizzes/create`)
                )
            ) {
                setExpandedQuizzes((prev) => ({...prev, [course.id]: true}));
            }
        });
    }, [location.pathname, courses]);

    const toggleChapters = (courseId) => {
        setExpandedChapters((prev) => ({
            ...prev,
            [courseId]: !prev[courseId],
        }));
    };

    const toggleQuizzes = (courseId) => {
        setExpandedQuizzes((prev) => ({
            ...prev,
            [courseId]: !prev[courseId],
        }));
    };

    const isActiveLink = (path) => location.pathname === path;

    const isActivePrefix = (prefix) => location.pathname.startsWith(prefix);

    const courseCreatePath = `/courses/create`;
    const courseCreateIsActive = isActiveLink(courseCreatePath);

    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <Link to="/" className="sidebar-logo">
                    Nursing AcademIQ
                </Link>
            </div>

            <nav className="sidebar-menu">
                <ul className="menu-list">
                    {token ? (
                        <>
                            <li className={isActiveLink('/my-courses') ? 'active' : ''}>
                                <Link to="/my-courses" className="menu-link">
                                    Διαχείριση Μαθημάτων
                                </Link>
                            </li>
                            <li className={isActiveLink('/my-quizzes') ? 'active' : ''}>
                                <Link to="/my-quizzes" className="menu-link">
                                    Διαχείριση Quiz
                                </Link>
                            </li>
                            <li className={isActiveLink('/courses') || location.pathname.startsWith('/courses/') ? 'active' : ''}>
                                <Link to="/courses" className="menu-link">
                                    Μαθήματα
                                </Link>
                            </li>
                            <li className={isActiveLink('/create-tutor') ? 'active' : ''}>
                                <Link to="/create-tutor" className="menu-link">
                                    Προσθήκη Καθηγητή
                                </Link>
                            </li>
                        </>
                    ) : (
                        <>
                            <li className={isActiveLink('/courses') || location.pathname.startsWith('/courses/') ? 'active' : ''}>
                                <Link to="/courses" className="menu-link">
                                    Μαθήματα
                                </Link>
                            </li>
                            <li className={isActiveLink('/quizzes') || location.pathname.startsWith('/quizzes') ? 'active' : ''}>
                                <Link to="/quizzes" className="menu-link">
                                    Quizzes
                                </Link>
                            </li>
                        </>
                    )}
                </ul>

                <ul className="menu-list account-menu">
                    {token ? (
                        <li className={isActiveLink('/profile') ? 'active' : ''}>
                            <Link to="/profile" className="menu-link">
                                Προφίλ
                            </Link>
                        </li>
                    ) : (
                        <li className={isActiveLink('/login') ? 'active' : ''}>
                            <Link to="/login" className="menu-link">
                                Σύνδεση
                            </Link>
                        </li>
                    )}
                </ul>
            </nav>

            <div className="sidebar-courses">
                <Link to="/courses" className="sidebar-section-title">
                    <div className="sidebar-section-title">Μαθήματα</div>
                </Link>

                {/* Show create course button if token exists */}
                {token && (
                    <Link to="/courses/create"
                          className={`btn-create  ${courseCreateIsActive ? 'active' : ''}`}>
                        Δημιουργία Μαθήματος
                    </Link>
                )}

                {(token
                        ? courses
                        : courses.filter(
                            (course) =>
                                (Array.isArray(course.chapters) &&
                                    course.chapters.length > 0) || (
                                    Array.isArray(course.quizzes) &&
                                    course.quizzes.length > 0)
                        )
                ).map((course) => {
                    const coursePath = `/courses/${course.id}`;
                    const courseIsActive = isActivePrefix(coursePath);

                    const courseEditPath = `/courses/${course.id}/edit`;
                    const courseEditIsActive = isActiveLink(courseEditPath);
                    const courseQuestionsManagePath = `/courses/${course.id}/questions/manage`;
                    const courseQuestionsManageIsActive = isActiveLink(courseQuestionsManagePath);

                    const chapterCreatePath = `/courses/${course.id}/chapters/create`;
                    const chapterCreateIsActive = isActiveLink(chapterCreatePath);

                    return (
                        <div
                            key={course.id}
                            className={`course-container ${courseIsActive ? 'active' : ''}`}
                        >
                            {/* Course title as Link */}
                            <Link to={coursePath} className={`course-title ${courseIsActive ? 'active' : ''}`}>
                                {course.title}
                            </Link>

                            {/* Edit button under course title if token */}
                            {token && (
                                <Link to={`/courses/${course.id}/edit`}
                                      className={`edit-btn  ${courseEditIsActive ? 'active' : ''}`}>
                                    Επεξεργασία Μαθήματος
                                </Link>
                            )}

                            {token && (
                                <Link to={`/courses/${course.id}/questions/manage`}
                                      className={`edit-btn  ${courseQuestionsManageIsActive ? 'active' : ''}`}>
                                    Διαχείριση Ερωτήσεων
                                </Link>
                            )}

                            {/* Chapters Section */}
                            <div className="sidebar-section">
                                <button
                                    onClick={() => toggleChapters(course.id)}
                                    className={`toggle-btn ${expandedChapters[course.id] ? 'expanded' : ''}`}
                                    aria-expanded={expandedChapters[course.id] ? 'true' : 'false'}
                                >
                                    Θεματικές Ενότητες
                                </button>

                                {expandedChapters[course.id] && (
                                    <>
                                        {/* Create chapter button if token */}
                                        {token && (
                                            <Link
                                                to={`/courses/${course.id}/chapters/create`}
                                                className={`btn-create  ${chapterCreateIsActive ? 'active' : ''}`}
                                            >
                                                Δημιουργία Ενότητας
                                            </Link>
                                        )}
                                        {course.chapters?.length > 0 && (
                                            <ul className="sidebar-list">
                                                {course.chapters.map((chapter) => {
                                                    const chapterPath = `/courses/${course.id}/chapters/${chapter.ID}/quizzes`;
                                                    const chapterEditPath = `/courses/${course.id}/chapters/${chapter.ID}/edit`;
                                                    const chapterIsActive = isActivePrefix(chapterPath) || isActiveLink(chapterEditPath);
                                                    const chapterEditIsActive = isActiveLink(chapterEditPath);

                                                    return (
                                                        <li key={chapter.ID} className="sidebar-list-item">
                                                            <Link
                                                                to={chapterPath}
                                                                className={`sidebar-link ${chapterIsActive ? 'active' : ''}`}
                                                            >
                                                                {chapter.Title}
                                                            </Link>

                                                            {/* Edit button under chapter if token */}
                                                            {token && (
                                                                <Link to={chapterEditPath}
                                                                      className={`edit-btn  ${chapterEditIsActive ? 'active' : ''}`}>
                                                                    Επεξεργασία
                                                                </Link>
                                                            )}
                                                        </li>
                                                    );
                                                })}
                                            </ul>
                                        )}
                                    </>
                                )}
                            </div>


                            {/* Quizzes Section */}
                            <div className="sidebar-section">
                                <button
                                    onClick={() => toggleQuizzes(course.id)}
                                    className={`toggle-btn ${expandedQuizzes[course.id] ? 'expanded' : ''}`}
                                    aria-expanded={expandedQuizzes[course.id] ? 'true' : 'false'}
                                >
                                    Quizzes
                                </button>

                                {expandedQuizzes[course.id] && (
                                    <>
                                        {/* Create quiz buttons if token */}
                                        {token && (
                                            <>
                                                <Link
                                                    to={`/courses/${course.id}/quizzes/create`}
                                                    className="btn-create"
                                                >
                                                    Δημιουργία Quiz
                                                </Link>
                                            </>
                                        )}
                                        {course.quizzes?.length > 0 && (
                                            <ul className="sidebar-list">
                                                {(token
                                                        ? course.quizzes
                                                        : course.quizzes.filter((quiz) => quiz.Visibility === true)
                                                ).map((quiz) => {
                                                    const quizPath = `/courses/${course.id}/quizzes/${quiz.ID}`;
                                                    const quizEditPaths = [
                                                        `/courses/${course.id}/quizzes/${quiz.ID}/edit`,
                                                        `/courses/${course.id}/quizzes/${quiz.ID}/edit/step-two`,
                                                        `/courses/${course.id}/quizzes/${quiz.ID}/edit/step-three`,
                                                    ];
                                                    const quizIsActive =
                                                        isActivePrefix(quizPath) ||
                                                        quizEditPaths.some((p) => isActivePrefix(p));

                                                    return (
                                                        <li key={quiz.ID} className="sidebar-list-item">
                                                            <Link
                                                                to={quizPath}
                                                                className={`sidebar-link ${quizIsActive ? 'active' : ''}`}
                                                            >
                                                                {quiz.Title}
                                                            </Link>

                                                            {/* Edit button under quiz if token */}
                                                            {token && (
                                                                <Link to={`${quizPath}/edit`} className="edit-btn">
                                                                    Επεξεργασία
                                                                </Link>
                                                            )}
                                                        </li>
                                                    );
                                                })}
                                            </ul>
                                        )}
                                    </>
                                )}
                            </div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
};

export default Sidebar;
