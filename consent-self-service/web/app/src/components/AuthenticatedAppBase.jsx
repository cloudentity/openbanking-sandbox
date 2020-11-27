import React, {Suspense} from 'react';
import {Switch} from 'react-router';
import {Route} from "react-router-dom";
import Dashboard from './Dashboard';

export default function AuthenticatedAppBase ({authorizeURL, clientId, scopes, userinfo = {}}) {
  return (
    <Suspense>
      <Switch>
        <Route exact path={"/"} render={() =>
          <Dashboard
            authorizeURL={authorizeURL}
            userinfo={userinfo}
          />}/>
      </Switch>
    </Suspense>
  )
}
