import { createContext, useState, useContext } from "react";

const QuizContext = createContext();

const initialQuizState = {
    course: null,
    title: "",
    description: "",
    isVisible: false,
    isShowSubsetChecked: false,
    subsetSize: 0,
    questions: [],
};

export function QuizProvider({ children }) {
    const [quiz, setQuiz] = useState(initialQuizState);

    const resetQuiz = () => setQuiz(initialQuizState);

    return (
        <QuizContext.Provider value={{ quiz, setQuiz, resetQuiz }}>
            {children}
        </QuizContext.Provider>
    );
}

export const useQuiz = () => useContext(QuizContext);
