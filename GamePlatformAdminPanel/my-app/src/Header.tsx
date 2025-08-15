import { Box } from '@mui/material'
// import GoToOtherPageButton from '../../GoToHomeButton'
import { grey } from '@mui/material/colors'
import { ProfileMenu } from './ProfileMenu'
import { useAppSelector } from './app/hooks'
import { selectUserData } from './reducers/UserSlice'

interface AdminHeaderProps {
    pathToPage: string
}

export default function Header({ pathToPage } : AdminHeaderProps) {
    const user_data = useAppSelector(selectUserData)
    return (
        <>
            <Box sx={{ display: 'flex',
                alignItems: 'center',
                justifyContent: 'space-between'}}>
                {/* <GoToOtherPageButton pathToPage={pathToPage} /> */}
                <h2 style={{margin: 10, color: grey[900], fontSize: 24}}>Админ-панель</h2>
                <ProfileMenu user_data={user_data}/>
            </Box>
        </>
    )
}