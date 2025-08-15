import { CircularProgress } from "@mui/material";
import { red } from "@mui/material/colors";

export default function Loader() {
    return (
        <>
            <CircularProgress sx={{color: red[500]}} />
        </>
    )
}