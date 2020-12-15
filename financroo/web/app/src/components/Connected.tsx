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
import {pathOr, pick} from "ramda";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    height: '100%;'
  }
}));

export default function Connected({banks, onConnectClick, onDisconnect}) {
  const classes = useStyles();
  const [filtering, setFiltering] = useState({
    accounts: [],
    months: [],
    categories: []
  });

  const {
    isLoading: fetchAccountsProgress,
    isError: fetchAccountsError,
    data: accountsRes
  } = useQuery('fetchAccounts', api.fetchAccounts, {
    refetchOnWindowFocus: false,
    retry: false,
    onSuccess: data => {
      setFiltering(m => ({...m, accounts: (data.accounts || []).map(a => a.AccountId)}));
    }
  });

  if (fetchAccountsError) {
    //
  }

  const {isLoading: fetchBalancesProgress, data: balancesRes} = useQuery('fetchBalances', api.fetchBalances, {
    refetchOnWindowFocus: false,
    retry: false
  });

  const accounts = accountsRes ? pathOr([], ['accounts'], accountsRes) : [];
  const balances = balancesRes ? pathOr([], ['balances'], balancesRes) : [];

  const {isLoading: fetchTransactionsProgress, data} = useQuery('fetchTransactions', api.fetchTransactions, {
    refetchOnWindowFocus: false,
    retry: false
  });

  const isLoading = fetchAccountsProgress || fetchBalancesProgress || fetchTransactionsProgress;

  if (isLoading) {
    return <Progress/>;
  }

  const transactions = applyFiltering(pick(['accounts'], filtering), data?.transactions);

  return (
    <Grid container className={classes.root}>
      <Grid item xs={4} style={{background: '#F7FAFF', padding: '16px 32px', borderRight: '1px solid #EAECF1'}}>
        <Accounts
          banks={banks}
          accounts={accounts}
          balances={balances}
          filtering={filtering}
          onChangeFiltering={f => setFiltering({...filtering, ...f})}
          onConnectClick={onConnectClick}
          onDisconnect={onDisconnect}
        />
      </Grid>
      <Grid item xs={8} style={{background: '#FCFCFF', padding: '32px 32px 16px 32px'}}>
        <Analytics transactions={transactions} filtering={filtering} onChangeFiltering={f => setFiltering({...filtering, ...f})}/>
      </Grid>
    </Grid>
  )
};
