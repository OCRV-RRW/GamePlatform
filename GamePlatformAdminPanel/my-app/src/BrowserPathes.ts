//ROOT
export const HOME_PATH: string = "/"

//AUTH
export const LOGIN_PATH: string = "/login"
export const REGISTER_PATH: string = "/register"
export const FORGOT_PASSWORD_PATH: string = "/forgot_password"
export const LOGOUT_PATH: string = "/logout"
export const REGISTER_VERIFY_EMAIL_PATH = REGISTER_PATH + '/verify/'
export const RESET_PASSWORD_PATH = FORGOT_PASSWORD_PATH + '/reset/'

//GAME
export const GAME_PATH: string = '/game'


//ADMIN
export const ADMIN_PANEL_PATH: string = "/admin-panel"
export const GAMES_LIST_PATH: string = ADMIN_PANEL_PATH + "/games-list"
export const UPDATE_GAME_PATH: string = ADMIN_PANEL_PATH + "/update-game"
export const SKILLS_LIST_PATH: string = ADMIN_PANEL_PATH + "/skills-list"
export const UPDATE_SKILL_PATH: string = ADMIN_PANEL_PATH + "/update-skill"
export const UPDATE_USER_PATH: string = ADMIN_PANEL_PATH + "/update-user"