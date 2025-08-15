import { Box } from "@mui/material";
import { red } from "@mui/material/colors";
import WarningRoundedIcon from '@mui/icons-material/WarningRounded';

export default function NotFoundPage() {
    return (
        <>
            <Box sx={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                flexDirection: 'column',
                width: '100vw',
                height: '90vh'
            }}>
                <WarningRoundedIcon 
                    sx={{
                            width: '10vmin',
                            height: '10vmin',
                            color: red[700]
                        }} 
                    />
                <h1 style={{color: red[700], margin: 1, fontSize: 100}}>404</h1>
                <h1>Упс... Страница не существует</h1>
            </Box>
        </>
    )
}