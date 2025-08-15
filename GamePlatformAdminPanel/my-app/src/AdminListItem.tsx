import { Box, Divider, ListItem } from "@mui/material"
import { grey } from "@mui/material/colors"
import styles from './css_modules/style.module.css'
import { DeleteButtonAdminListItem, UpdateButtonAdminListItem } from "./ButtonsAdminList";

interface AdminListItemProps {
    title: string,
    eleName: string,
    update_path?: string,
    query?: string,
    delete_fetch?: () => Promise<void>,
}

export function AdminListItem({title, eleName, update_path, query, delete_fetch}: AdminListItemProps) {
    return <ListItem sx={{paddingBottom: 0}}>
    <Box className={styles.gameListLink}>
        <h4 style={{color: grey[500], paddingLeft: 10}}>{title}</h4>
        <Box sx={{display: 'flex', alignItems: 'center'}}>
            {update_path && <>
                <Divider orientation='vertical' sx={{borderColor: grey[100]}}/>
                <UpdateButtonAdminListItem eleName={eleName} path={update_path} query={query} />
            </>}
            {delete_fetch && <>
                <Divider orientation='vertical' sx={{borderColor: grey[100]}}/>
                <DeleteButtonAdminListItem eleName={eleName} deleteEntityFetch={delete_fetch!} />
            </>}
        </Box>
    </Box>
</ListItem>
}
