import { GOOD_STATUS_SERVER_RESPONSE_REG_EXP } from "../reg-exp"
import { get_access_token } from "../local-storage/access_token"
import update_token_middleware from "../middlewares/update_token_middleware"
import { ApiForm } from "./api_forms_interfaces"

export type fetchAPIData = {
    access_token: string, 
    updated: boolean,
    response: Response
}

interface FetchAPIParams {
    fetch_func: (access_token: string, query_search?: string, body?: ApiForm) => Promise<Response>,
    query_search?: string,
    body?: ApiForm
}

export default function fetchAuthAPI(params: FetchAPIParams) 
    : Promise<fetchAPIData> {
    let response_data : fetchAPIData = {
        access_token: "",
        updated: false,
        response: new Response()
    }

    let afterFirstQueryResponse = new Response()
    return new Promise<fetchAPIData>((resolve, reject) => {
        params.fetch_func(get_access_token(), params.query_search, params.body)
            .then((response) => {
                afterFirstQueryResponse = response
                return update_token_middleware(response)})
            .then((updated_token_data) => {
                response_data.access_token = updated_token_data.access_token
                response_data.updated = updated_token_data.updated
                if (!response_data.updated) {
                    response_data.response = afterFirstQueryResponse
                    GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response_data.response.status.toString())
                        ? resolve(response_data) 
                        : reject(response_data.response.status)
                }
                else {
                    return params.fetch_func(response_data.access_token, params.query_search, params.body)
                        .then((response) => {
                            response_data.response = response
                            GOOD_STATUS_SERVER_RESPONSE_REG_EXP.test(response_data.response.status.toString())
                                ? resolve(response_data)
                                : reject(response_data.response.status.toString())
                        })
                }
            })
    })
} 