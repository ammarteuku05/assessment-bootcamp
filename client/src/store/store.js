import thunk from 'redux-thunk'

import { applyMiddleware, createStore } from 'redux'

import root from './rootReducer'

const store =  createStore(root, applyMiddleware(thunk))

export default store