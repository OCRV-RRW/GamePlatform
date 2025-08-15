import { ApiForm, UpdateUserForm } from "../app/api_forms_interfaces";
import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { API_UPDATE_USER_PATH } from "../ApiPathes";

const fetch_update_user_response = (access_token: string, query_search?: string, body?: ApiForm) : Promise<Response> => 
    fetch(API_UPDATE_USER_PATH + query_search, set_request_options({method: "PATCH", access_token: access_token, body_form: <UpdateUserForm>body})
)

export function fetch_update_user(body: UpdateUserForm, query_search: string) : Promise<{access_token: string, response: Response}> {
    console.log(body)
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_update_user_response, body: body, query_search})
            .then(
                (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                (reason) => reject(reason)
            )
        }
    )
}