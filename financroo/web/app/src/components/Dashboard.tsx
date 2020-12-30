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
  const [isProgress, setProgress] = useState(false);

  const {
    isLoading: fetchBanksProgress, data: banksRes, refetch: refetchBanks
  } = useQuery('fetchBanks', api.fetchBanks, {refetchOnWindowFocus: false, retry: false});

  const banks = banksRes ? pathOr([], ['connected_banks'], banksRes) : []

  const handleAllowAccess = ({bankId, permissions}) => {
    setProgress(true);
    api.connectBank(bankId, {permissions})
      .then(res => {
        window.location.href = res.login_url;
      })
      .catch(() => setProgress(false));
  };

  const handleDisconnectBank = bankId => () => {
    setProgress(true);
    api.disconnectBank(bankId)
      .then(refetchBanks)
      .finally(() => setProgress(false));
  }

  const handleReconnectBank = (bankId, permissions) => () => {
    setProgress(true);
    api.connectBank(bankId, {permissions})
      .then(res => {
        window.location.href = res.login_url;
      })
      .catch(() => setProgress(false));
  }

  const showProgress = isProgress || fetchBanksProgress;

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
      {showProgress && <Progress/>}

      {!showProgress && (
        <>
          {banks.length === 0 && (
            <PageContainer withBackground>
              <Welcome onConnectClick={() => setConnectAccountOpen(true)}/>
            </PageContainer>
          )}
          {banks.length > 0 && (
            <PageContent>
              <Connected
                banks={banks}
                onConnectClick={() => setConnectAccountOpen(true)}
                onDisconnect={handleDisconnectBank}
                onReconnect={handleReconnectBank}
              />
            </PageContent>
          )}
        </>
      )}

      {connectAccountOpen && <ConnectAccount connected={banks} onAllowAccess={handleAllowAccess} onClose={() => setConnectAccountOpen(false)}/>}
    </div>
  )
};
