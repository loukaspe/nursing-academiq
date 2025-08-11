import React, {useEffect, useState} from "react";
import "./QuizStart.css";
import {Link, useNavigate, useParams} from "react-router-dom";
import Breadcrumb from "../Utilities/Breadcrumb";

const QuizStart = () => {
    const [quiz, setQuiz] = useState({});
    const [course, setCourse] = useState({});

    const params = useParams();
    let quizID = params.quizID;
    let courseID = params.courseID;

    let navigate = useNavigate();

    useEffect(() => {
        const fetchQuiz = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

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

                if (result.quiz === undefined) {
                    throw Error("error getting quiz questions");
                }
                setQuiz(result.quiz);

                if (result.quiz.Course === undefined) {
                    throw Error("error getting quiz questions");
                }
                setCourse(result.quiz.Course);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        }

        fetchQuiz();
    }, []);

    const handleSubmit = () => {

    };

    return (
        <React.Fragment>
            <Breadcrumb
                actualPath={`/courses/${courseID}/quizzes/${quizID}`}
                namePath={`/Μαθήματα/${course.Title}/Quiz/${quiz.Title}`}
            />
            <div className="singleQuizName">
                {
                    quiz.ShowSubset
                        ? `${quiz.Title} - ${quiz.SubsetSize} Ερωτήσεις`
                        : `${quiz.Title} - ${quiz.NumberOfQuestions} Ερωτήσεις`
                }
            </div>
            <div className="quiz-container">
                <div className="questionCard">
                    <div className="quizDescription">{quiz.Description}</div>
                    <div className="quizButtons">
                        <div className="quizButtonsRow">
                            <Link className="quizSmallButton"
                                  to={`/courses/${courseID}/quizzes/${quizID}/complete`}>Έναρξη</Link>
                        </div>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );
};


export default QuizStart;