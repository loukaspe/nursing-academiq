// components/Sidebar.js
import React, { useEffect, useState } from "react";
import { Link, useLocation } from "react-router-dom";
import "./Sidebar.css";
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
                const response = await fetch(process.env.REACT_APP_API_URL + `/courses`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                    credentials: "include",
                });
                const result = await response.json();
                if (!response.ok) throw new Error(result.message || "Failed to fetch courses");
                if (!result.courses) throw new Error("No courses found");
                setCourses(result.courses);
            } catch (error) {
                console.error("Error fetching courses:", error);
            }
        };

        fetchCourses();
    }, []);

    const toggleChapters = (courseId) => {
        setExpandedChapters((prev) => ({ ...prev, [courseId]: !prev[courseId] }));
    };

    const toggleQuizzes = (courseId) => {
        setExpandedQuizzes((prev) => ({ ...prev, [courseId]: !prev[courseId] }));
    };

    const isActive = (path) => {
        if (!path) return false;
        if (location.pathname === path) return true;
        if (location.pathname.startsWith(path + "/")) return true;
        return false;
    };

    useEffect(() => {
        // Auto-expand chapters/quizzes if active
        const newExpandedChapters = {};
        const newExpandedQuizzes = {};
        courses.forEach((course) => {
            if (
                course.chapters?.some(
                    (ch) =>
                        isActive(`/courses/${course.id}/chapters/${ch.ID}`) ||
                        isActive(`/courses/${course.id}/chapters/${ch.ID}/edit`)
                )
            ) {
                newExpandedChapters[course.id] = true;
            }
            if (
                course.quizzes?.some(
                    (quiz) =>
                        isActive(`/courses/${course.id}/quizzes/${quiz.ID}`) ||
                        isActive(`/courses/${course.id}/quizzes/${quiz.ID}/edit`)
                )
            ) {
                newExpandedQuizzes[course.id] = true;
            }
        });
        setExpandedChapters((prev) => ({ ...prev, ...newExpandedChapters }));
        setExpandedQuizzes((prev) => ({ ...prev, ...newExpandedQuizzes }));
    }, [location.pathname, courses]);

    return (
        <aside className="sidebar">
            <header className="sidebar-header">
                <Link to="/" className="sidebar-logo">
                    Nursing AcademIQ
                </Link>
            </header>

            <section className="sidebar-menu">
                {token ? (
                    <>
                        <ul className="menu-list">
                            <li className={isActive("/my-courses") ? "active" : ""}>
                                <Link to="/my-courses" className="menu-link">
                                    Διαχείριση Μαθημάτων
                                </Link>
                            </li>
                            <li className={isActive("/my-quizzes") ? "active" : ""}>
                                <Link to="/my-quizzes" className="menu-link">
                                    Διαχείριση Quiz
                                </Link>
                            </li>
                            <li className={isActive("/courses") ? "active" : ""}>
                                <Link to="/courses" className="menu-link">
                                    Μαθήματα
                                </Link>
                            </li>
                            <li className={isActive("/create-tutor") ? "active" : ""}>
                                <Link to="/create-tutor" className="menu-link">
                                    Προσθήκη Καθηγητή
                                </Link>
                            </li>
                        </ul>
                        <ul className="menu-list account-menu">
                            <li className={isActive("/profile") ? "active" : ""}>
                                <Link to="/profile" className="menu-link">
                                    Προφίλ
                                </Link>
                            </li>
                        </ul>
                    </>
                ) : (
                    <>
                        <ul className="menu-list">
                            <li className={isActive("/courses") ? "active" : ""}>
                                <Link to="/courses" className="menu-link">
                                    Μαθήματα
                                </Link>
                            </li>
                            <li className={isActive("/quizzes") ? "active" : ""}>
                                <Link to="/quizzes" className="menu-link">
                                    Quizzes
                                </Link>
                            </li>
                        </ul>
                        <ul className="menu-list account-menu">
                            <li className={isActive("/login") ? "active" : ""}>
                                <Link to="/login" className="menu-link">
                                    Σύνδεση
                                </Link>
                            </li>
                        </ul>
                    </>
                )}
            </section>

            <section className="sidebar-courses">
                <h2 className="sidebar-section-title">Μαθήματα</h2>

                {/* Create Course button */}
                <Link
                    to="/courses/create"
                    className={`btn-create ${isActive("/courses/create") ? "active" : ""}`}
                >
                    Δημιουργία Μαθήματος
                </Link>

                {courses.map((course) => {
                    const coursePath = `/courses/${course.id}`;
                    const courseActive = isActive(coursePath);

                    return (
                        <div key={course.id} className="course-container">
                            <Link
                                to={coursePath}
                                className={`course-title ${courseActive ? "active" : ""}`}
                            >
                                {course.title}
                            </Link>

                            <div className="btn-group-create">
                                <Link
                                    to={`/courses/${course.id}/chapters/create`}
                                    className={`btn-create small ${
                                        isActive(`/courses/${course.id}/chapters/create`) ? "active" : ""
                                    }`}
                                >
                                    Δημιουργία Θεματικής Ενότητας
                                </Link>
                                <Link
                                    to={`/courses/${course.id}/quizzes/create`}
                                    className={`btn-create small ${
                                        isActive(`/courses/${course.id}/quizzes/create`) ? "active" : ""
                                    }`}
                                >
                                    Δημιουργία Quiz
                                </Link>
                            </div>

                            {/* Chapters */}
                            {course.chapters?.length > 0 && (
                                <div className="sidebar-section">
                                    <button
                                        className="toggle-btn"
                                        onClick={() => toggleChapters(course.id)}
                                        aria-expanded={expandedChapters[course.id] ? "true" : "false"}
                                    >
                                        {expandedChapters[course.id] ? "▲" : "▼"} Θεματικές Ενότητες
                                    </button>
                                    {expandedChapters[course.id] && (
                                        <ul className="sidebar-list">
                                            {course.chapters.map((chapter) => {
                                                const chapterPath = `/courses/${course.id}/chapters/${chapter.ID}`;
                                                const chapterEditPath = `${chapterPath}/edit`;
                                                return (
                                                    <li key={chapter.ID} className="sidebar-list-item">
                                                        <Link
                                                            to={`${chapterPath}/quizzes`}
                                                            className={`sidebar-link ${
                                                                isActive(chapterPath) ? "active" : ""
                                                            }`}
                                                        >
                                                            {chapter.Title}
                                                        </Link>
                                                        <Link
                                                            to={chapterEditPath}
                                                            className={`edit-btn ${
                                                                isActive(chapterEditPath) ? "active" : ""
                                                            }`}
                                                        >
                                                            Επεξεργασία
                                                        </Link>
                                                    </li>
                                                );
                                            })}
                                        </ul>
                                    )}
                                </div>
                            )}

                            {/* Quizzes */}
                            {course.quizzes?.length > 0 && (
                                <div className="sidebar-section">
                                    <button
                                        className="toggle-btn"
                                        onClick={() => toggleQuizzes(course.id)}
                                        aria-expanded={expandedQuizzes[course.id] ? "true" : "false"}
                                    >
                                        {expandedQuizzes[course.id] ? "▲" : "▼"} Quizzes
                                    </button>
                                    {expandedQuizzes[course.id] && (
                                        <ul className="sidebar-list">
                                            {course.quizzes.map((quiz) => {
                                                const quizPath = `/courses/${course.id}/quizzes/${quiz.ID}`;
                                                const quizEditPath = `/courses/${course.id}/quizzes/${quiz.ID}/edit`;
                                                return (
                                                    <li key={quiz.ID} className="sidebar-list-item">
                                                        <Link
                                                            to={quizPath}
                                                            className={`sidebar-link ${isActive(quizPath) ? "active" : ""}`}
                                                        >
                                                            {quiz.Title}
                                                        </Link>
                                                        <Link
                                                            to={quizEditPath}
                                                            className={`edit-btn ${isActive(quizEditPath) ? "active" : ""}`}
                                                        >
                                                            Επεξεργασία
                                                        </Link>
                                                    </li>
                                                );
                                            })}
                                        </ul>
                                    )}
                                </div>
                            )}
                        </div>
                    );
                })}
            </section>
        </aside>
    );
};

export default Sidebar;
