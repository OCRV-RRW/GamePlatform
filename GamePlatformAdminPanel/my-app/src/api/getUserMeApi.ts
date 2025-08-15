import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { User } from "../app/user_type";
import { API_USER_ME_PATH } from "../ApiPathes";

const fetch_user_me_response = (access_token: string) : Promise<Response> => fetch(API_USER_ME_PATH, 
    set_request_options({method: "GET", access_token: access_token})
)

export function fetch_user_me() : Promise<{access_token: string, user_data: User | null}> {
    let response_access_token = ""
    return new Promise<{access_token: string, user_data: User | null}>((resolve, reject) => {
        fetchAuthAPI({fetch_func: fetch_user_me_response})
            .then((fetch_api_data) => {
                response_access_token = fetch_api_data.access_token
                return fetch_api_data.response.json().then((response_data) => {
                    return resolve({access_token: response_access_token, user_data: response_data?.data?.user})}, (reason) => reject(reason))
        })
    })
}