import { List, Pagination } from "@mui/material"
import { CreateButtonAdminListItem } from "./ButtonsAdminList";

interface AdminListProps { 
    listWithCreate?: boolean,
    children: JSX.Element[] | JSX.Element | undefined
}

export function AdminList({listWithCreate = true, children} : AdminListProps) {
    return <>
        <List sx={{width: '100%'}}>
            {children}
        </List>
        {listWithCreate && <CreateButtonAdminListItem/>}
        <Pagination count={1} variant="outlined" shape="rounded" />
    </>
}