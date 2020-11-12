import {base, baseWithCustomBaseUrl, toJson} from './api-base';
import {origin} from "../components/AuthPage";

export const api = {
  fetchAccounts: () => base.get({url: `/accounts`}),
  updateAccounts: body => base.put({url: `/accounts`, body}),
  userinfo: () => baseWithCustomBaseUrl('/', origin).get({url: `/default/financroo/userinfo`}),
};
