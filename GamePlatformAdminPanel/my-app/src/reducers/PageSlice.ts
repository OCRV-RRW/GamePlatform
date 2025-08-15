import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { store } from "../app/store";
import { PageState, PageStatus } from "../app/states_interfaces";

const initialState : PageState = {
    current_page: "",
    status: 'ok'
}

const pageSlice = createSlice({
    name: 'page',
    initialState, 
    reducers: {
        set_page(state, action: PayloadAction<string>) {
            state.current_page = action.payload
        },
        set_status(state, action: PayloadAction<PageStatus>) {
            state.status = action.payload
        }
    }
})

export const selectCurrentPage = () => store.getState().page.current_page
export const selectStatusPage = () => store.getState().page.status
export const { set_page, set_status } = pageSlice.actions
export default pageSlice.reducer 