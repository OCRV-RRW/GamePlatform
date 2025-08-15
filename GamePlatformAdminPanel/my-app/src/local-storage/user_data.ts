import { User } from "../app/user_type"

const USER_DATA_NAME = "user_data"

export const save_userdata = (data: User | null) : void => {
    if (localStorage.getItem(USER_DATA_NAME)) {
        console.warn("userdata already exists in local storage")
        drop_userdata()
    }
    if (data)
        localStorage.setItem(USER_DATA_NAME, JSON.stringify(data))
}

export const get_userdata = () : User => {
    let userdata_json : User | null = null
    if (!localStorage.getItem(USER_DATA_NAME)) {
        console.warn("local storage not contain userdata")
    }
    else {
        userdata_json = JSON.parse(localStorage.getItem(USER_DATA_NAME) ?? "")
    }

    return {
        id: userdata_json?.id, 
        name: userdata_json?.name, 
        email: userdata_json?.email,
        is_admin: userdata_json?.is_admin || false,
        birthday: null,
        gender: "",
    }
}

export const drop_userdata = () : void => {
    if (!localStorage.getItem(USER_DATA_NAME)) {
        console.warn("local storage not contain userdata")
    }
    else {
        localStorage.removeItem(USER_DATA_NAME)
    }
}


export const rewrite_user_data = (data: User | null) : void => {
    drop_userdata()
    save_userdata(data)
}