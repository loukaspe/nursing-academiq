import React, {useState} from "react";
import "./CsvImport.css";
import axios from "axios";


const CsvImport = () => {
    const [file, setFile] = useState(null);
    const [message, setMessage] = useState('');

    const onFileChange = (event) => {
        setFile(event.target.files[0]);
    };

    const onFileUpload = async () => {
        if (!file) {
            setMessage('Please select a file first.');
            return;
        }

        const formData = new FormData();
        formData.append('file', file);

        // TODO: take course ID for real
        let courseID = 1;
        let apiUrl = process.env.REACT_APP_API_URL + `/courses/${courseID}/questions/import`

        try {
            const response = await axios.post(apiUrl, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
                }
            });
            console.log('File uploaded successfully.', response);
            setMessage('File uploaded successfully.');
        } catch (error) {
            console.error('Error uploading the file', error);
            setMessage('Error uploading the file.');
        }
    };

    return (
        <div>
            <h2>CSV File Upload</h2>
            <input type="file" accept=".csv" onChange={onFileChange} />
            <button onClick={onFileUpload}>Upload</button>
            {message && <p>{message}</p>}
        </div>
    );
};


export default CsvImport;