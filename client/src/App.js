import React from "react"
import HomePage from "./page/homePage";
import Resister from "./page/register";
import Login from "./page/login";
import Password from "./page/password";

import {BrowserRouter as Router, Switch, Route} from "react-router-dom";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact>
          <HomePage />
        </Route>
        <Route path="/register" exact>
          <Resister />
        </Route>
        <Route path="/login" exact>
          <Login />
        </Route>
        <Route path="/password" exact>
          <Password />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
