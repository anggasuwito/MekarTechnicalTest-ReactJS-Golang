import React from 'react'
import { Link } from 'react-router-dom'
import { Navbar, Nav } from 'react-bootstrap'

const MainNavbar = (props) => {
    const { onLogout, auth } = props
    let showPage
    if (auth) {
        showPage = <Navbar bg="light" expand="lg">
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="mr-auto">
                    <Link to={"/home"}><button className="btn navbar-btn" >Home</button></Link>
                    <Link to={"/user"}><button className="btn navbar-btn" >User</button></Link>
                    <Link><button className="btn navbar-btn" onClick={() => onLogout()}>Logout</button></Link>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
    }
    return (
        <div>
            {showPage}
        </div>
    );
}

export default MainNavbar
