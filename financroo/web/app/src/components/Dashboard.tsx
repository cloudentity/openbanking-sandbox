import React, { useEffect, useState } from "react";
import PageToolbar from "./common/PageToolbar";
import Connected from "./Connected";
import Welcome from "./Welcome";
import ConnectAccount from "./ConnectAccount";
import { useQuery } from "react-query";
import { api } from "../api/api";
import Progress from "./Progress";
import PageContent from "./common/PageContent";
import PageContainer from "./common/PageContainer";
import { pathOr } from "ramda";
import { useLocation } from "react-router";
import Snackbar from "@material-ui/core/Snackbar";
import IconButton from "@material-ui/core/IconButton";
import CloseIcon from "@material-ui/icons/Close";
import Alert from "@material-ui/lab/Alert";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  alert: {
    width: "100%",
    "& > div:last-of-type": {
      position: "relative",
      top: 2,
      fontWeight: 600,
    },
  },
}));

export default function Dashboard({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}) {
  const [connectAccountOpen, setConnectAccountOpen] = useState(false);
  const [isProgress, setProgress] = useState(false);
  const [snackbar, setSnackbar] = useState("");
  const classes = useStyles();

  const {
    state,
  }: { state: undefined | { bankNeedsReconnect: boolean } } = useLocation();

  useEffect(() => {
    if (state?.bankNeedsReconnect) {
      setSnackbar("Error: unauthorized. Bank needs reconnect");
    }
  }, [state]);

  const {
    isLoading: fetchBanksProgress,
    data: banksRes,
    refetch: refetchBanks,
  } = useQuery("fetchBanks", api.fetchBanks, {
    refetchOnWindowFocus: false,
    retry: false,
  });

  const banks = banksRes ? pathOr([], ["connected_banks"], banksRes) : [];

  const handleAllowAccess = ({ bankId, permissions }) => {
    setProgress(true);
    api
      .connectBank(bankId, { permissions })
      .then((res) => {
        window.location.href = res.login_url;
      })
      .catch(() => setProgress(false));
  };

  const handleDisconnectBank = (bankId) => () => {
    setProgress(true);
    api
      .disconnectBank(bankId)
      .then(refetchBanks)
      .finally(() => setProgress(false));
  };

  const handleReconnectBank = (bankId, permissions) => () => {
    setProgress(true);
    api
      .connectBank(bankId, { permissions })
      .then((res) => {
        window.location.href = res.login_url;
      })
      .catch(() => setProgress(false));
  };

  const showProgress = isProgress || fetchBanksProgress;

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="main"
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />

      {showProgress && <Progress />}

      {!showProgress && (
        <>
          {banks.length === 0 && (
            <PageContainer withBackground>
              <Welcome onConnectClick={() => setConnectAccountOpen(true)} />
            </PageContainer>
          )}
          {banks.length > 0 && (
            <PageContent>
              <Connected
                banks={banks}
                onConnectClick={() => setConnectAccountOpen(true)}
                onDisconnect={handleDisconnectBank}
                onReconnect={handleReconnectBank}
              />
            </PageContent>
          )}
        </>
      )}

      {connectAccountOpen && (
        <ConnectAccount
          connected={banks}
          onAllowAccess={handleAllowAccess}
          onClose={() => setConnectAccountOpen(false)}
        />
      )}

      <Snackbar
        anchorOrigin={{
          vertical: "top",
          horizontal: "center",
        }}
        open={!!snackbar}
        autoHideDuration={6000}
        onClose={() => setSnackbar("")}
        action={
          <>
            <IconButton
              size="small"
              aria-label="close"
              color="inherit"
              onClick={() => setSnackbar("")}
            >
              <CloseIcon fontSize="small" />
            </IconButton>
          </>
        }
      >
        <Alert severity="error" className={classes.alert}>
          {snackbar}
        </Alert>
      </Snackbar>
    </div>
  );
}
