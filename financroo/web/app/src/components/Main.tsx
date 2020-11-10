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

export default ({}) => {
  const [connectAccountOpen, setConnectAccountOpen] = useState(false);

  const {isLoading, data} = useQuery('accounts', api.fetchAccounts)

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
      </PageToolbar>
      {isLoading && <Progress/>}

      {data && data.accounts.length === 0 && <Welcome onConnectClick={() => setConnectAccountOpen(true)}/>}
      {data && data.accounts.length > 0 && <Connected accounts={data.accounts} onConnectClick={() => setConnectAccountOpen(true)}/>}

      {connectAccountOpen && <ConnectAccount onClose={() => setConnectAccountOpen(false)}/>}
    </div>
  )
};
