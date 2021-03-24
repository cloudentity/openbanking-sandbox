import React, { useEffect, useMemo, useState } from "react";
import IconButton from "@material-ui/core/IconButton";
import { makeStyles } from "@material-ui/core/styles";
import { ArrowLeft, Lock } from "react-feather";
import { useHistory } from "react-router";
import { useQuery } from "react-query";
import { path } from "ramda";

import PageContainer from "../common/PageContainer";
import PageToolbar, { subHeaderHeight } from "../common/PageToolbar";
import InvestmentsContributeAmount from "./InvestmentsContributeAmount";
import InvestmentsContributeMethod from "./InvestmentsContributeMethod";
import InvestmentsContributeSummary from "./InvestmentsContributeSummary";
import { api } from "../../api/api";
import Progress from "../Progress";
import { banks as banksArray } from "../banks";

export type Balance = {
  AccountId: string;
  Amount: { Amount: string; Currency: string };
  BankId: string;
  CreditDebitIndicator: string;
  CreditLine: any;
  DateTime: string;
  Type: string;
};

const useStyles = makeStyles((theme) => ({
  toolbarButton: {
    color: theme.palette.primary.main,
  },
  title: {
    display: "flex",
    alignItems: "center",
  },
  footer: {
    padding: 34,
    display: "flex",
    alignItems: "center",
    ...theme.custom.caption,
  },
  spacer: {
    flex: 1,
  },
}));

const stepsTitle = {
  0: "Contributions",
  1: "Select the payment method",
  2: "Investment summary",
};

export default function InvestmentsContribute() {
  const classes = useStyles();
  const history = useHistory();
  const [step, setStep] = useState(0);
  const [amount, setAmount] = useState("");
  const [bank, setBank] = useState("");
  const [account, setAccount] = useState("");
  const [alert, setAlert] = useState("");

  const [isProgress] = useState(false);

  const {
    isLoading: fetchBanksProgress,
    error: fetchBanksError,
    data: banksRes,
  } = useQuery("fetchBanks", api.fetchBanks, {
    refetchOnWindowFocus: false,
    retry: false,
  });

  const { isLoading: fetchBalancesProgress, error: fetchBalancesError, data: balancesRes } = useQuery(
    "fetchBalances",
    api.fetchBalances,
    {
      refetchOnWindowFocus: false,
      retry: false,
    }
  );

  useEffect(() => {
    if (step === -1) {
      history.push("/investments");
      setStep(0);
      setAmount("");
    }
  }, [step, history, amount, bank, account]);

  function handleBack() {
    setStep((step) => step - 1);
  }

  function handleNext() {
    if (step === 2) {
      // FIXME request with transaction
      console.log("BANK", amount, bank, account);
      history.push("/investments/contribute/mock-id/success");
    } else {
      setStep((step) => step + 1);
    }
  }

  const balances = useMemo(() => {
    const tmpBallances = balancesRes?.balances ?? [];
    if (tmpBallances.length) {
      setAccount(tmpBallances[0].AccountId);
    }
    return tmpBallances;
  }, [balancesRes]);

  const banks = useMemo(() => {
    const tmpBanks = banksRes?.connected_banks ?? [];
    if (tmpBanks.length) {
      setBank(tmpBanks[0]);
    }
    return tmpBanks.map(
      (b) => banksArray.find((v) => v.value === b) || { value: b, name: b }
    );
  }, [banksRes]);

  const showProgress =
    isProgress || fetchBanksProgress || fetchBalancesProgress;

  useEffect(() => {
    const bankNeedsReconnect =
      path(["response", "error", "status"], fetchBanksError) === 401 ||
      path(["response", "error", "status"], fetchBalancesError) === 401;

    if (bankNeedsReconnect) {
      history.push({ pathname: "/", state: { bankNeedsReconnect } });
    }
  }, [fetchBanksError, fetchBalancesError, history]);

  return (
    <div style={{ position: "relative" }}>
      <PageToolbar
        mode="onlySubheader"
        subHeaderTitle={
          <div className={classes.title}>
            <IconButton
              className={classes.toolbarButton}
              onClick={() => {
                setStep((step) => step - 1);
              }}
            >
              <ArrowLeft style={{ color: "#36C6AF" }} />
            </IconButton>
            <span>{stepsTitle[step]}</span>
          </div>
        }
      />
      {showProgress ? (
        <Progress top={150} />
      ) : (
        <PageContainer
          withOnlySubheader
          style={{
            paddingTop: 48,
            marginTop: subHeaderHeight,
            position: "relative",
          }}
          containerStyle={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            minHeight: "100%",
          }}
        >
          {step === 0 && (
            <InvestmentsContributeAmount
              amount={amount}
              setAmount={setAmount}
              handleBack={handleBack}
              handleNext={handleNext}
              account={account}
              setAlert={setAlert}
            />
          )}
          {step === 1 && (
            <InvestmentsContributeMethod
              amount={amount}
              handleBack={handleBack}
              handleNext={handleNext}
              bank={bank}
              setBank={setBank}
              account={account}
              setAcccount={setAccount}
              banks={banks || []}
              balances={balances || []}
              alert={alert}
              setAlert={setAlert}
            />
          )}
          {step === 2 && (
            <InvestmentsContributeSummary
              amount={amount}
              handleBack={handleBack}
              handleNext={handleNext}
              bank={bank}
              account={account}
              balances={balances}
            />
          )}
          <div className={classes.spacer} />
          <div className={classes.footer}>
            <Lock style={{ marginRight: 12 }} />
            We use multi-level ecryption measures to protect your data
          </div>
        </PageContainer>
      )}
    </div>
  );
}
