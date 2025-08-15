import { ApiForm, UpdateGameForm } from "../app/api_forms_interfaces";
import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { API_UPDATE_GAME_PATH } from "../ApiPathes";

const fetch_update_game_response = (access_token: string, query_search?: string, body?: ApiForm) : Promise<Response> => 
    fetch(API_UPDATE_GAME_PATH + query_search, set_request_options({method: "PATCH", access_token: access_token, body_form: <UpdateGameForm>body})
)

export function fetch_update_game(body: UpdateGameForm, query_search: string) : Promise<{access_token: string, response: Response}> {
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_update_game_response, body: body, query_search})
            .then(
                (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                (reason) => reject(reason)
            )
        }
    )
}