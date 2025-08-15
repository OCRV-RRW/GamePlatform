import { Button, colors, styled } from "@mui/material";
import { grey } from "@mui/material/colors";

export const AdminPanelDrawerButton = styled(Button)(({ theme }) => ({
    '&.MuiButton-text': {
        color: grey[400]
    },
    '&.MuiButton-contained': {
        backgroundColor: grey[700]
    },
    [theme.containerQueries.down(200)]: {
        fontSize: 0
    }
}))