import { combineReducers } from 'redux'

import userReducer from './reducer/userReducer'
// import bookReducer from './reducer/bookReducer'


const rootReducer = combineReducers({
    user : userReducer,
    // book : bookReducer,
})

export default rootReducer