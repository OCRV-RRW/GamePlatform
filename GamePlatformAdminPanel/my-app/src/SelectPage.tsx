import { useContext, useEffect, useRef, useState } from "react";
import { PathContext } from "./Page";
import { 
    ADMIN_PANEL_PATH,
    FORGOT_PASSWORD_PATH,
    LOGIN_PATH, 
    NOT_ADMIN_WARNING_PAGE_PATH, 
    REGISTER_PATH, 
    REGISTER_VERIFY_EMAIL_PATH, 
    RESET_PASSWORD_PATH, 
    UPDATE_GAME_PATH,
    UPDATE_USER_PATH
} from "./BrowserPathes";
import Register from "./Register";
import Login from "./Login";
import VerifyEmail from "./VerifyEmail";
import UpdateGamePage from "./UpdateGamePage";
import AdminPanelHome from "./Home";
import UpdateUserPage from "./UpdateUserPage";
import ForgotPassword from "./ForgotPassword";
import ResetPassword from "./ResetPasword";
import NotAdminWarningPage from "./NotAdminWarningPage";

export default function SelectPath() {
    const path = useContext(PathContext)
    const [currentComponent, setCurrentComponent] = useState<JSX.Element>()
    const pages = useRef<{path: string, page: JSX.Element}[]>(
        [
            {path: REGISTER_PATH, page: <Register />},
            {path: LOGIN_PATH, page: <Login />},
            {path: FORGOT_PASSWORD_PATH, page: <ForgotPassword />},
            {path: REGISTER_VERIFY_EMAIL_PATH, page: <VerifyEmail />},
            {path: RESET_PASSWORD_PATH, page: <ResetPassword />},
            {path: UPDATE_GAME_PATH, page: <UpdateGamePage />},
            {path: ADMIN_PANEL_PATH, page: <AdminPanelHome />},
            {path: UPDATE_USER_PATH, page: <UpdateUserPage />},
            {path: NOT_ADMIN_WARNING_PAGE_PATH, page: <NotAdminWarningPage />}
        ]
    )
    
    const selectComponent = (path_context: string) : JSX.Element => {
        const el = pages.current.find(({path}) => path === path_context)
        return el?.page!
    }

    useEffect(() => {
        setCurrentComponent(selectComponent(path))
    }, [path])

    return(
        <>
            {currentComponent}
        </>
    )
}