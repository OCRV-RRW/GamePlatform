import { ApiForm, CreatePreviewForm } from "../app/api_forms_interfaces";
import fetchAuthAPI from "../app/fetchAPI";
import { set_request_options } from "../app/set_request_options";
import { API_CREATE_PREVIEW_PATH } from "../ApiPathes";
import { BACKEND_DOMAIN } from "../Settings";

const fetch_create_preview_response = (access_token: string, query_search?: string, body?: ApiForm) : Promise<Response> => {
    let requestBody : FormData = new FormData()
    requestBody.append("gameId", (body as CreatePreviewForm).game_id)
    requestBody.append("image", (body as CreatePreviewForm).image)
    if ((body as CreatePreviewForm).video)
        requestBody.append("video", (body as CreatePreviewForm).video!)
    return fetch(API_CREATE_PREVIEW_PATH, {method: "POST", body: requestBody, credentials: "include", headers: {
        'Access-Control-Allow-Origin' : BACKEND_DOMAIN,
        'Authorization': access_token ? 'Bearer ' + access_token : ""
    }})
} 

export function fetch_create_preview(body: CreatePreviewForm, query_search?: string) : Promise<{access_token: string, response: Response}> {
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_create_preview_response, body: body, query_search})
            .then(
                (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                (reason) => reject(reason)
            )
        }
    )
}