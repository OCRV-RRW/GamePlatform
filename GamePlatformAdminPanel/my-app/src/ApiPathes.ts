import { BACKEND_DOMAIN } from "./Settings"

//ROOT
export const API_PATH : string = BACKEND_DOMAIN + "/api/v1"

//AUTH
export const API_AUTH_PATH : string = API_PATH + '/auth'
export const API_REGISTER_PATH : string = API_AUTH_PATH + '/register'
export const API_LOGIN_PATH : string = API_AUTH_PATH + '/login'
export const API_LOGOUT_PATH : string = API_AUTH_PATH + '/logout'
export const API_FORGOT_PASSWORD_PATH : string = API_AUTH_PATH + '/forgot-password'
export const API_RESET_PASWORD_PATH : string = API_AUTH_PATH + '/reset-password/'
export const API_REFRESH_PATH : string = API_AUTH_PATH + '/refresh'
export const API_VERIFY_REGISTER_EMAIL : string =  API_AUTH_PATH + '/verify-email/'

//USER
export const API_USER_PATH : string = API_PATH + '/users'
export const API_USER_ME_PATH : string = API_USER_PATH + '/me'
export const API_GET_USERS_PATH : string = API_USER_PATH

//GAME
export const API_GAMES_PATH : string = API_PATH + '/games'
export const API_GET_GAME_PATH : string = API_GAMES_PATH + '/'

//SKILL 
export const API_SKILLS_PATH : string = API_PATH + "/skills"

//USER ME SKILLS
export const API_USER_ME_SKILLS_PATH : string = API_USER_ME_PATH + "/skills"
export const API_GET_USER_ME_SKILLS_PATH : string = API_USER_ME_SKILLS_PATH
export const API_ADD_SCORE_USER_ME_SKILLS_PATH : string = API_USER_ME_SKILLS_PATH

//ADMIN/GAMES
export const API_UPDATE_GAME_PATH : string = API_GAMES_PATH + '/'
export const API_CREATE_GAME_PATH : string = API_GAMES_PATH + '/'
export const API_DELETE_GAME_PATH : string = API_GAMES_PATH + '/'
export const API_CREATE_PREVIEW_PATH : string = API_GAMES_PATH + '/previews'
export const API_DELETE_PREVIEW_PATH : string = API_GAMES_PATH + '/previews/'

//ADMIN/SKILLS
export const API_CREATE_SKILL_PATH : string = API_SKILLS_PATH + '/'
export const API_DELETE_SKILL_PATH : string = API_SKILLS_PATH + '/'
export const API_UPDATE_SKILL_PATH : string = API_SKILLS_PATH + '/'
export const API_GET_SKILL_PATH : string = API_SKILLS_PATH

//ADMIN/USER
export const API_DELETE_USER_PATH: string = API_USER_PATH + '/'
export const API_UPDATE_USER_PATH: string = API_USER_PATH + '/'
export const API_GET_USER_PATH: string = API_USER_PATH