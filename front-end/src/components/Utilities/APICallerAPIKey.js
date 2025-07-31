import axios from "axios";
import refreshToken from "../Login/RefreshToken"; // Import the refresh token function
import Cookies from "universal-cookie";
import {jwtDecode} from "jwt-decode";
import {useNavigate} from "react-router-dom";

const cookies = new Cookies();

const apiWithAPIKey = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    withCredentials: true,
});

apiWithAPIKey.interceptors.request.use(
    async (config) => {
        config.headers = {
            ...config.headers,
            Authorization: `Bearer ${process.env.REACT_APP_API_KEY}`,
        };

        if (!config.headers["Content-Type"]) {
            config.headers["Content-Type"] = "application/json";
        }

        return config;
    },
    (error) => Promise.reject(error)
);

export default apiWithAPIKey;
