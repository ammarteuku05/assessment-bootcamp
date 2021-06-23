import React from 'react'

import Navbar from "../component/navbar"
import { useDispatch, useSelector } from 'react-redux'
import { useState, useEffect } from 'react'

import { useHistory, useLocation } from "react-router-dom"
import {fetchPass, updatePass} from "../store/action/passAction"

function UpdatePass() {
    const history = useHistory()
    const dispatch = useDispatch()

    const [website, setWebsite] = useState("")
    const [pass, setPass] = useState("")

    const { error } = useSelector(state => state.user)

    const location = useLocation()

    const passId = location.pathname.substr(location.pathname.lastIndexOf("/")+1)

    // useEffect(() => {
    //     console.log(passId)
    // }, [])


    const submitUpdatePass = (e) => {
        e.preventDefault()
        const data = {
            Website : website,
            Pass : pass,
        }

        console.log(data)

        dispatch(updatePass(passId,data))

        if (!error) {
            dispatch(updatePass(passId,data))
            dispatch(fetchPass())
            history.push("/pass")
        }

    }
    const back = () => {
        dispatch(fetchPass())
        history.push("/pass")
    }


    return (
        <div>
            <Navbar/>
            <div className="container">    
            <form style={{textAlign:"center", paddingTop:"150px", paddingBottom:"150px"}} onSubmit={submitUpdatePass}>
            <h2>Create New Password</h2>
                <div className="mb-3">
                    <label for="title" className="form-label">Title</label>
                    <input type="text" className="form-control" id="title" onChange={e => {
                        e.preventDefault()
                        setWebsite(e.target.value)
                    }}/>
                </div>
                <div className="mb-3">
                    <label for="auhtor" className="form-label">Auhtor</label>
                    <input type="text" className="form-control" id="auhtor" onChange={e => {
                        e.preventDefault()
                        setPass(e.target.value)
                    }}/>
                </div>
                <button type="submit" className="btn btn-danger" style={{ margin : "0px 100px"}} onClick={back}>Back</button>
                <button type="submit" className="btn btn-primary" style={{ margin : "0px 100px"}}>Create</button>
            </form>
            </div> 
            {/* <Footer/> */}
        </div>
    )
}


export default UpdatePass