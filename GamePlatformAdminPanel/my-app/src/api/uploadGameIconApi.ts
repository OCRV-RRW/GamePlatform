import { API_GAMES_PATH } from "../ApiPathes";
import { ApiForm, UploadIconForm } from "../app/api_forms_interfaces";
import fetchAuthAPI from "../app/fetchAPI";
import { BACKEND_DOMAIN } from "../Settings";

const fetch_upload_icon_response = (access_token: string, query_search?: string, body?: ApiForm) : Promise<Response> => {
    let requestBody : FormData = new FormData()
    requestBody.append("icon", (body as UploadIconForm).image)
    return fetch(API_GAMES_PATH + '/' + query_search + '/icon', {method: "POST", body: requestBody, credentials: "include", headers: {
        'Access-Control-Allow-Origin' : BACKEND_DOMAIN,
        'Authorization': access_token ? 'Bearer ' + access_token : ""
    }})
} 

export function fetch_upload_icon(body: UploadIconForm, query_search?: string) : Promise<{access_token: string, response: Response}> {
    return new Promise<{access_token: string, response: Response}>(
        (resolve, reject) => {
            fetchAuthAPI({fetch_func: fetch_upload_icon_response, body: body, query_search})
            .then(
                (fetch_api_data) => resolve({access_token: fetch_api_data.access_token, response: fetch_api_data.response}),
                (reason) => reject(reason)
            )
        }
    )
}