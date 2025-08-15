import { Box } from "@mui/material";
import { ChangePageFuncContext } from "./Home";
import { AdminPanelDrawerButton } from "./AdminPanelDrawerButton";
import UsersIcon from '@mui/icons-material/People';
import { red } from "@mui/material/colors";

export default function ToUsersListButton() {
    return (<>
        <ChangePageFuncContext.Consumer>
            {page => (
                <AdminPanelDrawerButton variant={page.page === 'users' ? 'contained' : 'text'} onClick={() => page.change_page('users')}>
                    <Box sx={{display: 'flex', alignItems: 'center'}}>
                        <UsersIcon sx={{padding: 1, color: red[400]}} />
                            Пользователи
                    </Box>
                </AdminPanelDrawerButton>
            )}
        </ChangePageFuncContext.Consumer>
    </>)
}