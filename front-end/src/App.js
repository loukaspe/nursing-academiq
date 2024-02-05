import React from 'react';
import {Routes, Route, Link} from "react-router-dom";
import LimitedCoursesList from "./components/CoursesList/LimitedCoursesList";
import Homepage from "./components/Homepage/Homepage";
import Layout from "./components/Layout/Layout";
import QuestionsWrapper from "./components/Questions/QuestionsWrapper/QuestionsWrapper";
import {questions} from "./questions";
import ProtectedRoutes from "./routes/ProtectedRoute";
import Cookies from "universal-cookie";
import LoginPage from "./components/Login/LoginPage";
import CoursesList from "./components/CoursesList/CoursesList";
import QuizzesList from "./components/QuizzesList/QuizzesList";

const cookies = new Cookies();


const App = () => {
    return (
        <>
            <Routes>
                <Route path="/" element={<Layout/>}>
                    <Route
                        index
                        element={
                            <ProtectedRoutes>
                                <Homepage/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="questions"
                        element={
                            <ProtectedRoutes>
                                <QuestionsWrapper
                                    questions={questions}
                                />
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="courses"
                        element={
                            <ProtectedRoutes>
                                <CoursesList/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route
                        path="quizzes"
                        element={
                            <ProtectedRoutes>
                                <QuizzesList/>
                            </ProtectedRoutes>
                        }
                    />
                    <Route path="login" element={<LoginPage/>}/>
                    {/*<Route path="loginForm" element={<LoginForm/>}/>*/}
                    <Route path="logout" element={<Logout/>}/>

                    <Route path="*" element={<NotFound/>}/>
                </Route>
            </Routes>
        </>

        // <>
        //     <Header/>
        //     <div style={{width: '100%'}}>
        //         <div style={{float: 'left', width: '50%'}}>
        //             <CoursesList coursesList={coursesList}/>
        //         </div>
        //         <div style={{float: 'right', width: '50%'}}>
        //             <LoginForm/>
        //         </div>
        //     </div>
        //     <div style={{clear: 'both'}}></div>
        // </>
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
    cookies.remove("user", {path: "/"});
    window.location.href = "/";
}

export default App