import { useLocation, Link } from "react-router-dom";
import { useQuiz } from "../../context/QuizContext";
import "./CreationProgressBar.css";

const steps = [
    { path: "/quizzes/create", label: "1. Επιλογή Μαθήματος" },
    { path: "/quizzes/create/step-two", label: "2. Λεπτομέρειες Quiz" },
    { path: "/quizzes/create/step-three", label: "3. Ερωτήσεις" },
    { path: "/quizzes/create/step-four", label: "4. Ολοκλήρωση" },
];

export default function CreationProgressBar() {
    const location = useLocation();
    const { quiz } = useQuiz();

    const currentStepIndex = steps.findIndex((step) => step.path === location.pathname);
    const currentStep = steps[currentStepIndex];
    const hasCourseSelected = quiz.course !== null;

    return (
        <div className="step-progress">
            {steps.map((step, index) => {
                const isActive = index === currentStepIndex;
                const isDisabled = index > 0 && !hasCourseSelected;
                
                return (
                    <Link
                        key={step.path}
                        to={isDisabled ? "#" : step.path}
                        className={`step ${isActive ? "active" : ""} ${isDisabled ? "disabled" : ""}`}
                        onClick={(e) => {
                            if (isDisabled) {
                                e.preventDefault();
                            }
                        }}
                    >
                        {step.label}
                    </Link>
                );
            })}
        </div>
    );
}
