import api from "../../API/api"

export const registerUser = (payload) =>{
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