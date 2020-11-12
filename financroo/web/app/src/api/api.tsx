import {base, baseWithCustomBaseUrl, toJson} from './api-base';

export const api = {
  fetchAccounts: () => base.get({url: `/accounts`}),
  updateAccounts: body => base.put({url: `/accounts`, body}),
  userinfo: (authorizationServerURL, tenantId, authorizationServerId) => baseWithCustomBaseUrl('/', authorizationServerURL).get({url: `/${tenantId}/${authorizationServerId}/userinfo`}),
};
