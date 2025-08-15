import { createContext, useEffect } from "react"
import { useAppDispatch, useAppSelector } from "./app/hooks"
import { selectCurrentPage, set_page } from "./reducers/PageSlice"
import SelectPath from "./SelectPage";

export interface PageProps {
    path: string
}

export const PathContext = createContext("");

export default function Path({path}: PageProps) {
    const dispath = useAppDispatch()
    const current_page = useAppSelector(selectCurrentPage)

    useEffect(() => {
        dispath(set_page(path))
    }, [path])

    return (
        <>
            <PathContext.Provider value={current_page}>
                <SelectPath />
            </PathContext.Provider>
        </>
    )
}