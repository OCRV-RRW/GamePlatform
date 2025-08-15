import { useLocation } from "react-router"
import { useAppDispatch } from "./app/hooks"
import { useCallback, useEffect, useState } from "react"
import { User, UserGender } from "./app/user_type"
import { Controller, useForm } from "react-hook-form"
import { fetch_get_user } from "./api/getUserApi"
import { updateToken } from "./reducers/UserSlice"
import { FORBIDDEN } from "./ResponseCodes"
import { set_status } from "./reducers/PageSlice"
import AdminHeader from "./Header"
import Loader from "./Loader"
import { UpdateUserForm } from "./app/api_forms_interfaces"
import { fetch_update_user } from "./api/updateUserApi"
import { Box, Button, FormControl, MenuItem, Select, Switch, TextField, Tooltip } from "@mui/material"
import { blue, green, grey, red } from "@mui/material/colors"
import { BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP } from "./reg-exp"
import style from "./css_modules/style.module.css"
import { ADMIN_PANEL_PATH } from "./BrowserPathes"
import SupervisorAccountIcon from '@mui/icons-material/SupervisorAccount';
import MaleIcon from '@mui/icons-material/Male';
import FemaleIcon from '@mui/icons-material/Female';

type UpdateUserFormFields = {
    birthday: string,
    gender: UserGender,
    is_admin: boolean,
    name: string
}

