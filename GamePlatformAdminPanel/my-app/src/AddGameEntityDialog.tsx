import { CreateGameForm } from "./app/api_forms_interfaces"
import { Button, Dialog, DialogActions, DialogContent, DialogTitle, TextField } from "@mui/material"
import { OpenCreateDialogWindowContext } from "./GamesList"
import { useAppDispatch } from "./app/hooks"
import { updateToken } from "./reducers/UserSlice"

export interface AddGameEntityDialogProps {
    createGameEntityFetch: (form: CreateGameForm) => Promise<void>
    isOpen: boolean
}

export function AddGameEntityDialog({createGameEntityFetch: createGameEntityFetch, isOpen}: AddGameEntityDialogProps) {
    const dispatch = useAppDispatch()

    return <>
        <OpenCreateDialogWindowContext.Consumer>
            {setOpen => (
                <Dialog open={isOpen}
                    slotProps={{
                        paper: {
                            component: 'form',
                            onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
                                event.preventDefault()
                                const form_data = new FormData(event.currentTarget)
                                const form_json = Object.fromEntries((form_data as any).entries())
                                const created_entity_title = form_json.title
                                const created_entity_description = form_json.description
                                const created_entity_src = form_json.src
                                const created_entity_icon = form_json.icon

                                createGameEntityFetch({
                                    'title': created_entity_title,
                                    'description': created_entity_description,
                                    'src': created_entity_src,
                                    'icon': created_entity_icon
                                }).then(() => setOpen(false))    
                            }
                        }
                    }}>
                    <DialogTitle>Создание игры</DialogTitle>
                    <DialogContent>
                        <TextField 
                            autoFocus
                            required
                            margin="dense"
                            id="title"
                            name="title"
                            label="название игры"
                            type='text'
                            fullWidth
                            variant='outlined'/>
                        <TextField 
                            autoFocus
                            required
                            margin="dense"
                            id="description"
                            name="description"
                            label="описание"
                            type='text'
                            fullWidth
                            variant='outlined' />
                        <TextField 
                            autoFocus
                            required
                            margin="dense"
                            id="src"
                            name="src"
                            label="url на источник"
                            type='text'
                            fullWidth
                            variant='outlined' />
                        <TextField 
                            autoFocus
                            required
                            margin="dense"
                            id="icon"
                            name="icon"
                            label="иконка"
                            type='text'
                            fullWidth
                            variant='outlined' />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={() => setOpen(false)}>Отменить</Button>
                        <Button type='submit'>Создать</Button>
                    </DialogActions>
                </Dialog>
            )}
        </OpenCreateDialogWindowContext.Consumer>
    </>
}