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

export default ({userinfo}) => {
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
        <Button variant={"outlined"} onClick={() => logout()}>Logout</Button>
      </PageToolbar>
      {isProgress && <Progress/>}

      {!isProgress && (
        <>
          {data && data.accounts.length === 0 && <Welcome onConnectClick={() => setConnectAccountOpen(true)}/>}
          {data && data.accounts.length > 0 && <Connected accounts={data.accounts} onConnectClick={() => setConnectAccountOpen(true)}/>}
        </>
      )}


      {connectAccountOpen && <ConnectAccount onAllowAccess={handleAllowAccess} onClose={() => setConnectAccountOpen(false)}/>}
    </div>
  )
};
