import {Urls} from '../Urls'

export class FinancrooLoginPage {
  private readonly loginButtonLocator: string = `.login-button`;

  public visit(): void {
    cy.visit(Urls.financrooUrl);
  }

  public login(): void {
    cy.get(this.loginButtonLocator).click();
  }
}
