import { ForgotPasswordForm } from "../app/api_forms_interfaces"
import { set_request_options } from "../app/set_request_options"
import { API_FORGOT_PASSWORD_PATH } from "../ApiPathes"
import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp"

export function fetch_forgot_password(forgot_password_form: ForgotPasswordForm) : Promise<void> {
    return new Promise<void> ((resolve, reject) => {
        fetch(API_FORGOT_PASSWORD_PATH, set_request_options({method: "POST", body_form: forgot_password_form}))
            .then((response) => GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString()) 
                ? resolve() 
                : reject(response.status.toString()))
    })
}