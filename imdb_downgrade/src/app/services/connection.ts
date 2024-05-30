const HOST = `localhost`;
const PORT = 5050;
const PREFIX = `http://${HOST}:${PORT}`;

const AUTH_PREFIX = `${PREFIX}/api/auth`;
const V1_PREFIX = `${PREFIX}/api/v1`;

export const REQUEST_GOOGLE_AUTH = `${AUTH_PREFIX}/google/login`;
export const REQUEST_LOGIN = `${AUTH_PREFIX}/login`;
export const REQUEST_REGISTER = `${AUTH_PREFIX}/register`;
export const REQUEST_SEARCH = (keyword: string) => `${V1_PREFIX}/search?keyword=${keyword}`;
