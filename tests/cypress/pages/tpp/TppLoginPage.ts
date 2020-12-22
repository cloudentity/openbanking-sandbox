export class TppLoginPage {
  private readonly loginButtonLocator = '[onclick="onLogin()"]';
  private readonly statusTextLocator = '.aut-demo-info-text';

  public login(): void {
    cy.get(this.loginButtonLocator).click();
  }

  public assertSuccess(): void {
    cy.get(this.statusTextLocator).should('contain.text', 'Authenticated')
  }
}
