import axios from "axios";

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
