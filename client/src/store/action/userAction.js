import api from "../../APIs/api"

export const register = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const { data } = await api({
                method : "POST",
                url : "/users/register",
                data : payload,
            })

            console.log(data)

            return dispatch({ type : "USER_REGISTER", payload : data})

        } catch(err) {
            console.log(err.response)
        }
    }
}

export const login = (payload, history) => {
    
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const { data } = await api({
                method : "POST",
                url : "/users/login",
                data : payload,
            })

            localStorage.setItem("access_token", data.data.authorization)

            // console.log("data",data.data.authorization)
            console.log(payload)

            history.push("/pass")

            return dispatch({ type : "USER_LOGIN", payload : data})


        } catch(err) {
            console.log(payload)
            dispatch({type : "USER_ERROR", payload : err})
            console.log(err.response)
        }
    }
}

export const logoutUser = () => ({
    type : "LOGOUT_USER"
})
