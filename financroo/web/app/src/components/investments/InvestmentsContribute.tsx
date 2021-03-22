import React, { useEffect, useState } from "react";
import IconButton from "@material-ui/core/IconButton";
import { makeStyles } from "@material-ui/core/styles";
import { ArrowLeft, Lock } from "react-feather";
import { useHistory } from "react-router";

import PageContainer from "../common/PageContainer";
import PageToolbar, { subHeaderHeight } from "../common/PageToolbar";
import InvestmentsContributeAmount from "./InvestmentsContributeAmount";
import InvestmentsContributeMethod from "./InvestmentsContributeMethod";

const useStyles = makeStyles((theme) => ({
  toolbarButton: {
    color: theme.palette.primary.main,
  },
  title: {
    display: "flex",
    alignItems: "center",
  },
  footer: {
    position: "absolute",
    bottom: 0,
    left: "50%",
    transform: "translateX(-50%)",
    display: "flex",
    alignItems: "center",
    fontSize: 12,
    lineHeight: "22px",
    color: "#626576",
    padding: 34,
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
  const [amount, setAmount] = useState("1");
  const [paymentMethod] = useState("bank-transfer");
  const [bank, setBank] = useState("go-bank");
  const [account, setAccount] = useState("available-accocunt");

  useEffect(() => {
    if (step === -1) {
      history.push("/investments");
      setStep(0);
      setAmount("");
    }
  }, [step, history]);

  function handleBack() {
    setStep((step) => step - 1);
  }

  function handleNext() {
    setStep((step) => step + 1);
  }

  console.log(amount, paymentMethod, bank, account);

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
      <PageContainer
        withOnlySubheader
        style={{
          paddingTop: 48,
          marginTop: subHeaderHeight,
          position: "relative",
        }}
      >
        {step === 0 && (
          <InvestmentsContributeAmount
            amount={amount}
            setAmount={setAmount}
            handleBack={handleBack}
            handleNext={handleNext}
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
          />
        )}
        {/* FIXME */}
        {/* <div className={classes.footer}>
          <Lock style={{ marginRight: 12 }} />
          We use multi-level ecryption measures to protect your data
        </div> */}
      </PageContainer>
    </div>
  );
}
