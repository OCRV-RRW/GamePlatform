import { Box } from "@mui/material";
import WarningRoundedIcon from '@mui/icons-material/WarningRounded';
import { red } from "@mui/material/colors";

export default function ClientWarningPage() {
    return <>
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
                <h1 style={{color: red[700], margin: 1, fontSize: 100}}>4xx</h1>
                <h1>Ой... Я в чем-то провинился</h1>
            </Box>
    </>
}