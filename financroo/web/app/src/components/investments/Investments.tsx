import React from "react";
import { useHistory } from "react-router";

import PageContainer from "../common/PageContainer";
import PageToolbar from "../common/PageToolbar";
import dashboardImg from "../../assets/investments-dashboard.svg";

export default function Investments({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const history = useHistory();
  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="main"
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
        tab="investments"
        subHeaderTitle="Investments"
        subHeaderButton={{
          title: "Contribute now",
          onClick: () => {
            history.push("/investments/contribute");
          },
        }}
      />
      <PageContainer
        withSubheader
        style={{ paddingTop: 48, paddingBottom: 48 }}
      >
        <img
          alt="financroo logo"
          src={dashboardImg}
          style={{ width: "100%" }}
        />
      </PageContainer>
    </div>
  );
}
