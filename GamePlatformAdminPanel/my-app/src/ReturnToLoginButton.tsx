import { useNavigate } from "react-router"
import { LOGIN_PATH } from "./BrowserPathes"
import { Button } from "@mui/material"
import { grey } from "@mui/material/colors"

export default function ReturnToLoginButton() {
    const navigate = useNavigate()

    return (
        <>
            <Button sx={{
                    '&.MuiButton-outlined': {
                        color: grey[900],
                        borderColor: grey[900]
                    }}}
                    style={{margin: 10}}
                name="back_to_login"
                type='submit' variant='outlined'
                onClick={() => navigate(LOGIN_PATH)}>
                    Вернуться к логину
            </Button>
        </>
    )
}