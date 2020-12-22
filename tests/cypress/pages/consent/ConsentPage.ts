export class ConsentPage {
  private readonly confirmButtonLocator = `[value="confirm"]`;

  public confirm(): void {
    cy.get(this.confirmButtonLocator).click();
  }
}
