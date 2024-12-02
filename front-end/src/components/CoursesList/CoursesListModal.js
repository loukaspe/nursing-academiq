import React, {useEffect, useState} from "react";
import "./CoursesListModal.css";

const CoursesListModal = ({onCourseSelect, isOpen}) => {
    const [courses, setCourses] = useState([]);
    const [selectedCourse, setSelectedCourse] = useState("");


    useEffect(() => {
        console.log("anoiksaaa " + isOpen);
        const fetchCourses = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/courses`;

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

                if (result.courses === undefined) {
                    throw Error("error getting courses for modal");
                }
                setCourses(result.courses);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        if (isOpen) {
            fetchCourses();
        }

    }, [isOpen]);

    const handleSelectChange = (e) => {
        setSelectedCourse(e.target.value);
    };

    const handleConfirm = () => {
        onCourseSelect(selectedCourse);
    };

    if (!isOpen) return null;

    return (
        <React.Fragment>
            <div className="coursesListModal">
                <div className="coursesListModalContent">
                    <h2>Επιλέξτε μάθημα</h2>
                    <select value={selectedCourse} onChange={handleSelectChange}>
                        <option value="">-- Επιλέξτε μάθημα --</option>
                        {courses.map((course) => (
                            <option key={course.id} value={course.id}>
                                {course.title}
                            </option>
                        ))}
                    </select>
                    <button onClick={handleConfirm} disabled={!selectedCourse}>
                        Συνέχεια
                    </button>
                </div>
            </div>
        </React.Fragment>
    );
};

export default CoursesListModal;