import {TppHomePage} from '../pages/tpp/TppHomePage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';

describe('Example tests', () => {
  const tppHomePage: TppHomePage = new TppHomePage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();

  it('Example test', () => {
    tppHomePage.visit();
    tppHomePage.next();
    tppLoginPage.login();
    acpLoginPage.login('user', 'p@ssw0rd!');
    consentPage.confirm();
    tppLoginPage.assertSuccess()
  })

  it('Failing test', () => {
    throw new Error('test fails here')
  })
})
