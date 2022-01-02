import React from 'react';
import {
  Nav,
  NavLink,
  Bars,
  NavMenu
} from './NavbarElements';

const Navbar = () => {
  return (
    <>
      <Nav>
      <NavLink to='/' activeStyle>
          Home
          </NavLink>
        <Bars />
        <NavMenu>
          <NavLink to='/top3' activeStyle>
            Top 3 Departamentos 
          </NavLink>
          <NavLink to='/graficacircular' activeStyle>
            Grafica Circular
          </NavLink>
          <NavLink to='/ultimos5' activeStyle>
          Ultimos 5 Casos
          </NavLink>
          <NavLink to='/graficabarras' activeStyle>
          Grafica Barras Edades
          </NavLink>
          <NavLink to='/dashboard' activeStyle>
          All
          </NavLink>
          
          {/* Second Nav */}
          {/* <NavBtnLink to='/sign-in'>Sign In</NavBtnLink> */}
        </NavMenu>
       
      </Nav>
    </>
  );
};

export default Navbar;
