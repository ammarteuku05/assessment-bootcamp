const initState = {
    passw : [],
    pass : null,
    isLoading : false,
    error : null
}

const passReducer = (state = initState, action) => {
    switch(action.type) {
        case "PASS_LOADING":
            return { ...state, isLoading : true}
        case "FETCH_PASS":
            return { ...state, isLoading: false, passw : action.payload}
        case "CREATE_PASS":
            return { ...state, isLoading:false, pass : action.payload}
        case "UPDATE_PASS":
            return { ...state, isLoading:false, pass : action.payload}
        case "DELETE_PASS":
            return { ...state, isLoading:false, pass : null}
        case "ERROR_PASS":
            return { ...state, isLoading: false}
        default:
            return state
    }
}

export default passReducer