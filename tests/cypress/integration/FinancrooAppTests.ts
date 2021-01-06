import {TppAuthenticatedPage} from '../pages/tpp/TppAuthenticatedPage';
import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from "../pages/ErrorPage";
import {FinancrooWelcomePage} from "../pages/financroo/FinancrooWelcomePage";
import {FinancrooConnectAccountPage} from "../pages/financroo/FinancrooConnectAccountPage";

describe('Example tests', () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();

  it(`Example test`, () => {
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login('test', 'p@ssw0rd!')




    financrooWelcomePage.connect()
    financrooConnectAccountPage.connectGoBank()
    financrooConnectAccountPage.allow()
    acpLoginPage.login('user', 'p@ssw0rd!')
    consentPage.assertPermissions([`ReadAccountsDetail`, `ReadAccountsBasic`, `ReadBalances`,
      `ReadTransactionsBasic`, `ReadTransactionsDetail`, `ReadTransactionsCredits`, `ReadTransactionsDebits`])
    consentPage.confirm()
  })
})
