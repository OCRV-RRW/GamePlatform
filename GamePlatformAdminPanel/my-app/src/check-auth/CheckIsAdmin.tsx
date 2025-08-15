import { useNavigate } from "react-router"
import { useAppSelector } from "../app/hooks"
import { selectUserData } from "../reducers/UserSlice"
import { useEffect } from "react"
import { HOME_PATH } from "../BrowserPathes"

export interface CheckIsAdminProps {
    children: JSX.Element
}

export default function CheckIsAdmin({children} : CheckIsAdminProps) {
    const isAdmin = useAppSelector(selectUserData)?.is_admin
    const navigate = useNavigate()

    useEffect(() => {
        if (!isAdmin)
            navigate(HOME_PATH)
    }, [isAdmin, navigate])
    
    return <>
        {isAdmin && children}
    </>
}