import React, { useState } from 'react';
import axios from 'axios';

const FileUpload = ({text}) => {
    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
    };

    const handleUpload = () => {
        const formData = new FormData();
        formData.append('image', selectedFile);

        axios.post('http://localhost:8080/test', formData)
            .then(response => {
                console.log('Upload successful');
            })
            .catch(error => {
                console.error('Error uploading file', error);
            });
    };

    // Implement user profile picture retrieval using axios.get for the second endpoint

    return (
        <div>
            <input type="file" onChange={handleFileChange} />
            <button onClick={handleUpload}>Upload</button>
            {/* Render user interface to retrieve profile picture */}
        </div>
    );
};

export default FileUpload;