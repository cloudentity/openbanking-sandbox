import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {FinancrooLoginPage} from '../pages/financroo/FinancrooLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from "../pages/ErrorPage";
import {FinancrooWelcomePage} from "../pages/financroo/FinancrooWelcomePage";
import {FinancrooConnectAccountPage} from "../pages/financroo/FinancrooConnectAccountPage";
import {FinancrooDashboardPage} from "../pages/financroo/FinancrooDashboardPage";

describe('Example tests', () => {
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();
  const financrooLoginPage: FinancrooLoginPage = new FinancrooLoginPage();
  const financrooWelcomePage: FinancrooWelcomePage = new FinancrooWelcomePage();
  const financrooConnectAccountPage: FinancrooConnectAccountPage = new FinancrooConnectAccountPage();
  const financrooDashboardPage: FinancrooDashboardPage = new FinancrooDashboardPage();

  const billsAccount: string = "Bills";
  const householdAccount: string = "Household";

  [
    [billsAccount, householdAccount],
    [billsAccount],
    [householdAccount],
    []
  ].forEach(accounts => {
    it(`Example test ${accounts}`, () => {
      financrooLoginPage.visit()
      financrooLoginPage.login()
      acpLoginPage.login('test', 'p@ssw0rd!')
      financrooWelcomePage.disconnect()
      financrooWelcomePage.connect()
      financrooConnectAccountPage.connectGoBank()
      financrooConnectAccountPage.allow()
      acpLoginPage.login('user', 'p@ssw0rd!')
      consentPage.checkAccounts(accounts)
      consentPage.assertPermissions([`ReadAccountsDetail`, `ReadAccountsBasic`, `ReadBalances`,
        `ReadTransactionsBasic`, `ReadTransactionsDetail`, `ReadTransactionsCredits`, `ReadTransactionsDebits`])
      consentPage.confirm()
      financrooDashboardPage.assertAccounts(accounts)
    })
  })

  it(`Example test 2`, () => {
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.cancel()
    errorPage.assertError("The user rejected the authentication")
  })

  it(`Example test 3`, () => {
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login('test', 'p@ssw0rd!')
    financrooWelcomePage.disconnect()
    financrooWelcomePage.connect()
    financrooConnectAccountPage.connectGoBank()
    financrooConnectAccountPage.allow()
    acpLoginPage.cancel()
    errorPage.assertError("The user rejected the authentication")
  })

  it(`Example test 4`, () => {
    financrooLoginPage.visit()
    financrooLoginPage.login()
    acpLoginPage.login('test', 'p@ssw0rd!')
    financrooWelcomePage.disconnect()
    financrooWelcomePage.connect()
    financrooConnectAccountPage.connectGoBank()
    financrooConnectAccountPage.allow()
    acpLoginPage.login('user', 'p@ssw0rd!')
    consentPage.cancel()
    errorPage.assertError("rejected")
  })
})
