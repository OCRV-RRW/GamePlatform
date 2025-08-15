import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { API_GET_USERS_PATH } from "../ApiPathes";

const fetch_get_users_response = (access_token: string, query_search?: string) : Promise<Response> => 
    fetch(API_GET_USERS_PATH, set_request_options({method: "GET", access_token: access_token})
)

export function fetch_get_users() : Promise<{access_token: string, response: Response}> {
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_get_users_response})
                .then(
                    (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                    (reason) => reject(reason)
                )
        }
    )
}