import { set_request_options } from "../app/set_request_options";
import { API_REFRESH_PATH } from "../ApiPathes";
import { refresh_token_middleware } from "../middlewares/update_token_middleware";

export default function fetch_refresh(signal?: AbortSignal) : Promise<{access_token: string, expired_in: string}> {
    return new Promise<{ access_token: string, expired_in: string}>((resolve) => {
        fetch(API_REFRESH_PATH, {signal: signal, ...set_request_options({method: "POST"})})
            .then((response) => 
                refresh_token_middleware(response.status, response.json()))
            .then((data) => resolve({access_token: data.data.access_token, expired_in: data.data.expired_in}), 
                () => resolve({access_token: "", expired_in: ""}))
    })
}