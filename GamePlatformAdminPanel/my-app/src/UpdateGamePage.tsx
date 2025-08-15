import { ChangeEvent, useCallback, useEffect, useState } from "react"
import { useParams } from "react-router"
import { Game } from "./app/game_type"
import { fetch_get_game } from "./api/getGameApi"
import { useAppDispatch } from "./app/hooks"
import { updateToken } from "./reducers/UserSlice"
import AdminHeader from "./Header"
import { useForm } from "react-hook-form"
import Box from "@mui/material/Box"
import { Button, TextField } from "@mui/material"
import { UpdateGameForm, UploadIconForm } from "./app/api_forms_interfaces"
import { fetch_update_game } from "./api/updateGameApi"
import './index.css'
import Loader from "./Loader"
import { FORBIDDEN, NOT_FOUND } from "./ResponseCodes"
import { set_status } from "./reducers/PageSlice"
import { ADMIN_PANEL_PATH } from "./BrowserPathes"
import { Preview } from "./app/preview_type"
import UpdateGamePreviewItem from "./UpdateGamePreviewItem"
import { fetch_upload_icon } from "./api/uploadGameIconApi"
import LoadFileComponent from "./LoadFileComponent"

type UpdateGameFormFields = {
    description: string,
    title: string,
    icon: string,
    src: string
}

export default function UpdateGamePage() {
    const dispatch = useAppDispatch()
    const { name } = useParams()
    const [gameData, setGameData] = useState<Game>()
    const [isSetValues, setIsSetValues] = useState<boolean>(false)
    const [imageFile, setImageFile] = useState<File | null>(null)
    const [icon, setIcon] = useState<string | undefined>();
    const [gamePreviews, setGamePreviews] = useState<Array<Preview>>([])
    const [loadingImage, setLoadingImage] = useState<boolean>(false)

    const { register, handleSubmit, formState: {errors}, reset, control, getValues, setValue } = useForm<UpdateGameFormFields>(
        {
            mode: 'onChange',
            defaultValues: {description: "", title: "", icon: "", src: ""}
        }
    )

    const fetchData = useCallback(() => {
        if (!name) return
        fetch_get_game(name)
        .then(async (fetch_data) => {
            dispatch(updateToken({access_token: fetch_data.access_token}))
            return fetch_data.response.json().then((json) => {
                setGameData(json.data.game as Game)
            })
        }, (reason) => {
            if (reason === FORBIDDEN.toString()) {
                dispatch(updateToken({access_token: ""}))
                return
            }
            if (reason === NOT_FOUND.toString()) return
            dispatch(set_status(reason))
        })
    }, [])

    useEffect(() => {
        fetchData()
    }, [fetchData])

    useEffect(() => {
        if (!gameData) return
        console.log(gameData)
        setGamePreviews(gameData.previews ?? [])
        setValue('description', gameData.description)
        setValue('title', gameData.title)
        setValue('src', gameData.src) 
        setIcon(gameData.icon)
        setIsSetValues(true)
    }, [gameData])

    const onUpdateGame = (form_data: UpdateGameFormFields) => {
        let updateGameFormData: UpdateGameForm = {
            description: form_data.description,
            title: form_data.title,
            icon: form_data.icon,
            src: form_data.src
        }
        console.log(updateGameFormData)
        fetch_update_game(updateGameFormData, name ?? "").then(() => {
            window.location.reload()
        }, (reason) => {
            if (reason === FORBIDDEN.toString()) {
                dispatch(updateToken({access_token: ""}))
                return
            }
            dispatch(set_status(reason))
        }) 
    }

    const loadImage = (event: ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0] || null
        if (!file) return
        setImageFile(file)
        setIcon(URL.createObjectURL(file))
    }

    const addGamePreview = () => {
        setGamePreviews([...gamePreviews, ({id: null, image: "", video: ""}) as Preview])
    }

    useEffect(() => {
        if (!imageFile) return
        if (!gameData?.id) return
        let upload_icon_form : UploadIconForm = {image: imageFile}
        
        fetch_upload_icon(upload_icon_form, gameData.id)
            .then(() => {
                console.log('good')
            }, (reason) => {
                console.log(reason)
            })
    }, [imageFile, icon])

    return(<>
        <AdminHeader pathToPage={ADMIN_PANEL_PATH} />
        {!gameData && <Loader />}
        {gameData && 
        <>
        <Box sx={{display: 'flex', justifyContent: 'center', alignItems: 'center', flexDirection: 'row'}}>
            <LoadFileComponent 
                buttonText="Выберите картинку"
                emptyWindowText="Нет картинки"
                url={icon}
                loadFunction={loadImage}
                accept="image/*"
                isVideo={false}
                />
            </Box>
            <Box sx={{display: 'flex', justifyContent: 'center', alignItems: 'center', padding: 2, flexDirection: 'column'}}>
                {gamePreviews &&
                gamePreviews.map((p, index) => 
                    <UpdateGamePreviewItem key={index} game_id={gameData.id} initPreviewData={p} />
                )}
            </Box>
            <Button style={{margin: 10}} type='submit' variant='outlined' onClick={addGamePreview}>
                +
            </Button>
         <form onSubmit={handleSubmit((data) => onUpdateGame(data))}>
            <div style={{margin: 10}}>
                <TextField id="description" {...register('description')} placeholder="описание..." label="Описание игры" />
            </div>
            <div style={{margin: 10}}>
                <TextField id="title" {...register('title')} placeholder="название..." label="Название игры" />
            </div>
            <div style={{margin: 10}}>
                <TextField id="icon" {...register('icon')} placeholder="иконка..." label="Иконка" />
            </div>
            <div style={{margin: 10}}>
                <TextField id="src" {...register('src')} placeholder="источник..." label="Источник игры (URL)" />
            </div>
            <Button style={{margin: 10}} type='submit' variant='outlined'>Обновить</Button>
         </form>
        </>
        }
    </>)
}