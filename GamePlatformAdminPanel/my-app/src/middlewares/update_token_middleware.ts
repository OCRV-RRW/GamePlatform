import { set_request_options } from "../app/set_request_options"
import { API_REFRESH_PATH } from "../ApiPathes"
import { FORBIDDEN } from "../ResponseCodes"
import { get_access_token } from "../local-storage/access_token"

export default function update_token_middleware(response: Response) : Promise<{access_token: string, updated: boolean}> {
    return new Promise<{access_token: string, updated: boolean}>(
        (resolve) => {
            if (response.status === 401 || response.status === 403) {
                fetch(API_REFRESH_PATH, set_request_options({method: "POST"}))
                    .then((response) => 
                        refresh_token_middleware(response.status, response.json()))
                    .then((data) => resolve({access_token: data.data.access_token, updated: true}), 
                        () => resolve({access_token: "", updated: true}))
            }
            else {
                resolve({access_token: get_access_token(), updated: false})
            }
    })
}

export function refresh_token_middleware(status: number, json: Promise<any>) : Promise<any> {
    return new Promise<any>(
        (resolve, reject) => {
            if (status !== FORBIDDEN) {
                return resolve(json)
            } 
            else {
                return reject(FORBIDDEN)
            }
        }
    )
}