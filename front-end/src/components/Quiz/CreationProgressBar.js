import { useLocation, Link } from "react-router-dom";
import "./CreationProgressBar.css";

const steps = [
    { path: "/quizzes/create", label: "1. Επιλογή Μαθήματος" },
    { path: "/quizzes/create/step-two", label: "2. Λεπτομέρειες Quiz" },
    { path: "/quizzes/create/step-three", label: "3. Ερωτήσεις" },
    { path: "/quizzes/create/step-four", label: "4. Ολοκλήρωση" },
];

export default function CreationProgressBar() {
    const location = useLocation();

    const currentStepIndex = steps.findIndex((step) => step.path === location.pathname);
    const currentStep = steps[currentStepIndex];

    return (
        <div className="step-progress">
            {steps.map((step, index) => {
                const isActive = index === currentStepIndex;
                return (
                    <Link
                        key={step.path}
                        to={step.path}
                        className={`step ${isActive ? "active" : ""}`}
                    >
                        {step.label}
                    </Link>
                );
            })}
        </div>
    );
}
