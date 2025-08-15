import { Box, Button, styled } from "@mui/material";
import { Preview } from "./app/preview_type";
import { red } from "@mui/material/colors";
import { ChangeEvent, useEffect, useState } from "react";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { fetch_create_preview } from "./api/createPreviewApi";
import { CreatePreviewForm } from "./app/api_forms_interfaces";
import { useAppDispatch } from "./app/hooks";
import { updateToken } from "./reducers/UserSlice";
import { FORBIDDEN, NOT_FOUND } from "./ResponseCodes";
import { set_status } from "./reducers/PageSlice";
import { fetch_delete_preview } from "./api/deletePreviewApi";
import LoadFileComponent from "./LoadFileComponent";

export interface UpdateGamePreviewItemProps {
    game_id: string,
    initPreviewData: Preview
}

export default function UpdateGamePreviewItem({initPreviewData, game_id} : UpdateGamePreviewItemProps) {
    const dispatch = useAppDispatch()
    const VisuallyHiddenInput = styled('input')({
    clip: 'rect(0 0 0 0)',
    clipPath: 'inset(50%)',
    height: 1,
    overflow: 'hidden',
    position: 'absolute',
    bottom: 0,
    left: 0,
    whiteSpace: 'nowrap',
    width: 1,
    });
    const [previewData, setPreviewData] = useState<Preview>(initPreviewData)
    const [imageFile, setImageFile] = useState<File | null>(null)
    const [image, setImage] = useState<string | undefined>(previewData.image);
    const [video, setVideo] = useState<string | undefined>(previewData.video);
    const [videoFile, setVideoFile] = useState<File | undefined>(undefined)

    const loadImage = (event: ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0] || null
        if (!file) return
        setImageFile(file)
        setImage(URL.createObjectURL(file))
    }

    useEffect(() => {
        setImage(previewData.image)
        setVideo(previewData.video)
    }, [previewData])

    const loadVideo = (event: ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0] || null
        if (!file) return
        setVideoFile(file)
        setVideo(URL.createObjectURL(file))
    }

    const createGamePreviewItem = () => {
        if (imageFile === null) return
        const createPreviewBody : CreatePreviewForm = {
            game_id,
            image: imageFile,
            video: videoFile
        }
        fetch_create_preview(createPreviewBody)
        .then(async (fetch_data) => {
            dispatch(updateToken({access_token: fetch_data.access_token}))
            return fetch_data.response.json().then((json) => {
                setPreviewData(json.data.preview)
            })
        }, (reason) => {
            if (reason === FORBIDDEN.toString()) {
                dispatch(updateToken({access_token: ""}))
                return
            }
            if (reason === NOT_FOUND.toString()) return
            dispatch(set_status(reason))
        })
    }

    const deleteGamePreviewItem = () => {
        if (previewData.id === null) return
        fetch_delete_preview(previewData.id)
        .then(async (fetch_data) => {
            dispatch(updateToken({access_token: fetch_data.access_token}))
            window.location.reload()
        }, (reason) => {
            if (reason === FORBIDDEN.toString()) {
                dispatch(updateToken({access_token: ""}))
                return
            }
            if (reason === NOT_FOUND.toString()) return
            dispatch(set_status(reason))
        })
    }

    return (
        <>
        <Box sx={{display: 'flex', justifyContent: 'center', alignItems: 'center', flexDirection: 'row'}}>
            <LoadFileComponent 
                buttonText="Выберите картинку" 
                emptyWindowText="Нет картинки" 
                url={image}
                loadFunction={loadImage}
                accept="image/*" 
                isVideo={false}/>
            <LoadFileComponent 
                buttonText="Выберите видео" 
                emptyWindowText="Нет видео" 
                url={video!}
                loadFunction={loadVideo}
                accept="video/*"
                isVideo={true} />
            <Button
                    component="label"
                    role={undefined}
                    variant="contained"
                    disabled={previewData.id !== null}
                    tabIndex={-1}
                    sx={{margin: 1}}
                    onClick={createGamePreviewItem}
                    >
                    Создать
            </Button>
            <Button
                    component="label"
                    role={undefined}
                    variant="contained"
                    disabled={previewData.id === null}
                    tabIndex={-1}
                    sx={{margin: 1}}
                    onClick={deleteGamePreviewItem}
                    >
                    Удалить
            </Button>
        </Box>
        </>
    )
}