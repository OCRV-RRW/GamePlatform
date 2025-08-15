import { FieldErrors, ResolverResult } from "react-hook-form"
import { RegisterForm } from "../../app/api_forms_interfaces"
import EmailValidate from "../email_validate"
import PasswordValidate from "../password_validate"
import { INVALID_PATTERN_EMAIL_ERROR, INVALID_PATTERN_PASSWORD_ERROR, NOT_SAME_PASSWORDS_ERROR, REQUIRED_FIELD_ERROR } from "../../FormValidErrors"


export default async function RegisterFormResolver(values: RegisterForm): Promise<ResolverResult<RegisterForm>> {
    let errors : FieldErrors<RegisterForm> = {}

    if (values.name === "") {
        errors.name = REQUIRED_FIELD_ERROR
    }

    if (values.email === "") {
        errors.email = REQUIRED_FIELD_ERROR
    }
    else if (!EmailValidate(values.email)) {
        errors.email = INVALID_PATTERN_EMAIL_ERROR
    }

    if (values.password === "") {
        errors.password = REQUIRED_FIELD_ERROR
    }
    else if (!PasswordValidate(values.password)) {
        errors.password = INVALID_PATTERN_PASSWORD_ERROR
    }
    
    if (values.password_confirm === "") {
        errors.password_confirm = REQUIRED_FIELD_ERROR
    }
    else if (values.password_confirm !== values.password) {
        errors.password_confirm = NOT_SAME_PASSWORDS_ERROR
    }

    return {
        values: values,
        errors: errors
    }
}