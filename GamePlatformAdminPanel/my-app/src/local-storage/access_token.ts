const ACCESS_TOKEN_NAME = "access_token"

export const save_access_token = (token: string) : void => {
    if (localStorage.getItem(ACCESS_TOKEN_NAME)) {
        console.warn("token already exists in local storage")
        drop_access_token()
    }
    localStorage.setItem(ACCESS_TOKEN_NAME, token)
}

export const get_access_token = () : string => {
    let token: string = ""
    if (!localStorage.getItem(ACCESS_TOKEN_NAME)) {
        console.warn("local storage not contain token")
    }
    else {
        token = localStorage.getItem(ACCESS_TOKEN_NAME) ?? ""
    }

    return token
}

export const drop_access_token = () : void => {
    if (!localStorage.getItem(ACCESS_TOKEN_NAME)) {
        console.warn("local storage not contain token")
    }
    else {
        localStorage.removeItem(ACCESS_TOKEN_NAME)
    }
}

export const rewrite_access_token = (token: string) : void => {
    drop_access_token()
    save_access_token(token)
}