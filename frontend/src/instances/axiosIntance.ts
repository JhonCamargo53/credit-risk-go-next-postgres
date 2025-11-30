import { DEVELOP_BASE_URL, JWT_COOKIE_NAME, NODE_ENV, PRODUCTION_BASE_URL } from "@/config/env.config";
import { getCookieValueService } from "@/utils/cookieUtils";
import axios from "axios";

export const BASE_URL = (NODE_ENV == 'production' ? PRODUCTION_BASE_URL : DEVELOP_BASE_URL);

export const axiosInstance = axios.create({
    baseURL: BASE_URL,
});

axiosInstance.interceptors.request.use(
    async function (config) {
        const token = getCookieValueService(JWT_COOKIE_NAME);
        config.headers.Authorization = token ? `Bearer ${token}` : '';
        return config;
    },
    function (error) {
        return Promise.reject(error);
    }
);