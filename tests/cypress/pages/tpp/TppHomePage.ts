import {Urls} from '../Urls'

export class TppHomePage {
  private readonly nextButtonLocator = `[type="submit"]`;

  public visit(): void {
    cy.visit(Urls.tppUrl);
  }

  public next(): void {
    cy.get(this.nextButtonLocator).click();
  }
}
