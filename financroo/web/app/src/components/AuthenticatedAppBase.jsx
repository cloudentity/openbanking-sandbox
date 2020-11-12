import React, {Suspense} from 'react';
import {Switch} from 'react-router';
import {Route} from "react-router-dom";
import Dashboard from './Dashboard';
import {useSilentAuthentication} from './useSilentAuthentication';

export default function AuthenticatedAppBase ({authorizationServerId, tenantId, scopes, userinfo = {}}) {
  useSilentAuthentication(authorizationServerId, tenantId, scopes);

  return (
    <Suspense>
      <Switch>
        <Route exact path={"/"} render={() => <Dashboard userinfo={userinfo}/>}/>
      </Switch>
    </Suspense>
  )
}
