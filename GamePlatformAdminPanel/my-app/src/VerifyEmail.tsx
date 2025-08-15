import { useEffect, useState } from "react"
import { fetch_verify_email } from "./api/verifyEmailApi"
import { useNavigate, useParams } from "react-router"
import { LOGIN_PATH } from "./BrowserPathes";
import Loader from "./Loader";

export default function VerifyEmail() {
    const { id } = useParams();
    const [verifyMessage, setVerifyMessage] = useState<string>("")
    const navigate = useNavigate()
    const EMAIL_VERIFATED_MESSAGE = "Почта подтверждена"
    const EMAIL_FAILED_VERIFY_MESSAGE = "Ошибка подтверждения почты. Попробуйте снова"

    useEffect(() => {
        fetch_verify_email(id ?? "")
            .then(() => {
                    setVerifyMessage(EMAIL_VERIFATED_MESSAGE)
                    window.setTimeout(() => navigate(LOGIN_PATH), 2000)
                },
                (reason) => {
                    console.log(reason)
                    setVerifyMessage(EMAIL_FAILED_VERIFY_MESSAGE)
                }
            )
    }, [id])

    return (
        <>
            <h1>{verifyMessage === "" ? <Loader /> : verifyMessage}</h1>
        </>
    )
}