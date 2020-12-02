import React, {useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Connected from "./Connected";
import Welcome from "./Welcome";
import ConnectAccount from "./ConnectAccount";
import {useQuery} from "react-query";
import {api} from "../api/api";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {Button} from "@material-ui/core";
import {logout} from "./AuthPage";
import PageContent from "./PageContent";
import PageContainer from "./PageContainer";
import {pathOr} from "ramda";

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
  const [connectAccountOpen, setConnectAccountOpen] = useState(false);
  const [connectProgress, setConnectProgress] = useState(false);

  const {
    isLoading: fetchAccountsProgress,
    isError: fetchAccountsError,
    data: accountsRes
  } = useQuery('fetchAccounts', api.fetchAccounts, {refetchOnWindowFocus: false, retry: false});

  if (fetchAccountsError) {
    //
  }

  const accounts = accountsRes ? pathOr([], ['accounts'], accountsRes) : []

  const {isLoading: fetchBalancesProgress, data: balancesRes} = useQuery('fetchBalances', api.fetchBalances);

  const handleAllowAccess = ({permissions}) => {
    setConnectProgress(true);
    api.connectBank({permissions})
      .then(res => {
        window.location.href = res.login_url;
      })
      .catch(() => setConnectProgress(false));
  };

  const isProgress = fetchAccountsProgress || fetchBalancesProgress || connectProgress;

  return (
    <div style={{position: 'relative'}}>
      <PageToolbar>
        <Hidden mdUp>
          <IconButton edge="start" color="inherit" aria-label="menu">
            <MenuIcon/>
          </IconButton>
        </Hidden>
        <Hidden smDown>
          <Tabs
            value={'accounts'}
            indicatorColor="primary"
            onChange={() => {
            }}
            aria-label="menu tabs"
            style={{height: 64}}
          >
            <Tab label="Accounts" value={'accounts'} style={{height: 64}}/>
            <Tab label="Spending" value={'spending'} style={{height: 64}}/>
            <Tab label="Settings" value={'settings'} style={{height: 64}}/>
          </Tabs>
        </Hidden>
        <Button variant={"outlined"} onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
      </PageToolbar>
      {isProgress && <Progress/>}

      {!isProgress && (
        <>
          {accounts.length === 0 && (
            <PageContainer withBackground>
              <Welcome onConnectClick={() => setConnectAccountOpen(true)}/>
            </PageContainer>
          )}
          {accounts.length > 0 && (
            <PageContent>
              <Connected accounts={accounts} balances={balancesRes.balances}
                         onConnectClick={() => setConnectAccountOpen(true)}/>
            </PageContent>
          )}
        </>
      )}

      {connectAccountOpen && <ConnectAccount onAllowAccess={handleAllowAccess} onClose={() => setConnectAccountOpen(false)}/>}
    </div>
  )
};
