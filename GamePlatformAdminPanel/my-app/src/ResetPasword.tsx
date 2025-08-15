import { FieldErrors, useForm } from 'react-hook-form'
import reset_password_styles from './css_modules/style.module.css'
import { ResetPasswordForm } from './app/api_forms_interfaces'
import { PASSWORD_REG_EXP } from './reg-exp'
import { useNavigate, useParams } from 'react-router'
import ResetPasswordFormResolver from './validate/form_resolvers/reset_password_resolver'
import { fetch_reset_password } from './api/resetPasswordApi'
import { useState } from 'react'
import { grey, red } from '@mui/material/colors'
import { Button, IconButton, InputAdornment, InputLabel, OutlinedInput } from '@mui/material'
import { Visibility, VisibilityOff } from '@mui/icons-material';
import { LOGIN_PATH } from './BrowserPathes'

export default function ResetPassword() {
    const navigate = useNavigate()
    const { id } = useParams();
    const [showPassword, setShowPassword] = useState<boolean>(false)
    const [loading, setLoading] = useState<boolean>(false) 
    const [statusReset, setStatusReset] = useState<string>("")

    const handleClickShowPassword = () => setShowPassword((show) => !show)

    const {register, handleSubmit, formState: { errors }, trigger, reset } = useForm<ResetPasswordForm>( 
        { 
            resolver: ResetPasswordFormResolver, 
            mode: 'onChange', 
            defaultValues: { password: "", password_confirm: "" } 
        } 
    )
    
    const onResetPassword = (data: ResetPasswordForm) => {
        setLoading(true)
        fetch_reset_password(data, id ?? "").then((response) => {
            setStatusReset("Пароль сброшен")
            window.setTimeout(() => navigate(LOGIN_PATH), 2000)
            console.log(response)
            setLoading(false)
        }, (reason) => {
            setStatusReset("не получилось сбросить пароль")
            console.log(reason)
            setLoading(false)
        })
        reset()
    }

    const checkErrors = (errors: FieldErrors<ResetPasswordForm>) : boolean => {
        return ((errors.password !== undefined) || 
            (errors.password_confirm !== undefined))
    }
    
    return (
        <>
            <form onSubmit={handleSubmit((data) => onResetPassword(data))}>
                <h1 style={{color: grey[900]}}>Изменение пароля</h1>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="password" className={reset_password_styles.required}>Введи новый пароль:</InputLabel>
                    <div style={{margin: 10}}>
                        <OutlinedInput
                            sx={{width: "100%"}}
                            id='password'
                            {...register('password', {pattern: PASSWORD_REG_EXP, onChange: () => trigger("password_confirm")})} 
                            placeholder="пароль..."
                            error={errors.password !== undefined} 
                            type={showPassword ? 'text' : 'password'}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={handleClickShowPassword}>
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                    </IconButton>
                                </InputAdornment>
                            } />
                        {errors.password && <h5 style={{color: red[400], margin: 2, fontSize: 10, wordWrap: 'break-word', maxWidth: 200}}> {errors.password?.message}</h5>}
                    </div>
                </div>
                <div>
                    <InputLabel style={{padding: 10}} htmlFor="confirm_password" className={reset_password_styles.required}>Повтори пароль:</InputLabel>
                    <div style={{margin: 10}}>
                        <OutlinedInput 
                            sx={{width: "100%"}}
                            id='confirm_password' 
                            error={errors.password_confirm !== undefined}
                            {...register('password_confirm', {validate: (value, formValues) => value === formValues.password})} 
                            placeholder="повтори пароль..."
                            type={showPassword ? 'text' : 'password'}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={handleClickShowPassword}>
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                    </IconButton>
                                </InputAdornment>
                            } />
                        {errors.password_confirm && <h5 style={{color: red[400], margin: 2, fontSize: 10, wordWrap: 'break-word', maxWidth: 200}}> {errors.password_confirm?.message}</h5>}
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
                }} style={{margin: 10}} type='submit' variant='outlined' disabled={checkErrors(errors)} loading={loading}>Изменить пароль</Button>
                {<h5>{statusReset}</h5>}
            </form>
        </>
    )
}