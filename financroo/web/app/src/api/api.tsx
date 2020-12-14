import {base, baseWithCustomBaseUrl} from './api-base';

export const api = {
  fetchAccounts: () => base.get({url: `/accounts`}),
  connectBank: (bankId, body) => base.post({url: `/connect/${bankId}`, body, query: {}}),
  fetchTransactions: () => base.get({url: `/transactions`}),
  fetchBalances: () => base.get({url: `/balances`}),
  userinfo: (authorizationServerURL, tenantId, authorizationServerId) => baseWithCustomBaseUrl('/', authorizationServerURL).get({url: `/${tenantId}/${authorizationServerId}/userinfo`}),
};
