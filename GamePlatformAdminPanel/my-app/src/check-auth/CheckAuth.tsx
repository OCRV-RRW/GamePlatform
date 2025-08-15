import { useAppSelector } from "../app/hooks";
import { selectAccessToken } from "../reducers/UserSlice";
import { useNavigate } from "react-router";
import { LOGIN_PATH } from "../BrowserPathes";
import { useEffect } from "react";

interface CheckAuthProps {
    children: JSX.Element,
}

export default function CheckAuth({ children }: CheckAuthProps) {
    const access_token = useAppSelector(selectAccessToken)
    const navigate = useNavigate()

    useEffect(() => {
        if (access_token === "") {
            navigate(LOGIN_PATH)
        }
        
    }, [access_token, navigate])

    return (
        <>
            {access_token !== "" && children}
        </>
    )
}