import React from "react"

const Navbar = () =>{
    return(
        <Navbar>
  <Container>
    <Navbar.Brand href="#home">Navbar with text</Navbar.Brand>
    <Navbar.Toggle />
    <Navbar.Collapse className="justify-content-end">
        <Nav.Link href="#home">Login</Nav.Link>
        <Nav.Link href="#link"></Nav.Link>
    </Navbar.Collapse>
  </Container>
</Navbar>
    )
}

export default Navbar