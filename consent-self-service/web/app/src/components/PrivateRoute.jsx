import React, {Fragment, useEffect, useState} from 'react';
import {Redirect, Route} from 'react-router';
import {api} from '../api/api';
import {isTokenInStore, removeAllAuthDataFromStore} from './auth.utils';
import Progress from './Progress';

export default function PrivateRoute ({component: Component, login, ...rest}) {
  const [progress, setProgress] = useState(true);

  return (
    <Fragment>
      {progress && <Progress/>}
      {!progress && (
        <Route
          {...rest}
          render={props =>
            isTokenInStore()
              ? <Component {...props}/>
              : <Redirect
                to={{
                  pathname: '/auth',
                  state: {from: props.location}
                }}
              />}
        />)}
    </Fragment>
  )
}
