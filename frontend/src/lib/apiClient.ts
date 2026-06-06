import axios from 'axios';

// In production, API requests go through CloudFront via the /api path behaviour,
// which proxies to the ALB origin. This avoids mixed-content (HTTPS→HTTP) blocks.
// For local development, set VITE_API_BASE_URL to your local backend (e.g. http://localhost:8080).
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api';
export const apiClient = axios.create({
    baseURL: API_BASE_URL,
    withCredentials: true, // Crucial for httpOnly cookies
});
