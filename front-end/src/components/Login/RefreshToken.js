import Cookies from "universal-cookie";

const cookies = new Cookies();

async function refreshToken() {
    try {
        const response = await fetch(process.env.REACT_APP_API_URL + "/refresh-token", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${process.env.REACT_APP_API_KEY}`,
            },
            credentials: 'include',
        });

        if (!response.ok) {
            throw new Error("Failed to refresh token");
        }

        const result = await response.json();
        cookies.set("access_token", result.access_token, { path: "/" });
        return result.access_token;
    } catch (error) {
        console.error("Error refreshing token:", error);
        cookies.remove("access_token", {path: "/"});
        cookies.remove("result", {path: "/"});
        cookies.remove("user", {path: "/"});
        window.location.href = "/login";
    }
}

export default refreshToken;
