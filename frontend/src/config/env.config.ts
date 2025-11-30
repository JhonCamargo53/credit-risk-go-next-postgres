export const JWT_COOKIE_NAME = process.env.NEXT_PUBLIC_JWT_COOKIE_NAME || "management-token";

export const NODE_ENV = (process.env.NEXT_PUBLIC_NODE_ENV || "development") as "development" | "production";

export const DEVELOP_BASE_URL = process.env.NEXT_PUBLIC_DEVELOP_BASE_URL || 'localhost:3000';
export const PRODUCTION_BASE_URL = process.env.NEXT_PUBLIC_PRODUCTION_BASE_URL || 'localhost:5000';
