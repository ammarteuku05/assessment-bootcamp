const initialState = {
    user:null,
    isLoading:false,
    error: null,
}

const userReducer = (state = initialState, action) =>{
    switch (action.type) {
        case "USER_LOADING":
            return {...state, isLoading:true}
        case "USER_REGISTER":
            return{...state, user:action.payload, isLoading:false}
        case "USER_LOGIN":
            localStorage.setItem("accessss_token", action.payload.authorization)
            localStorage.setItem("accessToken", action.payload.authorization)
            return{...state, user:action.payload, isLoading:false}
        case "LOGOUT_USER":
            localStorage.removeItem("access_token")
            return{...state, user:null}
    }
}

export default userReducer