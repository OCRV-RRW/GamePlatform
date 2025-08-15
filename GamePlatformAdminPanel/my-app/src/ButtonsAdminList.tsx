import { Button } from "@mui/material"
import { useNavigate } from "react-router"
import { green, red } from "@mui/material/colors"
import UpdateIcon from '@mui/icons-material/Create';
import CreateIcon from '@mui/icons-material/Add';
import DeleteIcon from '@mui/icons-material/DeleteForever';
import { OpenCreateDialogWindowContext } from "./GamesList"

export interface ButtonAdminListItemProps {
    eleName?: string,
    path?: string,
    query?: string
}

export function UpdateButtonAdminListItem({path, eleName, query}: ButtonAdminListItemProps) {
    const navigate = useNavigate()
    return <Button onClick={() => query === undefined ? navigate(path + "/" + eleName) : navigate(path + query + eleName)}>
        <UpdateIcon />
    </Button>
}

export interface DeleteButtonAdminListItemProps extends ButtonAdminListItemProps {
    deleteEntityFetch: () => Promise<void>
}
export function CreateButtonAdminListItem() {
    return <>
        <OpenCreateDialogWindowContext.Consumer>
            {setOpen => (
                <Button sx={{color: green[800], margin: 1}} onClick={() => setOpen(true)}>
                    <CreateIcon />
                </Button>
            )}
        </OpenCreateDialogWindowContext.Consumer>
    </> 
}

export function DeleteButtonAdminListItem({deleteEntityFetch}: DeleteButtonAdminListItemProps) {
    async function onDelete()  {
        await deleteEntityFetch()
    }

    return <>
        <Button sx={{color: red[900]}} onClick={onDelete}>
            <DeleteIcon />
        </Button>
    </>
}