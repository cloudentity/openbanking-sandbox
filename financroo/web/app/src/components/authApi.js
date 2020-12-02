import superagent from 'superagent';

const authApi = {
  exchangeCodeForToken: ({body}) => superagent
    .post(`https://localhost:8443/default/financroo/oauth2/token`)
    .send(body)
    .then(res => res)
};

export default authApi;
