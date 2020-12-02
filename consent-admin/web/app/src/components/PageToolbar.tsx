import React from "react";
import Toolbar from "@material-ui/core/Toolbar";
import AppBar from "@material-ui/core/AppBar";
import {Typography} from "@material-ui/core";

export default function PageToolbar({children}) {

    return (
        <AppBar position="fixed" variant={'outlined'}>
            <Toolbar>
                <div>
                    <Typography style={{fontSize: 12, marginLeft: 24}}>Powered by <strong>Cloudentity</strong></Typography>
                </div>
                {children}
            </Toolbar>
        </AppBar>
    )
};
