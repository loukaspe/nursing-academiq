import {useQuiz} from "../../context/QuizContext";
import {Link, useParams} from "react-router-dom";
import CreationProgressBar from "./CreationProgressBar";
import React, {useEffect} from "react";
import "./CreateQuizStepTwo.css";
import axios from "axios";
import EditProgressBar from "./EditProgressBar";
import Breadcrumb from "../Utilities/Breadcrumb";

export default function EditQuizStepOne() {
    const {quiz, setQuiz} = useQuiz();

    const params = useParams();
    let quizID = params.id;

    useEffect(() => {
        const fetchQuiz = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/quiz/${quizID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                console.log(response.data.quiz)
                setQuiz(prevQuiz => ({
                    ...prevQuiz,
                    id: response.data.quiz.ID,
                    title: response.data.quiz.Title,
                    description: response.data.quiz.Description,
                    course: response.data.quiz.Course,
                    isVisible: response.data.quiz.Visibility,
                    isShowSubsetChecked: response.data.quiz.ShowSubset,
                    subsetSize: response.data.quiz.SubsetSize || null,
                    questions: response.data.quiz.Questions || null
                }));
            } catch (error) {
                console.error('Error fetching the quiz data', error);
            }
        };

        fetchQuiz();
    }, [quizID]);

    return (
        <div className="createQuizStepTwoContainer">
            {quiz.course && (
                <Breadcrumb
                    actualPath={`/courses/${quiz.course.ID}/quizzes/${quiz.id}/edit`}
                    namePath={`/Μαθήματα/${quiz.course.Title}/Quiz/${quiz.title}/Επεξεργασία`}
                />
            )}
            <EditProgressBar/>
            <div className="createQuizStepTwoHeaderRow">
                <div className="createQuizStepTwoHeader">
                    <div className="createQuizStepTwoInfo">
                        <span className="singleChapterQuizzesPageChapterName">1. Λεπτομέρειες Quiz</span>
                    </div>
                </div>
            </div>
            <div className="createQuizStepTwoDetailsRow">
                <div className="createQuizStepTwoDetailsRowColumn">
                    <div className="createQuizStepTwoDetailsRowInputGroup">
                        <label>*Όνομα Quiz</label>
                        <input type="text"
                               value={quiz.title}
                               className="createQuizStepTwoDetailsRowInputText"
                               onChange={(e) => setQuiz({...quiz, title: e.target.value || ""})}
                        />
                    </div>
                    <div className="createQuizStepTwoDetailsRowInputGroup">
                        <label>*Περιγραφή</label>
                        <input type="text"
                               value={quiz.description}
                               className="createQuizStepTwoDetailsRowInputText"
                               onChange={(e) => setQuiz({...quiz, description: e.target.value || ""})}
                        />
                    </div>
                </div>
            </div>

            <div className="createQuizStepTwoButtonsContainer">
                <Link
                    className={`createQuizStepTwoButton ${(quiz.title.trim() === "" || quiz.description.trim() === "") ? "disabled" : ""}`}
                    to={(quiz.title && quiz.description) ? `/courses/${quiz.course.ID}/quizzes/${quizID}/edit/step-two` : "#"}
                    onClick={(e) => {
                        if (quiz.title.trim() === "" || quiz.description.trim() === "") {
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
