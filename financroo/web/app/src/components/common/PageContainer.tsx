import React from "react";
import Container from "@material-ui/core/Container";
import PageContent from "./PageContent";

export default function PageContainer({
  children,
  fixed = true,
  withBackground = false,
  withSubheader = false,
  withOnlySubheader = false,
  style = {},
}) {
  return (
    <PageContent
      withBackground={withBackground}
      withSubheader={withSubheader}
      withOnlySubheader={withOnlySubheader}
      style={style}
    >
      <Container fixed={fixed}>{children}</Container>
    </PageContent>
  );
}
