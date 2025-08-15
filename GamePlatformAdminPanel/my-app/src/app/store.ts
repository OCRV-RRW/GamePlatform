import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import userReducer from '../reducers/UserSlice';
import counterReducer from '../features/counter/counterSlice'
import pageReducer from '../reducers/PageSlice';

export const store = configureStore({
  reducer: {
    user: userReducer,
    page: pageReducer,
    counter: counterReducer
  }
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;