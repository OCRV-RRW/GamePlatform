import AdminPanelSettingsIcon from '@mui/icons-material/AdminPanelSettings';
import { Button, Tooltip } from '@mui/material';
import { useNavigate } from 'react-router';
import { ADMIN_PANEL_PATH } from './BrowserPathes';
import { useAppSelector } from './app/hooks';
import { selectUserData } from './reducers/UserSlice';
import { useRef } from 'react';

export default function AdminPanelButton() {
    const navigate = useNavigate()
    const userData = useAppSelector(selectUserData)
    const visibility = useRef<'visible' | 'hidden'>(userData?.is_admin ? 'visible' : 'hidden')

    const toAdminPanel = () => {
        navigate(ADMIN_PANEL_PATH)
    }

    return <>
        <Tooltip title="перейти в панель администратора">
            <Button sx={{margin: 1, visibility: visibility.current}} onClick={toAdminPanel}>
                <AdminPanelSettingsIcon />
            </Button>
        </Tooltip>
    </>
}