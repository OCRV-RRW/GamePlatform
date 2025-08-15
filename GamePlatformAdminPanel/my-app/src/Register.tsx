import { useEffect, useState } from "react"
import { fetch_register } from "./api/registerApi"
import { RegisterForm } from "./app/api_forms_interfaces"
import ReturnToLoginButton from "./ReturnToLoginButton"
import { FieldErrors, useForm } from "react-hook-form"
import { EMAIL_REG_EXP, PASSWORD_REG_EXP } from "./reg-exp"
import RegisterFormResolver from "./validate/form_resolvers/register_resolver"
import register_styles from './css_modules/style.module.css'
import { grey, red } from "@mui/material/colors"
import { Button, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from "@mui/material"
import { BAD_REQUEST } from "./ResponseCodes"
import { Visibility, VisibilityOff } from '@mui/icons-material';

export default function Register() {
    const [registrationState, setRegistrationState] = useState<string>("")
    const [loading, setLoading] = useState<boolean>(false)
    const [showPassword, setShowPassword] = useState<boolean>(false);

    const handleClickShowPassword = () => setShowPassword((show) => !show);

    const checkErrors = (errors: FieldErrors<RegisterForm>) : boolean => {
        return ((errors.email !== undefined) || 
            (errors.name !== undefined) ||
            (errors.password !== undefined) || 
            (errors.password_confirm !== undefined))
    }

    const {register, handleSubmit, formState: { errors }, trigger, reset } = useForm<RegisterForm>( 
        { 
            resolver: RegisterFormResolver, 
            mode: 'onChange',
            defaultValues: { name: "", email: "", password: "", password_confirm: "" }
        } )

    const onRegister = (data: RegisterForm) => {
        setLoading(true)
        fetch_register(data)
            .then(
                () => {
                    setRegistrationState("Перейди на почту")
                    setLoading(false)
                }, 
                (reason) => {
                    if (reason === BAD_REQUEST.toString()) {
                        setRegistrationState("Проверь правильность почты")
                    }
                    setLoading(false)
                }
            )
        reset()
    }

    return(
        <>
         <div style={{display: 'flex', flexDirection: 'column',  alignItems: 'center'}}>
            <form onSubmit={handleSubmit((data) => onRegister(data))}>
                <h1 style={{color: grey[900]}}>Регистрация</h1>
                <h4 style={{color: grey[900]}}>{registrationState}</h4>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="name" className={register_styles.required}>Имя пользователя:</InputLabel>
                    <div style={{margin: 10}}>
                        <OutlinedInput 
                            onFocus={() => setRegistrationState("")}
                            sx={{width: "100%"}} 
                            id="name" {...register('name')} 
                            placeholder="имя пользователя..." 
                            error={errors.name !== undefined} />
                        {errors.name && <h5 style={{color: red[400], margin: 2, fontSize: 10}}>{errors.name.message}</h5>}
                    </div>
                </div>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="email" className={register_styles.required}>Электронная почта:</InputLabel>
                    <div style={{margin: 10}}>
                        <OutlinedInput 
                            onFocus={() => setRegistrationState("")}
                            sx={{width: "100%"}}
                            id="email" {...register('email', {pattern: EMAIL_REG_EXP})} 
                            placeholder="электронная почта..." 
                            error={errors.email !== undefined} />
                        {errors.email && <h5 style={{color: red[400], margin: 2, fontSize: 10}}>{errors.email.message}</h5>}
                    </div>
                </div>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="password" className={register_styles.required}>Пароль:</InputLabel>
                    <div style={{margin: 10, display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
                        <OutlinedInput 
                            onFocus={() => setRegistrationState("")}
                            sx={{width: "100%"}}
                            id="password" {...register('password', {pattern: PASSWORD_REG_EXP, onChange: () => trigger("password_confirm")})} 
                            placeholder="пароль..." 
                            type={showPassword ? 'text' : 'password'}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={handleClickShowPassword}>
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                    </IconButton>
                                </InputAdornment>
                            }
                            error={errors.password !== undefined}  />
                            {errors.password && <h5 style={{color: red[400], margin: 2, fontSize: 10, wordWrap: 'break-word', maxWidth: 200}}>{errors.password.message}</h5>}
                    </div>
                </div>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="password_confirm" className={register_styles.required}>Повтори пароль:</InputLabel>
                    <div style={{margin: 10}}>
                        <OutlinedInput 
                            onFocus={() => setRegistrationState("")}
                            sx={{width: "100%"}}
                            id="password_confirm" {...register('password_confirm', {validate: 
                                (value, formValues) => value === formValues.password})}
                            placeholder="повтори пароль..." 
                            type={showPassword ? 'text' : 'password'}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={handleClickShowPassword}>
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                    </IconButton>
                                </InputAdornment>
                            }
                            error={errors.password_confirm!== undefined}  
                        />
                        {errors.password_confirm && <h5 style={{color: red[400], margin: 2, fontSize: 10}}>{errors.password_confirm.message}</h5>}
                    </div>
                </div>
                <Button sx={{
                    '&.MuiButton-outlined': {
                        color: grey[900],
                        borderColor: grey[900]
                    },
                    '&.Mui-disabled': {
                        color: red[100],
                        borderColor: red[100]
                    },
                    '&.MuiButton-loading': {
                        color: grey[900],
                        borderColor: grey[900],
                        fontSize: 0
                    }
                }} style={{margin: 10}} type='submit' variant='outlined' disabled={checkErrors(errors)} loading={loading}>Зарегистрироваться</Button>
            </form>
            <div className="buttons">
                <ReturnToLoginButton />
            </div>
            <h1 style={{color: red[400]}}>ОЦРВ</h1>
         </div>
        </>
    )
} 