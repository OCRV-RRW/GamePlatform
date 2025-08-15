export const EMAIL_REG_EXP : RegExp = /\S+@\w+\.\w+/
export const PASSWORD_REG_EXP : RegExp = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\W).{8,}$/
export const QUERY_STRING_GAME_REG_EXP : RegExp = /^\?name=.{1,}$/
export const QUERY_STRING_USER_BY_ID_REG_EXP : RegExp = /^\?id=.{1,}$/
export const GOOD_STATUS_SERVER_RESPONSE_REG_EXP : RegExp = /^[2-3]\d{2}/
export const BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP: RegExp = /^4\d{2}/
export const BAD_STATUS_SERVER_RESPONSE_SERVER_WARNING_REG_EXP: RegExp = /^5\d{2}/