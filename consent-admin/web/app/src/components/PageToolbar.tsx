import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import griffinLogo from "../assets/griffin-logo.svg";
import AppBar from "@material-ui/core/AppBar";
import {Typography} from "@material-ui/core";

export default function PageToolbar({children}) {

    return (
        <AppBar position="fixed" variant={'outlined'}>
            <Toolbar>
                <div>
                    <img alt="app logo" src={griffinLogo} style={{marginRight: 32, width: 100}}/>
                    <Typography style={{fontSize: 12, marginLeft: 24}}>Powered by <strong>Axway</strong></Typography>
                </div>
                {children}
            </Toolbar>
        </AppBar>
    )
};
