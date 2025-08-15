import { set_request_options } from "../app/set_request_options"
import { API_VERIFY_REGISTER_EMAIL } from "../ApiPathes"
import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp"

export function fetch_verify_email(id: string) : Promise<void> {

    return new Promise<void>(
        (resolve, reject) => {
            fetch(API_VERIFY_REGISTER_EMAIL + id, set_request_options({method: "POST"})
                )
                .then((response) => GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString()) 
                    ? resolve() 
                    : reject(response.status.toString()))
        })
}