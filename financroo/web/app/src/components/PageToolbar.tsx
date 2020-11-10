import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";
import financrooLogo from "../assets/financroo-logo.svg";
import AppBar from "@material-ui/core/AppBar";

const useStyles = makeStyles((theme: Theme) =>{});

export default ({children}) => {
  const classes = useStyles();

  return (
    <AppBar position="fixed" color={'inherit'} variant={'outlined'}>
      <Toolbar>
        <img src={financrooLogo}/>
        <div style={{flex: 1}}/>
        {children}
      </Toolbar>
    </AppBar>
  )
};
