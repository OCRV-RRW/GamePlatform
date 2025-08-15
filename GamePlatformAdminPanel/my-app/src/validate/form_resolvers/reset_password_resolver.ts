import { FieldErrors, ResolverResult } from "react-hook-form"
import { ResetPasswordForm } from "../../app/api_forms_interfaces"
import PasswordValidate from "../password_validate"
import { INVALID_PATTERN_PASSWORD_ERROR, NOT_SAME_PASSWORDS_ERROR, REQUIRED_FIELD_ERROR } from "../../FormValidErrors"


export default async function ResetPasswordFormResolver(values: ResetPasswordForm): Promise<ResolverResult<ResetPasswordForm>> {
    let errors : FieldErrors<ResetPasswordForm> = {}

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