export default function UpdateUserPage() {
    const dispatch = useAppDispatch()
    const [userData, setUserData] = useState<User>()
    const location = useLocation()

    const { register, handleSubmit, control, getValues, setValue } = useForm<UpdateUserFormFields>(
        {
            mode: 'onChange',
            defaultValues: { birthday: "", gender: "", is_admin: false, name: "" }
        }
    )

    const fetchUser = useCallback(() => {
        fetch_get_user(location.search)
            .then(async (data) => {
                dispatch(updateToken({ access_token: data.access_token }))
                return data.response.json()
                    .then((json) => {
                        setUserData(json.data?.users[0] as User)
                    }
                )
            }, (reason) => {
                if (BAD_STATUS_SERVER_RESPONSE_CLIENT_WARNING_REG_EXP.test(reason)) {
                    dispatch(updateToken({access_token: ""}))
                    return
                }
                dispatch(set_status(reason))
            })
    }, [])

    const onUpdateUser = (form_data: UpdateUserFormFields) => {
        let birthday : Date = form_data.birthday ? new Date(form_data.birthday) : new Date(2000, 1, 1)
        let updateUserData: UpdateUserForm = {
            birthday: birthday,
            gender: form_data.gender,
            is_admin: form_data.is_admin,
            name: form_data.name
        }
        console.log('dsd')
        fetch_update_user(updateUserData, userData?.id ?? "")
            .then(() => {
                window.location.reload()
            },
            (reason) => {
                if (reason === FORBIDDEN.toString()) {
                    dispatch(updateToken({access_token: ""}))
                    return
                }
                dispatch(set_status(reason))
            })
    }

    console.log(userData)

    useEffect(() => {
        fetchUser()
    }, [fetchUser])

    useEffect(() => {
        if (!userData) return
        let birthday : string | null = 
            userData.birthday 
            ? userData.birthday.toString().slice(0, userData.birthday?.toString().indexOf("T")) 
            : null
        setValue('birthday', birthday!)
        setValue('gender', userData.gender)
        setValue('is_admin', userData.is_admin)
        setValue('name', userData.name ?? "")
    }, [userData])

    return (
        <>
            <AdminHeader pathToPage={ADMIN_PANEL_PATH} />
            {!userData 
                ? <Loader />
                : <>
                <h1 style={{color: grey[500], padding: 20}}>Пользователь: {userData.name}</h1>
                <form onSubmit={handleSubmit((data) => onUpdateUser(data))}>
                    <Box sx={{margin: 1}}>
                        <TextField id="name" {...register('name')} placeholder="имя..." label="Имя" />
                    </Box>
                    <Box sx={{
                        margin: 1
                    }}>
                        <Box sx={{
                            maxWidth: 223, width: "100%", display: 'inline-flex', position: 'relative', zIndex: 0
                        }}>
                                <Box sx={{
                                    maxWidth: 226, 
                                    width: "100%", 
                                    display: 'inline-flex', 
                                    position: 'relative',
                                    zIndex: 0,
                                    '&:hover': {
                                        borderWidth: "2px",
                                        borderColor: grey[900]
                                    },
                                    '&:focus': {
                                        borderWidth: "10px"
                                    }}}>
                                    <input type="date" id="birthday" 
                                        {...register('birthday')}
                                        placeholder="день рождения..." />
                                    <fieldset className={style.customFieldset}>
                                        <legend style={{
                                            fontSize: '0.75em',
                                            height: 11,
                                            padding: 0,
                                            visibility: 'hidden'
                                        }}>
                                            <span style={{
                                                paddingLeft: 5,
                                                paddingRight: 5,
                                                display: 'inline-block',
                                                opacity: 0,
                                                visibility: 'visible'
                                            }}>
                                                Дата рождения
                                            </span>
                                        </legend>
                                    </fieldset>
                                </Box>
                                <label style={{
                                    position: 'absolute',
                                    color: "rgba(0, 0, 0, 0.6)",
                                    fontWeight: 400,
                                    fontSize: '1rem',
                                    lineHeight: '1.4375em',
                                    letterSpacing: '0.00938em',
                                    padding: 0,
                                    display: 'block',
                                    whiteSpace: 'nowrap',
                                    overflow: 'hidden',
                                    textOverflow: 'ellipsis',
                                    left: 0,
                                    top: 0,
                                    transformOrigin: "top left",
                                    transform: "translate(14px, -9px) scale(0.75)",
                                    fontFamily: "Roboto, Helvetica, Arial, sans-serif",
                                    maxWidth: "calc(133% - 32px)"
                                }}>
                                    Дата рождения
                                </label>
                        </Box>
                    </Box>
                    {/* <Box sx={{
                        margin: 1
                    }}>
                         <FormControl>
                            <Controller render={
                                ({ field: { onChange, value }}) => (
                                    <>
                                        <label style={{color: grey[700], fontSize: 12, padding: 5}}>Уровень</label>
                                        <Select
                                            value={value}
                                            onChange={onChange}
                                            >
                                                <MenuItem value={0}>0</MenuItem>
                                                <MenuItem value={1}>1</MenuItem>
                                                <MenuItem value={2}>2</MenuItem>
                                                <MenuItem value={3}>3</MenuItem>
                                                <MenuItem value={4}>4</MenuItem>
                                                <MenuItem value={5}>5</MenuItem>
                                        </Select>
                                    </>
                                )
                            }
                            control={control}
                            name="grade"
                            />
                        </FormControl>
                    </Box> */}
                    <Box sx={{
                        margin: 1
                    }}>
                        <FormControl>
                            <Controller render={
                                ({ field: { onChange, value }}) => (
                                    <>
                                        <label style={{color: grey[700], fontSize: 12, padding: 5}}>Пол</label>
                                        <Select
                                            value={value}
                                            onChange={onChange}
                                            >
                                                <MenuItem sx={{display: 'flex', justifyContent: 'center'}} value="">Не выбрано</MenuItem>
                                                <MenuItem sx={{display: 'flex', justifyContent: 'center'}} value={"М"}>
                                                    <Tooltip title="Мужской" arrow placement="left">
                                                        <MaleIcon sx={{color: blue[700], display: 'block'}}/>
                                                    </Tooltip>
                                                </MenuItem>
                                                <MenuItem sx={{display: 'flex', justifyContent: 'center'}} value={"Ж"}>
                                                    <Tooltip title="Женский" arrow placement="left">
                                                        <FemaleIcon sx={{color: red[300], display: 'block'}}/>
                                                    </Tooltip>
                                                </MenuItem>
                                        </Select>
                                    </>
                                )
                            }
                            control={control}
                            name="gender"
                            />
                        </FormControl>
                    </Box>
                    <Box sx={{
                        margin: 1,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center'
                    }}>
                        <FormControl sx={{
                            display: 'flex',
                            flexDirection: 'row',
                            alignItems: 'center'
                        }}>
                            <Controller render={
                                ({ field: { onChange, value }}) => (
                                    <>
                                        <Tooltip title="У этого пользователя есть права администратора?" arrow placement="left">
                                            <SupervisorAccountIcon sx={{color: value ? blue[600] : grey[600]}}/>
                                        </Tooltip>
                                        <Switch id="is_admin" 
                                            checked={value}
                                            onChange={(_, checked) => {
                                            onChange(checked)}}
                                        />
                                    </>
                                )
                            } 
                            defaultValue={userData.is_admin}
                            control={control}
                            name="is_admin"
                            />
                        </FormControl>
                    </Box>
                    <Button style={{margin: 10}} type='submit' variant='outlined'>Обновить</Button>
                </form>
            </>}
        </>
    )
}