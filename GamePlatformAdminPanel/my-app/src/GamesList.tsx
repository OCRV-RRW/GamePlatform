import { createContext, useEffect, useState } from "react"
import { fetch_get_games } from "./api/getGamesApi"
import { useAppDispatch } from "./app/hooks"
import { updateToken } from "./reducers/UserSlice"
import { Game } from "./app/game_type"
import styles from './css_modules/style.module.css'
import { grey } from "@mui/material/colors"
import Loader from "./Loader"
import { AdminListItem } from "./AdminListItem"
import { UPDATE_GAME_PATH } from "./BrowserPathes"
import { AdminList } from "./AdminList"
import { fetch_delete_game } from "./api/deleteGameApi"
import { CreateGameForm } from "./app/api_forms_interfaces"
import { fetch_create_game } from "./api/createGameApi"
import { set_status } from "./reducers/PageSlice"
import { FORBIDDEN, NOT_FOUND } from "./ResponseCodes"
import { BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP } from "./reg-exp"
import { AddGameEntityDialog } from "./AddGameEntityDialog"

export const OpenCreateDialogWindowContext = createContext<(isOpen: boolean) => void>(() => {})
type GameListGamesName = {
    id: string,
    title: string
}


export default function GamesList() {
    const dispatch = useAppDispatch()
    const [gameNames, setGameNames] = useState<Array<GameListGamesName>>([])
    const [loading, setLoading] = useState<boolean>(false)
    const [createDialogWindowOpen, setCreateDialogWindowOpen] = useState<boolean>(false)
    
    const fetch_games = () => {
        setLoading(true)
        fetch_get_games()
            .then(async (fetch_data) => {
                dispatch(updateToken({access_token: fetch_data.access_token}))
                return fetch_data.response.json().then((json) => {
                    let games : Array<Game> = json.data.games as Array<Game>
                    setGameNames(games?.map<GameListGamesName>((game) => {
                        return {id: game.id, title: game.title}
                    }))
                    setLoading(false)
                })
            }, (reason) => {
                setLoading(false)
                if (reason === FORBIDDEN.toString()) {
                    dispatch(updateToken({access_token: ""}))
                    return
                }
                if (reason === NOT_FOUND.toString()) return
                dispatch(set_status(reason))
            })
    }

    useEffect(() => {
        fetch_games()
    }, [])

    return (
        <>
            <h1 style={{color: grey[500]}}>Список игр</h1>
            {loading && <Loader />}
            <OpenCreateDialogWindowContext.Provider value={(isOpen) => setCreateDialogWindowOpen(isOpen)}>
            {!loading && 
                <div className={styles.scrollableContainer}>
                    <AdminList>
                        {gameNames && gameNames.length > 0 ? gameNames?.map((g) => 
                            <AdminListItem 
                                key={g.id} 
                                title={g.title} 
                                eleName={g.id} 
                                update_path={UPDATE_GAME_PATH} 
                                delete_fetch={() => fetch_delete_game(g.id).then((data) => {
                                    fetch_games()
                                    dispatch(updateToken({access_token: data.access_token}))
                            }, (reason) => 
                                { 
                                    if (BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP.test(reason)) {
                                        dispatch(updateToken({access_token: ""}))
                                        return
                                    }
                                    dispatch(set_status(reason)) 
                                })} 
                            />) : <h2>Пока нет игр</h2>}
                    </AdminList>
                </div>}
                {<AddGameEntityDialog 
                    isOpen={createDialogWindowOpen}
                    createGameEntityFetch={(form: CreateGameForm) => fetch_create_game(form).then((data) => {
                        fetch_games()
                        dispatch(updateToken({access_token: data.access_token}))
                    }, (reason) => {
                        if (BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP.test(reason)) {
                            dispatch(updateToken({access_token: ""}))
                            return
                        }
                        dispatch(set_status(reason))
                    })}/>}
            </OpenCreateDialogWindowContext.Provider>
        </>
    )
}