export interface User  {
    id: string | undefined
    name: string | undefined
    email: string | undefined
    is_admin: boolean
    birthday: Date | null
    gender: UserGender
}

export type UserGender = "" | "М" | "Ж" 