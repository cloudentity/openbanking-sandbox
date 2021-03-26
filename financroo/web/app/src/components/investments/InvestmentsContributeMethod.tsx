import React from "react";

import ContributionCard from "./ContributionCard";

export default function InvestmentsContributeMethod({
  amount,
  handleBack,
  handleNext,
}) {
  return (
    <ContributionCard
      title="Payment total"
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{ title: "Next", onClick: handleNext }}
    >
      {amount}
    </ContributionCard>
  );
}
