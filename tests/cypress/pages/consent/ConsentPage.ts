export class ConsentPage {
  private readonly confirmButtonLocator: string = `[value="confirm"]`;
  private readonly cancelButtonLocator: string = `[value="deny"]`;
  private readonly permissionRowLocator: string = `.permission .account-header-title`;

  public assertPermissions(permission: string[]): void {
    cy.get(this.permissionRowLocator).invoke('text').should('equal', permission.join(``))
  }

  public confirm(): void {
    cy.get(this.confirmButtonLocator).click();
  }

  public cancel(): void {
    cy.get(this.cancelButtonLocator).click();
  }
}
