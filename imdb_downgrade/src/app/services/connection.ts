const HOST = `localhost`;
const PORT = 5050;
const PREFIX = `http://${HOST}:${PORT}`;

// Public
export const REQUEST_GOOGLE_LOGIN = `${PREFIX}/api/v1/public/auth/google/login`;
export const REQUEST_LOGIN = `${PREFIX}/api/v1/public/auth/login`;
export const REQUEST_REGISTER = `${PREFIX}/api/v1/public/auth/register`;
export const REQUEST_LOGOUT = `${PREFIX}/api/v1/public/auth/logout`;
export const REQUEST_ALL_MOVIES = `${PREFIX}/api/v1/public/movie/all`;
export const REQUEST_ALL_TV_SHOW = `${PREFIX}/api/v1/public/tv_show/all`;
export const REQUEST_USER_INFO = `${PREFIX}/api/v1/public/auth/user`;
export const REQUEST_CITIES_BY_COUNTRY = (country: string) => `${PREFIX}/api/v1/public/city/by-country?country=${country}`;
export const REQUEST_MOVIE_DETAILS = (id: string) => `${PREFIX}/api/v1/public/movie?id=${id}`;
export const REQUEST_TV_SHOW_DETAILS = (id: string) => `${PREFIX}/api/v1/public/tv_show?id=${id}`;
export const REQUEST_SEARCH = (keyword: string) => `${PREFIX}/api/v1/public/search?keyword=${keyword}`;
export const REQUEST_CELEBRITY_DETAILS = (id: string) => `${PREFIX}/api/v1/public/celebrity?id=${id}`;
export const REQUEST_ALL_COUNTRIES = `${PREFIX}/api/v1/public/countries/all`;
// Private
export const REQUEST_WATCH_LIST = `${PREFIX}/api/v1/private/watchlist`;
export const ADD_WATCH_LIST_MOVIE = (id: string) => `${PREFIX}/api/v1/private/movie/watchlist/add?movie_id=${id}`;
export const ADD_WATCH_LIST_TV = (id: string) => `${PREFIX}/api/v1/private/tv_show/watchlist?tv_show_id=${id}`;
export const REMOVE_WATCH_LIST_MOVIE = (id: string) => `${PREFIX}/api/v1/private/movie/watchlist/delete?movie_id=${id}`;
export const REMOVE_WATCH_LIST_TV = (id: string) => `${PREFIX}/api/v1/private/tv_show/watchlist?tv_show_id=${id}`;
export const REQUEST_RATE_MOVIE = (id: string, rate: number) => `${PREFIX}/api/v1/private/movie/rate?movie_id=${id}&rate=${rate}`;
export const REQUEST_RATE_TV_SHOW = (id: string, rate: number) => `${PREFIX}/api/v1/private/tv_show/rate?tv_show_id=${id}&rate=${rate}`;
