import { useState } from "react";
import AvatarColorCalculator from "./app/avatar_color_calculator";
import { User } from "./app/user_type";
import { Avatar, IconButton, Menu, MenuItem, Tooltip } from "@mui/material";
import Logout from "./Logout";

interface ProfileMenuProps {
    user_data: User | null
}

export function ProfileMenu({user_data}: ProfileMenuProps) {
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const avatarColor = new AvatarColorCalculator(user_data?.name, user_data?.email).calculateColors()
    const open = Boolean(anchorEl);

    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return(
        <>
            {user_data && <>
                <Tooltip title={user_data.name}>
                    <IconButton
                        size='large'    
                        sx={{ ml: 2, width: 40, height: 40, margin: 1 }}
                        aria-controls={open ? 'profile-menu' : undefined}
                        aria-haspopup="true"
                        onClick={handleClick}
                        aria-expanded={open ? 'true' : undefined}>
                        <Avatar sx={{ bgcolor: `rgba(${avatarColor.r}, ${avatarColor.g}, ${avatarColor.b}, 1)` }}>{user_data.name?.at(0)?.toUpperCase()}</Avatar>
                    </IconButton>
                </Tooltip>
                <Menu 
                    anchorEl={anchorEl} 
                    id="profile-menu"
                    open={open}
                    onClose={handleClose}
                    slotProps={{
                        paper: {
                            elevation: 0,
                            sx: {
                                textAlign: 'center',
                                backgroundColor: 'rgba(0, 0, 0, 0.5)',
                                overflow: 'visible',
                                mt: 1.5
                            }
                        }
                    }}
                    transformOrigin={{ horizontal: 'right', vertical: 'top' }}
                    anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}>
                        <MenuItem sx={{justifyContent: 'center'}}>
                            <Logout />
                        </MenuItem>
                </Menu>
            </>}
        </>
    )
}