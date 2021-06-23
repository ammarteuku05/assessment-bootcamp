import Axios from "axios"

const apiBase = Axios.create({
    baseURL : "http://localhost:8080"
})

export default apiBase