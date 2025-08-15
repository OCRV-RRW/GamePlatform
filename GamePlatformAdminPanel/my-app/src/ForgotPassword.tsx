import { useState } from "react";
import ReturnToLoginButton from "./ReturnToLoginButton";
import { fetch_forgot_password } from "./api/forgotPasswordApi";
import { grey, red } from "@mui/material/colors";
import { Button, InputLabel, TextField } from "@mui/material";
import styles from './css_modules/style.module.css'
import { NOT_FOUND } from "./ResponseCodes";

export default function ForgotPassword() {
    const [email, setEmail] = useState<string>("")
    const [forgotPasswordState, setForgotPasswordState] = useState<string>("")

    const onRestorePassword = () => {
        fetch_forgot_password({email: email})
            .then(() => setForgotPasswordState("Перейди на почту"), 
                (reason) => 
                    {
                        if (reason === NOT_FOUND.toString()) {
                            setForgotPasswordState("Проверь правильность почты")
                        }
                    })
    }

    return (
        <>
            <h1 style={{color: grey[900]}}>Забыл пароль</h1>
            <div className="email">
                <InputLabel htmlFor="email" className={styles.required} style={{padding: 10}}>Электронная почта:</InputLabel>
                <div style={{margin: 10}}>
                    <TextField id="email" value={email} placeholder="электронная почта..." label="Электронная почта" onChange={(event) => setEmail(event.target.value)} />
                </div>
            </div>
            <div style={{display: 'flex', flexDirection: 'column',  alignItems: 'center'}} className="buttons">
                <Button sx={{
                    '&.MuiButton-outlined': {
                        color: grey[900],
                        borderColor: grey[900]
                    }}}
                    onClick={onRestorePassword}
                    style={{margin: 10}} type='submit' variant='outlined'>Восстановить пароль
                </Button>
                <ReturnToLoginButton />
            </div>
            <h1 style={{color: red[200]}}>{forgotPasswordState}</h1>
            <h1 style={{color: red[400]}}>ОЦРВ</h1>
        </>
    )
}