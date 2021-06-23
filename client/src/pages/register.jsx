import React, { useState } from "react"

import Navbar from "../component/navbar"
import Footer from "../component/footer"
import { useDispatch } from "react-redux"

import { registerUser } from '../store/action/userAction'

function RegisterPage() {

    const dispatch = useDispatch()

    const [name, setName] = useState("")
    const [address, setAddress] = useState("")
    const [email, setEmail] = useState("")
    const [pass, setPass] = useState("")

    const registerSubmit = () => {
        const data = {
            name : name, 
            address : address,
            email : email,
            password :pass,
        }

        console.log(data)

        dispatch(registerUser(data))
    }

    return(
        <div>
        <Navbar/>
        <div className="container">    
            <form style={{textAlign:"center", paddingTop:"100px", paddingBottom:"30px"}} onSubmit={registerSubmit}>
            <h2>Register User</h2>
            <div className="mb-3">
                <label for="name" className="form-label">Name</label>
                <input type="text" className="form-control" id="name" onChange={e => {
                    e.preventDefault()
                    setName(e.target.value)
                }}/>
            </div>
            <div className="mb-3">
                <label for="adress" className="form-label">Address</label>
                <input type="text" className="form-control" id="address" onChange={e => {
                    e.preventDefault()
                    setAddress(e.target.value)
                }}/>
            </div>
            <div className="mb-3">
                <label for="date-birth" className="form-label">Date birth</label>
                <input type="text" className="form-control" id="date-birth" onChange={e => {
                    e.preventDefault()
                    setDate(e.target.value)
                }}/>
            </div>
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
            <button type="submit" className="btn btn-primary">Register</button>
            </form>
        </div>
        <Footer/>
        </div>
    )
}


export default RegisterPage