import { User } from "./user_type"

export type PageStatus = "ok" | "400" | "401" | "402" | "403" | "404" | "422" | "500" | "501" | "502" | "503"

export interface PageState {
    current_page: string,
    status: PageStatus
}

export interface UserState {
    user_data: User | null
    access_token: string
    expired_in?: string
}
