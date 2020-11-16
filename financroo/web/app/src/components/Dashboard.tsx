import React, {useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Connected from "./Connected";
import Welcome from "./Welcome";
import ConnectAccount from "./ConnectAccount";
import {useMutation, useQuery, useQueryCache} from "react-query";
import {api} from "../api/api";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {Button} from "@material-ui/core";
import {logout} from "./AuthPage";
import PageContent from "./PageContent";
import PageContainer from "./PageContainer";

export default ({authorizationServerURL, authorizationServerId, tenantId}) => {
  const [connectAccountOpen, setConnectAccountOpen] = useState(false);

  const queryCache = useQueryCache();
  const {isLoading: fetchProgress, data} = useQuery('fetchAccounts', api.fetchAccounts)
  const [updateAccounts, {isLoading: updateProgress}] = useMutation(api.updateAccounts, {
    onSuccess: (data, variables) => {
      queryCache.setQueryData('fetchAccounts', data);
    }
  })

  const handleAllowAccess = async (bank) => {
    try {
      setConnectAccountOpen(false);
      await updateAccounts({accounts: [...data.accounts, bank]});
    } catch {
      //
    }
  };

  const isProgress = fetchProgress || updateProgress;

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
          {data && data.Data.Account.length === 0 && (
            <PageContainer withBackground>
              <Welcome onConnectClick={() => setConnectAccountOpen(true)}/>
            </PageContainer>
          )}
          {data && data.Data.Account.length > 0 && (
            <PageContent>
              <Connected accounts={data.Data.Account} onConnectClick={() => setConnectAccountOpen(true)}/>
            </PageContent>
          )}
        </>
      )}

      {connectAccountOpen && <ConnectAccount onAllowAccess={handleAllowAccess} onClose={() => setConnectAccountOpen(false)}/>}
    </div>
  )
};
