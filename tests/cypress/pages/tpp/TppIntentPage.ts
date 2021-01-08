export class TppIntentPage {
  private readonly loginButtonLocator: string = `[onclick="onLogin()"]`;

  public login(): void {
    cy.get(this.loginButtonLocator).click();
  }
}
