import Sidebar from "./components/Sidebar";
import Home from "./pages/Home";
import About from "./pages/About";
import Add from "./pages/Add";
import Bike from "./pages/Bike";

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import { StyledApp, StyledMain } from "./App.styled";

function App(): JSX.Element {
  return (
    <Router>
      <StyledApp>
        <Sidebar />

        <StyledMain>
          <Switch>
            <Route path="/about">
              <About />
            </Route>
            <Route path="/add">
              <Add />
            </Route>
            <Route path="/bikes/:id">
              <Bike />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </StyledMain>
      </StyledApp>
    </Router>
  );
}

export default App;
