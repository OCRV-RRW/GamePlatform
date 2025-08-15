import { FieldErrors, ResolverResult } from "react-hook-form";
import { LoginForm } from "../../app/api_forms_interfaces";
import { REQUIRED_FIELD_ERROR } from "../../FormValidErrors";

export default async function LoginFormResolver(values: LoginForm) : Promise<ResolverResult<LoginForm>> {
    let errors : FieldErrors<LoginForm> = {}

    if (values.email === "") {
        errors.email = REQUIRED_FIELD_ERROR
    }

    if (values.password === "") {
        errors.password = REQUIRED_FIELD_ERROR
    }

    return {
        values: values,
        errors: errors
    }
}