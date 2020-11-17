import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import BankCard from "./BankCard";
import mainClasses from "./main.module.css"
import {groupBy, keys, pipe} from "ramda";
import Card from "@material-ui/core/Card";

const useStyles = makeStyles((theme: Theme) => ({}));

export default ({accounts, balances, filtering, onChangeFiltering, onConnectClick}) => {
  const classes = useStyles();

  const banks = pipe(groupBy((a: any) => a.BankID), keys)(accounts);
  const totalBalance = balances.reduce((total, b) => total + parseFloat(b.Amount.Amount), 0).toFixed(2);

  return (
    <div style={{height: '100%', display: 'flex', flexDirection: 'column'}}>
      <Card style={{padding: '32px 20px'}}>
        <div style={{display: 'flex', alignItems: 'center', justifyContent: 'space-between'}}>
          <Typography className={mainClasses.sectionTitle}>All accounts</Typography>
          <Typography><strong>GBP {totalBalance}</strong></Typography>
        </div>
        <Typography
          style={{display: 'inline-block', background: 'rgba(54, 198, 175, 0.08)', color: '#36C6AF', fontSize: 14, padding: 2, marginTop: 4}}
        >{accounts.length} accounts synced</Typography>
      </Card>

      {banks.map(bankId => (
        <BankCard
          bankId={bankId}
          accounts={accounts.filter(a => a.BankID === bankId)}
          balances={balances}
          filtering={filtering}
          onChangeFiltering={onChangeFiltering}
          style={{marginTop: 32}}/>
      ))}

      <div style={{flex: 1}}/>

      <Button color={'secondary'} variant={'contained'} size={'large'} style={{width: '100%'}} onClick={onConnectClick}>Connect your bank</Button>
    </div>
  )
};
