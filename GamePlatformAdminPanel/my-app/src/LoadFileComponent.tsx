import { Box, Button, styled } from "@mui/material";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { ChangeEvent } from "react";

export interface LoadFileProps {
    buttonText: string,
    emptyWindowText: string,
    url: string | undefined,
    loadFunction: (event: ChangeEvent<HTMLInputElement>) => void,
    accept: string,
    isVideo: boolean
}

export default function LoadFileComponent({buttonText, emptyWindowText, url, loadFunction, accept, isVideo}: LoadFileProps) {
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
    return (
        <Box sx={{display: 'flex', flexDirection: 'column', width:"300px"}}>
                    <Button
                            component="label"
                            role={undefined}
                            variant="contained"
                            tabIndex={-1}
                            startIcon={<CloudUploadIcon />}
                            sx={{margin: 1}}
                            >
                            {buttonText}
                            <VisuallyHiddenInput
                                type="file"
                                accept={accept}
                                onChange={loadFunction}
                            />
                        </Button>
                        {!url
                    ? <>
                        <Box style={{width: "300px", height: "300px"}}>
                            <h5>{emptyWindowText}</h5>
                        </Box>
                    </>
                    : <>
                        {isVideo ? 
                        <video style={{width: "300px", height: "300px"}}
                            src={url}
                        /> : 
                        <img style={{width: "300px", height: "300px"}}
                            src={url}
                        />
                        }
                    </>}
                    </Box>
    )
}