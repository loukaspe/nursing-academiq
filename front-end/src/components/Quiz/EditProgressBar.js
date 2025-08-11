import { useLocation, Link, useParams } from "react-router-dom";
import "./CreationProgressBar.css";

export default function EditProgressBar() {
    const location = useLocation();
    const { courseID, id } = useParams(); // extract IDs from the URL

    const steps = [
        { path: `/courses/${courseID}/quizzes/${id}/edit`, label: "1. Λεπτομέρειες Quiz" },
        { path: `/courses/${courseID}/quizzes/${id}/edit/step-two`, label: "2. Ερωτήσεις" },
        { path: `/courses/${courseID}/quizzes/${id}/edit/step-three`, label: "3. Ολοκλήρωση" },
    ];

    const currentStepIndex = steps.findIndex((step) => step.path === location.pathname);

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
