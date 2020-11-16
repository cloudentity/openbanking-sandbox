import React, {useState} from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Accounts from "./Accounts";
import Analytics from "./Analytics";
import {useQuery} from "react-query";
import {api} from "../api/api";
import Progress from "./Progress";
import {applyFiltering} from "./analytics.utils";
import {filter, pick} from "ramda";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    height: '100%;'
  }
}));

export default ({accounts, onConnectClick}) => {
  const classes = useStyles();
  const [filtering, setFiltering] = useState({
    accounts: accounts.map(a => a.AccountId),
    months: [],
    categories: []
  });
  const {isLoading, data} = useQuery('fetchTransactions', api.fetchTransactions);

  if (isLoading) {
    return <Progress/>;
  }

  const transactions = applyFiltering(pick(['accounts'], filtering), data.Data.Transaction);

  return (
    <Grid container className={classes.root}>
      <Grid item xs={4} style={{background: '#F7FAFF', padding: '16px 32px', borderRight: '1px solid #EAECF1'}}>
        <Accounts accounts={accounts} filtering={filtering} onChangeFiltering={f => setFiltering({...filtering, ...f})} onConnectClick={onConnectClick}/>
      </Grid>
      <Grid item xs={8} style={{background: '#FCFCFF', padding: '32px 32px 16px 32px'}}>
        <Analytics transactions={transactions} filtering={filtering} onChangeFiltering={f => setFiltering({...filtering, ...f})}/>
      </Grid>
    </Grid>
  )
};
