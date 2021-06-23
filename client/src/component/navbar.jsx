import React, {useState, useEffect} from "react";
import { useDispatch } from "react-redux";
import { logoutUser } from "../store/action/userAction";
import { useHistory, useLocation } from "react-router-dom";
import {Navbar, Nav} from "react-bootstrap"

// const Navbar = () => {

//   const dispatch = useDispatch()
//   const history = useHistory()

//   const handleLogout = () => {
//     dispatch(logoutUser());
//     history.push("/login");
//   };
//   return (
//     <>
//       <nav class="navbar navbar-expand-lg navbar-light bg-white">
//         <div class="container">
//           <a class="navbar-brand" href="/">
//             Password
//           </a>
//           <button
//             class="navbar-toggler"
//             type="button"
//             data-bs-toggle="collapse"
//             data-bs-target="#navbarNav"
//             aria-controls="navbarNav"
//             aria-expanded="false"
//             aria-label="Toggle navigation"
//           >
//             <span class="navbar-toggler-icon"></span>
//           </button>

//           <div
//             class="collapse navbar-collapse justify-content-end"
//             id="navbarNav"
//           >
//             <ul class="navbar-nav">
//               <li class="Login">
//                 <a class="nav-link active" aria-current="page" href="/">
//                   Login
//                 </a>
//               </li>
//               <li class="nav-item">
//                 <a class="nav-link active" aria-current="page" href="/">
//                   Register
//                 </a>
//               </li>
//               <li class="nav-item">

//                 <a class="nav-link active" aria-current="page" onClick={handleLogout}>
//                   Logout
//                 </a>
//               </li>
//             </ul>
//           </div>
//         </div>
//       </nav>


//     </>
//   );
// };

// export default Navbar;

const Header = () => {
  const [pageURL, setPageURL] = useState("");
  const authUser = !!localStorage.getItem("access_token");
  const location = useLocation();
  const dispatch = useDispatch();
  const history = useHistory();

  useEffect(() => {
    setPageURL(location.pathname);
  }, []);

  const handleLogout = () => {
    dispatch(logoutUser());
    history.push("/login");
  };

  return (
    <>
      <Navbar collapseOnSelect expand="lg" bg="light" variant="light">
        <div className="container">

        <Navbar.Brand href="/" className="ml-5">
          Password Manager
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse
          id="responsive-navbar-nav"
          className="justify-content-end"
        >
          <Nav>
            {authUser && (
              <>
                <Nav.Link className="mr-5" onClick={handleLogout}>
                  Logout
                </Nav.Link>
              </>
            )}
            {!authUser && (
              <>
                {pageURL === "/register" || (
                  <Nav.Link href="/register" className="mr-5">
                    Register
                  </Nav.Link>
                )}
                {pageURL === "/login" || (
                  <Nav.Link href="/login" className="mr-5">
                    Login
                  </Nav.Link>
                )}
              </>
            )}
          </Nav>
        </Navbar.Collapse>
        </div>
      </Navbar>
    </>
  );
};

export default Header;