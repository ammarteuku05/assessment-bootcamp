import { combineReducers } from 'redux'
import passReducer from './reducer/passReducer'
import user from './reducer/userReducer';


const rootReducer = combineReducers({
    user : user,
    pass : passReducer,
})

export default rootReducer;

