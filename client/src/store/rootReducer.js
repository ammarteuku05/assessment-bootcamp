import {combineReducers} from "redux"

import userReducer from "./reducer/userReducer"

const root = combineReducers({
    user:userReducer,
})
export default root