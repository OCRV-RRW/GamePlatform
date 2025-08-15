import { API_REGISTER_PATH } from "../ApiPathes"
import { RegisterForm } from "../app/api_forms_interfaces"
import { set_request_options } from "../app/set_request_options"
import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp"

export function fetch_register(register_form: RegisterForm) : Promise<void> {
    return new Promise<void>(
        (resolve, reject: (reason: string) => void) => {
            fetch(API_REGISTER_PATH, set_request_options({method: "POST", body_form: register_form}))
                .then((response) => GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString()) 
                ? resolve() 
                : reject(response.status.toString()))
        }
    )
}