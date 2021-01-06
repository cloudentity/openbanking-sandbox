export class FinancrooWelcomePage {
  private readonly connectButtonLocator: string = '.connect-button';

  public disconnect(): void {
    cy.intercept('POST', 'https://localhost:8443/default/financroo/oauth2/token').as('getToken')
    cy.wait('@getToken').then(interception => {
      cy.request({
        method: `DELETE`,
        url: `https://localhost:8091/api/disconnect/gobank`,
        auth: {bearer: interception.response.body.access_token}
      })
    })
    cy.reload()
  }

  public connect(): void {
    cy.get(this.connectButtonLocator).click()
  }


}
