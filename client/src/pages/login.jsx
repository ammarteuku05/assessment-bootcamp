import React, { useEffect, useState } from "react"

import Navbar from "../component/navbar"
import { useDispatch, useSelector } from "react-redux"

import { login } from '../store/action/userAction'
import { useHistory } from "react-router-dom"
import { fetchPass } from "../store/action/passAction"

function LoginPage() {

    const history = useHistory()

    const dispatch = useDispatch()
    const [email, setEmail] = useState("")
    const [pass, setPass] = useState("")

    const { user, error } = useSelector(state => state.user)

    useEffect(() => {
    }, [])

    const loginSubmit = () => {
        const data = {
            email : email,
            password :pass,
        }
        dispatch(login(data, history))

        if (!error) {
            dispatch(fetchPass(data))
            history.push("/pass")
        }
        // console.log(error)
    }

    return(
        <div>
        <Navbar/>
        <div className="container">    
            <form style={{textAlign:"center", paddingTop:"100px", paddingBottom:"30px"}} onSubmit={loginSubmit}>
            <h2>Login User</h2>
            <div className="mb-3">
                <label for="email" className="form-label">Email address</label>
                <input type="email" className="form-control" id="email" aria-describedby="emailHelp" onChange={e => {
                    e.preventDefault()
                    setEmail(e.target.value)
                }}/>
                <div id="emailHelp" className="form-text">We'll never share your email with anyone else.</div>
            </div>
            <div class="mb-3">
                <label for="password" className="form-label">Password</label>
                <input type="password" className="form-control" id="password" onChange={e => {
                    e.preventDefault()
                    setPass(e.target.value)
                }}/>
            </div>
            <button type="submit" className="btn btn-primary">Login</button>
            </form>
        </div>
        {/* <Footer/> */}
        </div>
    )
}


export default LoginPage