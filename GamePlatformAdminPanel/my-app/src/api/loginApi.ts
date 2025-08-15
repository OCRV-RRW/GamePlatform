import { API_USER_ME_PATH, API_LOGIN_PATH } from "../ApiPathes";
import { LoginForm } from "../app/api_forms_interfaces";
import { User } from "../app/user_type";
import { set_request_options } from "../app/set_request_options";
import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp";

export function fetch_log_in(login_form: LoginForm) {
    let access_token: string = ""
    let expired_in: string = ""
    
    return new Promise<{ access_token: string, user_data: User, expired_in: string }>(
        (resolve, reject: (reason: string) => void) => {
            fetch(API_LOGIN_PATH, set_request_options({method: "POST", body_form: login_form}))
            .then((response) => {
                if (GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString())) {
                    return response.json().then((response_data) => {
                        console.log(response_data)
                        access_token = response_data?.data?.access_token
                        expired_in = response_data?.data?.expired_in
                        return get_user_data_by_token(access_token)
                    })
                    .then((user_data) => resolve( { access_token: access_token, user_data: user_data.user_data, expired_in: expired_in }),
                        (reason) => reject(reason))
                }
                else {
                    return reject(response.status.toString())
                } 
            })
        }
    );
}

export function get_user_data_by_token(access_token: string) : Promise<{ user_data: User }> {
    return new Promise<{ user_data: User}>(
        (resolve, reject) => {
            fetch(API_USER_ME_PATH, set_request_options({method: "GET", access_token: access_token}))
            .then((response) => {
                if (GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response.status.toString())) {
                    return response.json().then((response_data) => resolve( { user_data: response_data?.data?.user } ))
                }
                else {
                    return reject(response.status.toString())
                }
            })
        }
    )
}