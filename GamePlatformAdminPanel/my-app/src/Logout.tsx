import { useAppDispatch } from "./app/hooks"
import { send_log_out, updateToken } from "./reducers/UserSlice"
import { Button, Tooltip } from "@mui/material"
import { grey } from "@mui/material/colors"
import LogoutOutlined from "@mui/icons-material/Logout"
import { unwrapResult } from "@reduxjs/toolkit"
import { useState } from "react"

export default function Logout() {
    const dispath = useAppDispatch()
    const [loading, setLoading] = useState<boolean>(false)

    const logout = () => {
        setLoading(true)
        dispath(send_log_out())
            .then(unwrapResult)
            .then(() => setLoading(false))
            .catch((error) => {
                setLoading(false)
                dispath(updateToken({access_token: ""}))
            })
    }

    return (
        <>
            <Tooltip title="Выйти">
                <Button sx={{
                        '&.MuiButton-text': {
                            color: grey[100]
                        },
                        '&.MuiButton-loading': {
                            color: grey[100]
                        }}}
                        type='submit' variant='text' onClick={logout}
                        loading={loading} >
                            {!loading && <LogoutOutlined />}
                </Button>
            </Tooltip>
        </>
    )
}