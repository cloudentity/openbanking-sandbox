import {base} from './api-base';

export const api = {
  fetchAccounts: () => base.get({url: `/accounts`}),
};
