import {Urls} from "../Urls";

export class FinancrooWelcomePage {
  private readonly connectButtonLocator: string = `.connect-button`;

  public disconnect(): void {
    cy.intercept(`POST`, `${Urls.acpUrl}/default/financroo/oauth2/token`).as(`getToken`)
    cy.wait(`@getToken`).then(interception => {
      cy.request({
        method: `DELETE`,
        url: `${Urls.financrooUrl}/api/disconnect/gobank`,
        auth: {bearer: interception.response.body.access_token}
      })
    })
    cy.reload()
  }

  public connect(): void {
    cy.get(this.connectButtonLocator).click()
  }


}
