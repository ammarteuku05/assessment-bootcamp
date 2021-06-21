import React, { useState } from "react"
import Navbar from "../component/navbar"
import {register} from '../store/action/userAction'

function Resister(){
   const dispatch = useDispatch()

   const[fullname, setFullName]= useState("")
   const[address, setAddress]= useState("")
   const[email, setEmail]= useState("")
   const[pass, setPass]= useState("")

   const registerSubmit = () =>{
       const data = {
           fullname = fullname,
           address = address,
           email = email, 
           pass = pass,
       }

       dispatch(register(data))
   }
     
    return(
        <>
        <Navbar/>
        <Form>
  <Form.Group className="mb-3" controlId="formNmae">
    <Form.Label>Full Name</Form.Label>
    <Form.Control type="name" placeholder="Enter Full Name" onchange={e=>{
        e.preventDefault()
        setFullName(e.target.value)
    }}/>
  </Form.Group>
  <Form.Group className="mb-3" controlId="formBasicAddress">
    <Form.Label>Address</Form.Label>
    <Form.Control type="email" placeholder="Enter Address" onchange={e=>{
        e.preventDefault()
        setAddress(e.target.value)
    }}/>
  </Form.Group>
  <Form.Group className="mb-3" controlId="formBasicEmail">
    <Form.Label>Email address</Form.Label>
    <Form.Control type="email" placeholder="Enter email" onchange={e=>{
        e.preventDefault()
        setEmail(e.target.value)
    }}/>
    <Form.Text className="text-muted">
      We'll never share your email with anyone else.
    </Form.Text>
  </Form.Group>

  <Form.Group className="mb-3" controlId="formBasicPassword">
    <Form.Label>Password</Form.Label>
    <Form.Control type="password" placeholder="Password" onchange={e=>{
        e.preventDefault()
        setPass(e.target.value)
    }}/>
  </Form.Group>
  <Button variant="primary" type="submit">
    Submit
  </Button>
</Form>
</>
    )
}

export default Resister