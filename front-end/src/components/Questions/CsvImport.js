import React, {useEffect, useState} from "react";
import "./CsvImport.css";
import axios from "axios";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDownload, faPenToSquare, faUpload} from "@fortawesome/free-solid-svg-icons";
import Breadcrumb from "../Utilities/Breadcrumb";
import {useNavigate, useParams} from "react-router-dom";
import api from "../Utilities/APICaller";

const CsvImport = () => {
    const [file, setFile] = useState(null);
    const [message, setMessage] = useState('');
    const [courseTitle, setCourseTitle] = useState('');
    const [fileName, setFileName] = useState("");
    const [createNewChaptersOption, setCreateNewChaptersOption] = useState(true);
    const [exportMistakesOption, setExportMistakesOption] = useState(false);

    let navigate = useNavigate();

    const params = useParams();
    let courseID = params.id;

    useEffect(() => {
        const fetchCourse = async () => {
            let apiUrl = process.env.REACT_APP_API_URL + `/course/${courseID}`

            try {
                const response = await axios.get(apiUrl, {
                    headers: {
                        Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                    },
                });
                setCourseTitle(response.data.course.Course.title);
            } catch (error) {
                console.error('Error fetching the course data', error);
            }
        };

        fetchCourse();
    }, []);

    const onFileChange = (event) => {
        setFile(event.target.files[0]);
        if (event.target.files.length > 0) {
            setFileName(event.target.files[0].name);
        } else {
            setFileName("");
        }
    };

    const handleCreateNewChaptersChange = (event) => {
        setCreateNewChaptersOption(event.target.checked);
    };

    const handleExportMistakesChange = (event) => {
        setExportMistakesOption(event.target.checked);
    };

    const onFileUpload = async () => {
        if (!file) {
            setMessage('Please select a file first.');
            return;
        }

        const formData = new FormData();
        const jsonData = {create_new_chapters: createNewChaptersOption};
        formData.append("jsonData", JSON.stringify(jsonData));
        formData.append('file', file);

        // TODO: take course ID for real
        let courseID = 1;
        let apiUrl = `/courses/${courseID}/questions/import`

        try {
            const response = await api.post(apiUrl, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                }
            });

            if (response.status === 200) {
                var contentType = response.headers.get("content-type")
                if (contentType === "text/csv" && exportMistakesOption) {
                    const blob = new Blob([response.data], {type: contentType});
                    const url = window.URL.createObjectURL(blob);
                    const a = document.createElement("a");
                    a.href = url;
                    a.download = "problematic_records.csv";
                    document.body.appendChild(a);
                    a.click();
                    a.remove();
                    window.URL.revokeObjectURL(url);
                    setMessage('Η εισαγωγή ολοκληρώθηκε μερικώς. Δείτε τις λανθασμένες ερωτήσεις στο αρχείο.');
                } else {
                    setMessage('Η εισαγωγή ολοκληρώθηκε μερικώς. Oι λανθασμένες απαντήσεις αγνοήθηκαν.');
                }
            } else if (response.status === 204) {
                setMessage('Η εισαγωγή ολοκληρώθηκε επιτυχώς.');
            }
            console.log('File uploaded successfully.', response);
        } catch (error) {
            console.error('Error uploading the file', error);
            setMessage('Η εισαγωγή απέτυχε. Παρακαλώ δοκιμάστε ξανά.');
        }
    };

    const downloadExample = () => {
        const csvRows = [
            ["Εκφώνηση", "Κατηγορία", "Πλήθος Απαντήσεων", "Επεξήγηση Λύσης", "Πηγή", "Απάντηση 1", "Ορθότητα 1", "Απάντηση 2", "Ορθότητα 2", "Απάντηση 3", "Ορθότητα 3", "Απάντηση 4", "Ορθότητα 4", "Απάντηση 5", "Ορθότητα 5", "Απάντηση 6", "Ορθότητα 6", "Απάντηση 7", "Ορθότητα 7", "Απάντηση 8", "Ορθότητα 8", "Εικόνα(Σύνδεσμος / Αρχείο)"],
            ["Ποιες από τις παρακάτω κατηγορίες φαρμάκων έχουν αντιυπερτασική δράση;", "Αντιυπερτασικά", 5, "Oι Β- Αδρενεργικοί ανταγωνιστές & τα Διουρητικά έχουν αντιυπερτασική δράση.", "Βασική & Κλινική φαρμακολογία Bertram G. Katzung /  σελ 147", "Β- Αδρενεργικοί ανταγωνιστές", 1, "Β- Αδρενεργικοί αγωνιστές", 0, "Διουρητικά", 1, "Αμινογλυκοσίδες", 0, "Κατεχολαμίνες", 0, "", "", "", "", "", "", ""],
        ];

        const csvContent = csvRows.map(row => row.join(",")).join("\n");
        const blob = new Blob([csvContent], {type: 'text/csv'});
        const url = URL.createObjectURL(blob);

        const a = document.createElement('a');
        a.href = url;
        a.download = 'example.csv';
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
    };

    return (
        <div>
            <Breadcrumb
                actualPath={`/courses/${courseID}/questions/import`}
                namePath={`/Διαχείριση Μαθημάτων/${courseTitle}/Ερωτήσεις/Εισαγωγή`}
            />
            <div className="csvImportPageContainer">
                <div className="csvImportPageHeader">
                    <div className="csvImportPageInfo">
                        <span className="singleChapterQuizzesPageChapterName">Εισαγωγή Ερωτήσεων</span>
                        <button className="backButton" onClick={() => navigate(-1)}>Πίσω</button>
                    </div>
                </div>
                <div className="csvImportPageImportRow">
                    <div className="csvImportPageImportBox">
                        <div className="csvImportPageText">
                            <FontAwesomeIcon icon={faUpload} className="chapterIcon"/> Επιλέξτε ή σύρετε ένα αρχείο.
                        </div>
                        {fileName && <p className="csvImportPageText">{fileName}</p>}
                        <div>
                            <label htmlFor="importFile" className="csvImportPageChooseFileButton">
                                Επιλογή Αρχείου
                            </label>
                            <div>
                                <input id="importFile" type="file" accept=".csv" onChange={onFileChange}
                                       style={{visibility: "hidden"}}/>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="csvImportPageMessageRow">
                    <div className="csvImportPageMessageText">
                        {message && <p>{message}</p>}
                    </div>
                </div>
                <div className="csvImportPageOptionsRow">
                    <div className="csvImportPageOptionsColumn">
                        <div className="csvImportPageOptions">
                            <label className="csvImportPageText">
                                Δημιουργία Νέων Θεματικών <input type="checkbox"
                                                                 onChange={handleCreateNewChaptersChange}
                                                                 checked={createNewChaptersOption}/>
                            </label>
                            <label className="csvImportPageText">
                                Εξαγωγή Λαθών σε Αρχείο <input type="checkbox"
                                                               onChange={handleExportMistakesChange}
                                                               checked={exportMistakesOption}/>
                            </label>
                        </div>
                    </div>
                    <div className="csvImportPageOptionsColumn">
                        <button className="csvImportPageButton" onClick={downloadExample}>
                            <FontAwesomeIcon icon={faDownload} className="csvImportPageFa"/> Λήψη Προτύπου
                        </button>
                    </div>
                    <div className="csvImportPageOptionsColumn">
                        <button className="csvImportPageButton" onClick={onFileUpload}>Ολοκλήρωση</button>
                    </div>
                </div>
            </div>
        </div>
    );
};


export default CsvImport;