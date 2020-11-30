import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import griffinLogo from "../assets/griffin-logo.svg";
import AppBar from "@material-ui/core/AppBar";

export default function PageToolbar({children}) {

  return (
    <AppBar position="fixed" variant={'outlined'}>
      <Toolbar>
        <img alt="financroo logo" src={griffinLogo} style={{marginRight: 32}}/>
        {children}
      </Toolbar>
    </AppBar>
  )
};
