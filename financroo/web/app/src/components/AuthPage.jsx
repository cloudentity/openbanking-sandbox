import React from 'react';
import {Redirect} from 'react-router';

import {
  getTokenFromStore,
  isTokenInStore,
  removeExpiresInFromStore,
  removeIATFromStore,
  removeTokenFromStore
} from './auth.utils';
import {generateRandomString, pkceChallengeFromVerifier} from './pkce.utils';
import Register from './Register';

export const origin = process.env.NODE_ENV === 'development'
  ? 'https://localhost:8443'
  : window.location.origin;

const calcAuthorizationUrl = async (tenantId, authorizationServerId, scopes = [], silent = false, idTokenHint = "") => {
  const authorizationUri = `${origin}/${tenantId}/${authorizationServerId}/oauth2/authorize`;

  // Create and store a random "state" value
  const state = generateRandomString();
  localStorage.setItem(`pkce_state`, state);

  // Create and store a new PKCE code_verifier (the plaintext random secret)
  const code_verifier = generateRandomString();
  localStorage.setItem(`pkce_code_verifier`, code_verifier);

  // Hash and base64-urlencode the secret to use as the challenge
  const code_challenge = await pkceChallengeFromVerifier(code_verifier);

  return authorizationUri
    + "?response_type=code"
    + "&client_id=" + encodeURIComponent("bumgdcphqf52c90ootpg")
    + "&state=" + encodeURIComponent(state)
    + "&scope=" + encodeURIComponent(scopes.join(' '))
    + "&redirect_uri=" + encodeURIComponent(window.location.origin + `/${silent ? 'silent' : 'callback'}`)
    + "&code_challenge=" + encodeURIComponent(code_challenge)
    + "&code_challenge_method=S256"
    + `${silent ? `&prompt=none&id_token_hint=${idTokenHint}` : ''}`
}

export const authorize = async (tenantId, authorizationServerId, scopes = []) => {

  // Authorization URL
  window.location.href = await calcAuthorizationUrl(tenantId, authorizationServerId, scopes)
};

const IFRAME_ID = 'silent-auth-iframe';
export const SILENT_AUTH_SUCCESS_MESSAGE = 'silentAuthSuccess';
export const SILENT_AUTH_ERROR_MESSAGE = 'silentAuthFailure';

export const silentAuthentication = async (tenantId, authorizationServerId, scopes, idTokenHint) => {
  const iframe = document.createElement("iframe");
  const src = await calcAuthorizationUrl(tenantId, authorizationServerId, scopes, true, idTokenHint);
  iframe.setAttribute("src", src)
  iframe.setAttribute("id", IFRAME_ID)
  iframe.style.display = 'none';

  const listener = e => {
    if (e.data === SILENT_AUTH_SUCCESS_MESSAGE || e.data === SILENT_AUTH_ERROR_MESSAGE) {
      const iframeToRemove = document.querySelector(`#${IFRAME_ID}`);
      iframeToRemove && document.body.removeChild(iframeToRemove);
      window.removeEventListener('message', listener);
    }
  };

  window.addEventListener("message", listener);

  document.body.appendChild(iframe);
}

export const logout = () => {
  removeTokenFromStore();
  removeExpiresInFromStore();
  removeIATFromStore();
  window.location.href = `${origin}/default/financroo/logout?redirect_to=http://localhost:3001`
};


const AuthPage = ({login, tenantId, authorizationServerId, scopes}) => {
  const handleLogin = () => {
    authorize(tenantId, authorizationServerId, scopes);
  }

  if (isTokenInStore()) {
    login({token: getTokenFromStore()});
    return (<Redirect to={'/'}/>)
  }

  return <Register onLogin={handleLogin}/>
};


export default AuthPage;
