const initState = {
    user : null,
    isLoading : false,
    error : null,
}

const useReducer = (state = initState, action) => {
    switch (action.type) {
        case "USER_LOADING":
            return { ...state, isLoading: true}
        case "USER_REGISTER":
            return { ...state, user : action.payload, isLoading : false }
        case "USER_LOGIN": 
            localStorage.setItem("access_token", action.payload.authorization)
            localStorage.setItem("accessToken", action.payload.authorization)
            return { ...state, user: action.payload, isLoading: false }
        case "LOGOUT_USER":
            localStorage.removeItem("access_token")
            return { ...state, user : null}
        case "USER_ERROR":
            return { ...state, error : action.payload, isLoading: false}
        default:
            return state
    }
}

export default useReducer