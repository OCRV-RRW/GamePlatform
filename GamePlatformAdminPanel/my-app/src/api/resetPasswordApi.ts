import { ResetPasswordForm } from "../app/api_forms_interfaces"
import { set_request_options } from "../app/set_request_options"
import { API_RESET_PASWORD_PATH } from "../ApiPathes"
import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp"

export function fetch_reset_password(reset_password_form: ResetPasswordForm, reset_code: string) : Promise<void> {
    return new Promise<void> ((resolve, reject) => {
        fetch(API_RESET_PASWORD_PATH + reset_code, 
                set_request_options({method: "PATCH", body_form: reset_password_form})
            )
            .then(
                (response) => GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString()) 
                ? resolve() 
                : reject(response.status.toString()))
    })
}