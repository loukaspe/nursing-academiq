import React from 'react';
import {Routes, Route, Link} from "react-router-dom";
import Homepage from "./components/Homepage/Homepage";
import Layout from "./components/Layout/Layout";
import QuestionsWrapper from "./components/Questions/QuestionsWrapper";
import {questions} from "./questions";
import ProtectedRoutes from "./routes/ProtectedRoute";
import Cookies from "universal-cookie";
import LoginPage from "./components/Login/LoginPage";
import CoursesList from "./components/CoursesList/CoursesList";
import UserProfile from "./components/UserProfile/UserProfile";
import ChangePassword from "./components/ChangePassword/ChangePassword";
import SingleCourse from "./components/Course/SingleCourse";
import CourseChaptersList from "./components/ChaptersList/CourseChaptersList";
import CourseQuizzesList from "./components/QuizzesList/CourseQuizzesList";
import CsvImport from "./components/Questions/CsvImport";
import ChapterQuizzesList from "./components/ChaptersList/ChapterQuizzesList";
import QuizStart from "./components/Quiz/QuizStart";
import MyCoursesList from "./components/CoursesList/MyCoursesList";
import MyQuizzesList from "./components/QuizzesList/MyQuizzesList";
import EditCourse from "./components/Course/EditCourse";
import CreateCourse from "./components/Course/CreateCourse";
import EditChapter from "./components/Chapter/EditChapter";
import CreateChapter from "./components/Chapter/CreateChapter";
import EditQuiz from "./components/Quiz/EditQuiz";
import CreateQuiz from "./components/Quiz/CreateQuiz";
import QuestionsManager from "./components/Questions/QuestionsManager";
import EditQuestion from "./components/Questions/EditQuestion";
import CreateQuestion from "./components/Questions/CreateQuestion";

const cookies = new Cookies();


const App = () => {
    return (
        <>
            <Routes>
                <Route path="/" element={<Layout/>}>
                    <Route index element={<Homepage/>}/>
                    {/* Courses */}
                    <Route
                        path="my-courses"
                        element={
                            <ProtectedRoutes>
                                <MyCoursesList/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses"
                        element={
                            <CoursesList/>
                        }
                    />
                    <Route
                        path="courses/:id"
                        element={
                            <SingleCourse/>
                        }
                    />
                    <Route
                        path="courses/:id/edit"
                        element={
                            <ProtectedRoutes>
                                <EditCourse/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/create"
                        element={
                            <ProtectedRoutes>
                                <CreateCourse/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:id/chapters"
                        element={
                            <CourseChaptersList/>
                        }
                    />
                    <Route
                        path="courses/:id/quizzes"
                        element={
                            <CourseQuizzesList/>
                        }
                    />
                    {/* Chapters */}
                    <Route
                        path="courses/:courseID/chapters/:chapterID/quizzes"
                        element={
                            <ChapterQuizzesList/>
                        }
                    />
                    <Route
                        path="courses/:courseID/chapters/:id/edit"
                        element={
                            <ProtectedRoutes>
                                <EditChapter/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/chapters/create"
                        element={
                            <ProtectedRoutes>
                                <CreateChapter/>
                            </ProtectedRoutes>
                        }
                    />
                    {/* Quizzes */}
                    <Route
                        path="my-quizzes"
                        element={
                            <ProtectedRoutes>
                                <MyQuizzesList/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/quizzes/:quizID"
                        element={
                            <QuizStart/>
                        }
                    />
                    <Route
                        path="courses/:courseID/quizzes/:quizID/complete"
                        element={
                            <QuestionsWrapper/>
                        }
                    />
                    <Route
                        path="courses/:courseID/quizzes/:id/edit"
                        element={
                            <ProtectedRoutes>
                                <EditQuiz/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/quizzes/create"
                        element={
                            <ProtectedRoutes>
                                <CreateQuiz/>
                            </ProtectedRoutes>
                        }
                    />
                    {/* Questions */}
                    <Route
                        path="questions"
                        element={
                            <QuestionsWrapper
                                questions={questions}
                            />
                        }
                    />
                    <Route
                        path="courses/:id/questions/import"
                        element={
                            <ProtectedRoutes>
                                <CsvImport/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/questions/manage"
                        element={
                            <ProtectedRoutes>
                                <QuestionsManager/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/chapters/:chapterID/questions/:id/edit"
                        element={
                            <ProtectedRoutes>
                                <EditQuestion/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses/:courseID/chapters/:chapterID/questions/create"
                        element={
                            <ProtectedRoutes>
                                <CreateQuestion/>
                            </ProtectedRoutes>
                        }
                    />
                    {/* User */}
                    <Route
                        path="profile"
                        element={
                            <ProtectedRoutes>
                                <UserProfile/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="change-password"
                        element={
                            <ProtectedRoutes>
                                <ChangePassword/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route path="login" element={<LoginPage/>}/>
                    <Route path="logout" element={<Logout/>}/>

                    <Route path="*" element={<NotFound/>}/>
                </Route>
            </Routes>
        </>
    )
}

function NotFound() {
    return (
        <div>
            <h2>Nothing to see here!</h2>
            <p>
                <Link to="/">Go to the home page</Link>
            </p>
        </div>
    );
}

function Logout() {
    cookies.remove("token", {path: "/"});
    cookies.remove("result", {path: "/"});
    cookies.remove("user", {path: "/"});
    window.location.href = "/";
}

export default App