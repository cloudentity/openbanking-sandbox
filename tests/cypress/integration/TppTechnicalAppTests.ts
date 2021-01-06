import {TppAuthenticatedPage} from '../pages/tpp/TppAuthenticatedPage';
import {TppIntentPage} from '../pages/tpp/TppIntentPage';
import {TppLoginPage} from '../pages/tpp/TppLoginPage';
import {AcpLoginPage} from '../pages/acp/AcpLoginPage';
import {ConsentPage} from '../pages/consent/ConsentPage';
import {ErrorPage} from "../pages/ErrorPage";

describe('Example tests', () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const consentPage: ConsentPage = new ConsentPage();
  const errorPage: ErrorPage = new ErrorPage();

  const basicPermission: string = 'ReadAccountsBasic';
  const detailPermission: string = 'ReadAccountsDetail';

  [
    [basicPermission, detailPermission],
    [basicPermission],
    [detailPermission],
    []
  ].forEach(permissions => {
    it(`Example test with permissions: ${permissions}`, () => {
      tppLoginPage.visit();
      tppLoginPage.checkBasicPermission(permissions.includes(basicPermission))
      tppLoginPage.checkDetailPermission(permissions.includes(detailPermission))
      tppLoginPage.next();
      if (!permissions.includes(basicPermission) && !permissions.includes(detailPermission)) {
        errorPage.assertError("Invalid create account access consent request")
      } else {
        tppIntentPage.login();
        acpLoginPage.login('user', 'p@ssw0rd!');
        consentPage.assertPermissions(permissions)
        consentPage.confirm();
        if (!permissions.includes(basicPermission) && permissions.includes(detailPermission)) {
          errorPage.assertError("failed to call bank get accounts")
        } else {
          tppAuthenticatedPage.assertSuccess()
        }
      }
    })
  })

  it(`Cancel 1`, () => {
    tppLoginPage.visit();
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.cancel();
    errorPage.assertError("The user rejected the authentication")
  })

  it(`Cancel 2`, () => {
    tppLoginPage.visit();
    tppLoginPage.next();
    tppIntentPage.login();
    acpLoginPage.login('user', 'p@ssw0rd!');
    consentPage.cancel()
    errorPage.assertError("rejected")
  })
})
