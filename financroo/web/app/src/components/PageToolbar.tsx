import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import financrooLogo from "../assets/financroo-logo.svg";
import AppBar from "@material-ui/core/AppBar";

export default function PageToolbar({children}) {

  return (
    <AppBar position="fixed" color={'inherit'} variant={'outlined'}>
      <Toolbar>
        <img alt="financroo logo" src={financrooLogo}/>
        <div style={{flex: 1}}/>
        {children}
      </Toolbar>
    </AppBar>
  )
};
