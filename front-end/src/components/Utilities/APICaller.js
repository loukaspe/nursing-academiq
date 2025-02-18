import axios from "axios";
import refreshToken from "../Login/RefreshToken"; // Import the refresh token function
import Cookies from "universal-cookie";
import {jwtDecode} from "jwt-decode";
import {useNavigate} from "react-router-dom";

const cookies = new Cookies();

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    withCredentials: true,
});

api.interceptors.request.use(
    async (config) => {
        let token = cookies.get("access_token");
        // let refresh_token = cookies.get("refreshToken");
        //
        // if (!refresh_token) {
        //     cookies.remove("access_token", {path: "/"});
        //     cookies.remove("result", {path: "/"});
        //     cookies.remove("user", {path: "/"});
        //     window.location.href = "/login";
        // }
        //
        // const refreshTokenExp = jwtDecode(refresh_token).exp * 1000; // Convert to ms
        // if (Date.now() >= refreshTokenExp) {
        //     console.log("mphkee")
        //     cookies.remove("access_token", {path: "/"});
        //     cookies.remove("result", {path: "/"});
        //     cookies.remove("user", {path: "/"});
        //     // window.location.href = "/login";
        // }

        const tokenExp = jwtDecode(token).exp * 1000; // Convert to ms
        try {
            if (Date.now() >= tokenExp) {
                token = await refreshToken();
            }
        } catch (error) {
            console.error("Invalid refresh token:", error);
        }

        config.headers = {
            ...config.headers,
            Authorization: `Bearer ${token}`,
        };

        if (!config.headers["Content-Type"]) {
            config.headers["Content-Type"] = "application/json";
        }

        return config;
    },
    (error) => Promise.reject(error)
);

export default api;
