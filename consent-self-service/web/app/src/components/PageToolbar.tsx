import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";

export default function PageToolbar({children}) {

  return (
    <AppBar position="fixed" variant={'outlined'}>
      <Toolbar>
        {children}
      </Toolbar>
    </AppBar>
  )
};
