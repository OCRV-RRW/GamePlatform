import { useEffect, useState } from "react";
import { useAppSelector } from "../app/hooks";
import { selectStatusPage } from "../reducers/PageSlice";
import NotFoundPage from "./NotFoundPage";
import { BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP, BAD_STATUS_SERVER_RESPONSE_SERVER_WARNING_REG_EXP } from "../reg-exp";
import ServerWarningPage from "./ServerWarningPage";
import ClientWarningPage from "./ClientWarningPage";

interface ControllerStatusCodePageProps {
    children: JSX.Element
}

export default function ControllerStatusCodePage(props: ControllerStatusCodePageProps) {
    const status = useAppSelector(selectStatusPage)
    const [page, setPage] = useState<JSX.Element>()

    useEffect(() => {
        if (status === 'ok') { setPage(props.children); return; }
        if (status === '404') { setPage(<NotFoundPage/>); return; }
        if (BAD_STATUS_SERVER_RESPONSE_SERVER_WARNING_REG_EXP.test(status)) { setPage(<ServerWarningPage />); return; }
        if (BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP.test(status)) { setPage(<ClientWarningPage />); return; }
    }, [status])

    return <>
        {page!} 
    </>
}