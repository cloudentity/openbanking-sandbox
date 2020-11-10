import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Accounts from "./Accounts";
import Analytics from "./Analytics";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    marginTop: 56,
    height: 'calc(100vh - 56px)',
    [theme.breakpoints.up('sm')]: {
      marginTop: 64,
      height: 'calc(100vh - 64px)',
    },
  }
}));

export default ({accounts, onConnectClick}) => {
  const classes = useStyles();

  return (
    <Grid container className={classes.root}>
      <Grid item xs={4} style={{background: '#F7FAFF', padding: '16px 32px', borderRight: '1px solid #EAECF1'}}>
        <Accounts accounts={accounts} onConnectClick={onConnectClick}/>
      </Grid>
      <Grid item xs={8} style={{background: '#FCFCFF', padding: '16px 32px'}}>
        <Analytics/>
      </Grid>
    </Grid>
  )
};
