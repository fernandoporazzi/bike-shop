import {
  Link
} from "react-router-dom";

import { StyledNav, StyledSidebar, StyledTitle } from "./Sidebar.styled";

const Sidebar = (): JSX.Element => {
  return (
    <StyledSidebar>
      <StyledTitle>The Bike Shop</StyledTitle>
      <StyledNav>
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/add">Add a new bike</Link></li>
          <li><Link to="/about">About</Link></li>
        </ul>
      </StyledNav>
    </StyledSidebar>
  );
};

export default Sidebar;
