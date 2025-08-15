import { FieldError } from "react-hook-form"

export const REQUIRED_FIELD_ERROR : FieldError | undefined = {
    type: 'pattern',
    message: 'это поле должно быть заполнено'
}

export const INVALID_PATTERN_PASSWORD_ERROR : FieldError | undefined = {
    type: 'pattern',
    message: 'пароль должен состоять не менее чем из 8 символом, содержать в себе цифры, специальные знаки, заглавные и маленькие буквы'
}

export const INVALID_PATTERN_EMAIL_ERROR : FieldError | undefined = {
    type: 'pattern',
    message: 'неккоректная электронная почта'
}

export const NOT_SAME_PASSWORDS_ERROR : FieldError | undefined = {
    type: 'validate',
    message: 'пароли должны совпадать'
}