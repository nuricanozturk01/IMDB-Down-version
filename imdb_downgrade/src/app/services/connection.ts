const HOST = `localhost`;
const PORT = 5050;
const PREFIX = `http://${HOST}:${PORT}`;

const AUTH_PREFIX = `${PREFIX}/api/auth`;
const V1_PREFIX = `${PREFIX}/api/v1`;

export const REQUEST_GOOGLE_AUTH = `${AUTH_PREFIX}/google/login`;
export const REQUEST_LOGIN = `${AUTH_PREFIX}/login`;
export const REQUEST_REGISTER = `${AUTH_PREFIX}/register`;
export const REQUEST_ALL_MOVIES = `${V1_PREFIX}/movie/all`;
export const REQUEST_WATCH_LIST = `${V1_PREFIX}/watchlist`;
export const ADD_WATCH_LIST_MOVIE = (id: string) => `${V1_PREFIX}/movie/watchlist/add?movie_id=${id}`;
export const LIKE_MOVIE = (id: string) => `${V1_PREFIX}/movie/like?movie_id=${id}`;
export const ADD_WATCH_LIST_TV = (id: string) => `${V1_PREFIX}/tv_show/watchlist/add?tv_show_id=${id}`;
export const REMOVE_WATCH_LIST_MOVIE = (id: string) => `${V1_PREFIX}/movie/watchlist/delete?movie_id=${id}`;
export const REMOVE_WATCH_LIST_TV = (id: string) => `${V1_PREFIX}/tv_show/watchlist/delete?movie_id=${id}`;
export const REQUEST_SEARCH = (keyword: string) => `${V1_PREFIX}/search?keyword=${keyword}`;
export const REQUEST_MOVIE_DETAILS = (id: string) => `${V1_PREFIX}/movie?id=${id}`;