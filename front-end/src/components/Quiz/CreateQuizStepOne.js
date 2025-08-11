import {useQuiz} from "../../context/QuizContext";
import {Link} from "react-router-dom";
import React, {useEffect, useState} from "react";
import CreationProgressBar from "./CreationProgressBar";
import "./CreateQuizStepOne.css";
import api from "../Utilities/APICaller";
import Cookies from "universal-cookie";
import Breadcrumb from "../Utilities/Breadcrumb";

export default function CreateQuizStepOne() {
    const {quiz, setQuiz} = useQuiz();
    const [courses, setCourses] = useState([]);

    const cookies = new Cookies();
    useEffect(() => {
        const fetchTutorCourses = async () => {
            let userCookie = cookies.get("user");
            let specificID = userCookie.specificID;

            let apiUrl = `/tutor/${specificID}/courses`;

            try {
                const response = await api.get(apiUrl);

                // TODO if 401 show unauthorized
                // TODO if 500 show server error
                if (response.status === 500) {
                    throw Error(response.data.message);
                }

                if (response.status === 401) {
                    throw Error("unauthorized: 401");
                }

                if (response.data.courses === undefined) {
                    throw Error("error getting courses for tutor");
                }
                setCourses(response.data.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchTutorCourses();

    }, []);

    const handleCourseChange = (e) => {
        const selectedId = parseInt(e.target.value, 10);
        const selectedCourse = courses.find((c) => c.id === selectedId);
        setQuiz({...quiz, course: selectedCourse || null});
    };

    return (
        <div>
            <Breadcrumb actualPath={`/quizzes/create`} namePath={`/Quiz/Δημιουργία - Βήμα 1`}/>
            <CreationProgressBar/>
            <div className="create-quiz-step-one-content">
                <h2 className="create-quiz-step-one-page-title">1. Επιλογή Μαθήματος</h2>
                <select
                    value={quiz.course?.id || ""}
                    onChange={handleCourseChange}
                >
                    <option value="">-- Επιλέξτε μάθημα --</option>
                    {courses.map((course) => (
                        <option key={course.id} value={course.id}>
                            {course.title}
                        </option>
                    ))}
                </select>

                <Link
                    className={`create-quiz-step-one-content-button ${!quiz.course ? "disabled" : ""}`}
                    to={quiz.course ? "/quizzes/create/step-two" : "#"}
                    onClick={(e) => {
                        if (!quiz.course) {
                            e.preventDefault();
                        }
                    }}
                >
                    Επόμενο
                </Link>
            </div>
        </div>
    );
}
