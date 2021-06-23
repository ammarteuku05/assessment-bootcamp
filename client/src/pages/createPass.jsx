import React from 'react'

import Navbar from "../component/navbar"
import { useDispatch, useSelector } from 'react-redux'
import { useState } from 'react'

import { useHistory } from "react-router-dom"
import {createPass, fetchPass} from "../store/action/passAction"

function CreatePass() {
    const history = useHistory()
    const dispatch = useDispatch()

    const [website, setWebsite] = useState("")
    const [pass, setPass] = useState("")

    const { error } = useSelector(state => state.user)

    const submitCreatePass = (e) => {
        e.preventDefault()
        const data = {
            Website : website,
            Pass : pass,
        }

        console.log(data)

        
        if (!error) {
            dispatch(createPass(data))
            dispatch(fetchPass)
            history.push("/pass")
        }

    }
    // const back = () => {
    //     history.push("/pass")
    // }



    return (
        <div>
            <Navbar/>
            <div className="container">    
            <form style={{textAlign:"center", paddingTop:"150px", paddingBottom:"150px"}} onSubmit={submitCreatePass}>
            <h2>Create New Password</h2>
                <div className="mb-3">
                    <label for="title" className="form-label">Website</label>
                    <input type="text" className="form-control" id="title" onChange={e => {
                        e.preventDefault()
                        setWebsite(e.target.value)
                    }}/>
                </div>
                <div className="mb-3">
                    <label for="auhtor" className="form-label">Password</label>
                    <input type="text" className="form-control" id="auhtor" onChange={e => {
                        e.preventDefault()
                        setPass(e.target.value)
                    }}/>
                </div>
                <button type="submit" className="btn btn-primary" style={{ margin : "0px 100px"}}>Create</button>
            </form>
            </div> 
            {/* <Footer/> */}
        </div>
    )
}


export default CreatePass