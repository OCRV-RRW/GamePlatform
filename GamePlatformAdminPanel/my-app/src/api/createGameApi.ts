import { ApiForm, CreateGameForm } from "../app/api_forms_interfaces";
import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { API_CREATE_GAME_PATH } from "../ApiPathes";

const fetch_create_game_response = (access_token: string, query_search?: string, body?: ApiForm) : Promise<Response> =>  {
    return fetch(API_CREATE_GAME_PATH, set_request_options({method: 'POST', access_token: access_token, body_form: body}))
}

export function fetch_create_game(body: CreateGameForm) : Promise<{access_token: string, response: Response}> {
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_create_game_response, body: body!})
                .then(
                    (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                    (reason) => reject(reason)
                )
        }
    )
}