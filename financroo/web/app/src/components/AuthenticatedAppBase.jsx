import React, {Suspense} from 'react';
import {Switch} from 'react-router';
import {Route} from "react-router-dom";
import Dashboard from './Dashboard';

export default function AuthenticatedAppBase ({authorizationServerURL, authorizationServerId, tenantId, clientId, scopes, userinfo = {}}) {

  return (
    <Suspense>
      <Switch>
        <Route exact path={"/"} render={() =>
          <Dashboard
            authorizationServerURL={authorizationServerURL}
            authorizationServerId={authorizationServerId}
            tenantId={tenantId}
            userinfo={userinfo}
          />}/>
      </Switch>
    </Suspense>
  )
}
