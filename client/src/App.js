import React from "react"
import RegisterPage from "./pages/register"
import LoginPage from "./pages/login";
import TablePass from "./pages/pass";
import CreatePass from "./pages/createPass";
import UpdatePass from "./pages/updatePass";

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"; 

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/register">
          <RegisterPage/>
        </Route>
        <Route path="/login">
          <LoginPage/>
        </Route>
        <Route path="/pass">
          <TablePass/>
        </Route>
        <Route path="/createPass">
          <CreatePass/>
        </Route>
        <Route path="/updatePass/:pass_id">
          <UpdatePass/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
