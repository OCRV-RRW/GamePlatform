import { useNavigate } from "react-router"
import { useAppSelector } from "../app/hooks"
import { selectUserData } from "../reducers/UserSlice"
import { useEffect } from "react"
import { ADMIN_PANEL_PATH } from "../BrowserPathes"

export interface CheckIsAdminProps {
    children: JSX.Element
}

export default function CheckNotIsAdmin({children} : CheckIsAdminProps) {
    const isAdmin = useAppSelector(selectUserData)?.is_admin
    const navigate = useNavigate()

    useEffect(() => {
        if (isAdmin)
            navigate(ADMIN_PANEL_PATH)
    }, [isAdmin, navigate])
    
    return <>
        {!isAdmin && children}
    </>
}