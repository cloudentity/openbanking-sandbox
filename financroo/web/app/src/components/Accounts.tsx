import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import BankCard from "./BankCard";
import mainClasses from "./main.module.css"

const useStyles = makeStyles((theme: Theme) => ({}));

export default ({accounts, onConnectClick}) => {
  const classes = useStyles();

  return (
    <div style={{height: '100%', display: 'flex', flexDirection: 'column'}}>
      <Typography className={mainClasses.sectionTitle}>All accounts</Typography>

      {accounts.map(account => (
        <BankCard account={account} selected={true} onToggleSelect={() => {}} style={{marginBottom: 32}}/>
      ))}

      <div style={{flex: 1}}/>

      <Button color={'primary'} variant={'contained'} size={'large'} style={{width: '100%'}} onClick={onConnectClick}>Connect your bank</Button>
    </div>
  )
};
