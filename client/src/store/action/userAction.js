import apiBase from '../../API/axios'

export const registerUser = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const { data } = await apiBase({
                method : "POST",
                url : "/users/register",
                data : payload,
            })

            console.log(data)

            return dispatch({ type : "USER_REGISTER", payload : data})

        } catch(err) {
            // err?.response?.data ? dispatch({ type : "USER_ERROR", payload: err.response.data}) : console.log(err.response.data)
            console.log(err.response)
        }
    }
}

export const loginUser = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({ type : "USER_LOADING"})
    
            const { data } = await apiBase({
                method : "POST",
                url : "/users/login",
                data : payload,
            })

            localStorage.setItem("access_token", data.authorization)

            console.log(data)

            return dispatch({ type : "USER_LOGIN", payload : data})

        } catch(err) {
            // err?.response?.data ? dispatch({ type : "USER_ERROR", payload: err.response.data}) : console.log(err.response.data)
        
            console.log(err.response)
        }
    }
}

export const logoutUser = () => ({
    type : "LOGOUT_USER"
})
