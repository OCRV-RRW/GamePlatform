import { Box } from "@mui/material";
import WarningRoundedIcon from '@mui/icons-material/WarningRounded';
import { red } from "@mui/material/colors";
import { ProfileMenu } from "./ProfileMenu";
import { useAppSelector } from "./app/hooks";
import { selectUserData } from "./reducers/UserSlice";

export default function NotAdminWarningPage() {
    const user_data = useAppSelector(selectUserData)
    return (
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
            <h1>У вас нет прав администратора!!!</h1>
            <ProfileMenu user_data={user_data} />
    </Box>
    )
}