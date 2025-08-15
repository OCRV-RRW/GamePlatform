export interface ApiForm {}

export interface LoginForm extends ApiForm {
    email: string
    password: string
}

export interface RegisterForm extends ApiForm {
    name: string
    email: string
    password: string
    password_confirm: string
}

export interface ForgotPasswordForm extends ApiForm {
    email: string
}

export interface ResetPasswordForm extends ApiForm {
    password: string
    password_confirm: string
}

// export interface AddScoreToSkillForm extends ApiForm {
//     score: number,
//     skill_name: string
// }

export interface UpdateGameForm extends ApiForm {
    description: string,
    icon: string,
    src: string,
    title: string
}

export interface CreateGameForm extends ApiForm {
    description: string,
    icon: string,
    src: string
    title: string
}

// export interface CreateSkillForm extends ApiForm {
//     description: string,
//     friendly_name: string,
//     name: string
// }

export interface UpdateUserForm extends ApiForm {
    birthday: Date,
    gender: string,
    is_admin: boolean,
    name: string
}

// export interface UpdateSkillForm extends ApiForm {
//     description: string,
//     friendly_name: string
// }

export interface UploadIconForm extends ApiForm {
    image: File
}  

export interface CreatePreviewForm extends ApiForm {
    game_id: string,
    image: File,
    video?: File
}