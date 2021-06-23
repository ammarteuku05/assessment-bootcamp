
import api from "../../APIs/api"

const access_token = !localStorage.getItem("access_token") ? "" : localStorage.getItem("access_token")

export const fetchPass = () => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const { data } = await api({
                method : "GET",
                url : "/pass",
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(data)

           return dispatch({ type : "FETCH_PASS", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }
}


export const createPass = (payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const { data } =  await api({
                method : "POST",
                url : "/pass",
                data : payload,
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(data)

            return dispatch({ type : "CREATE_PASS", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }

}


export const updatePass = (id, payload) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const { data } = await api({
                method : "PUT",
                url : `/pass/${id}`,
                data : payload,
                headers : {
                    "authorization" : access_token
                }
            })

            console.log(data)

            return dispatch({ type : "UPDATE_PASS", payload : data})

        } catch(err) {
            dispatch({ type : "ERROR_PASS"})
            console.log(err.response.data)
        }
    }
}


export const deletePass = (id, history) => {
    return async (dispatch) => {
        try {
            dispatch({type : "PASS_LOADING"})

            const { data } =  await api({
                method : "DELETE",
                url : `/pass/${id}`,
                headers : {
                    "authorization" : access_token
                }
            })
            history.push("/pass")

            console.log(data)

           return dispatch({ type : "DELETE_PASS", payload : data})
        } catch(err) {
            dispatch({ type : "ERROR_BOOKS"})
            console.log(err.response.data)
        }
    }
}
