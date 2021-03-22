import React from "react";
import Chip from "@material-ui/core/Chip";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import Radio from "@material-ui/core/Radio";
import { makeStyles } from "@material-ui/core/styles";
import clsx from "clsx";

import ContributionCard from "./ContributionCard";
import Field from "./Field";
import bankIcon from "../../assets/icon-bank.svg";
import cardIcon from "../../assets/icon-credit-card.svg";
import paypalIcon from "../../assets/icon-paypal.svg";
import walletIcon from "../../assets/icon-wallet.svg";

const useStyles = makeStyles((theme) => ({
  titleContainer: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    paddingLeft: "32px",
    paddingRight: "32px",
  },
  title: {
    fontWeight: "bold",
    fontSize: 12,
    lineHeight: "16px",
    textTransform: "uppercase",
    color: "#626576",
  },
  chip: {
    backgroundColor: theme.palette.primary.main,
    color: "white",
  },
  grid: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr 1fr",
    gridColumnGap: 14,
  },
  card: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    padding: "16px 14px",
    backgroundColor: "#FBFCFD",
    border: "1px solid #36C6AF",
    boxSizing: "border-box",
    boxShadow:
      "0px 1px 1px rgba(0, 0, 0, 0.04), 0px 3px 2px rgba(0, 0, 0, 0.04)",
    borderRadius: 4,
    "& > img": {
      width: 29,
    },
    "& > span": {
      fontSize: 12,
      lineHeight: "22px",
      color: "#626576",
    },
  },
  disabled: {
    border: "1px solid #E8EAED",
    opacity: 0.6,
    boxShadow: "none",
  },
  radioGroup: {
    display: "flex",
    "& > div": {
      fontWeight: 600,
      fontSize: 12,
      lineHeight: "24px",
      color: "#212533",
    },
  },
  accountSelect: {
    "& :first-child": {
      borderTopLeftRadius: 4,
      borderTopRightRadius: 4,
    },
    "& :last-child": {
      borderBottomLeftRadius: 4,
      borderBottomRightRadius: 4,
    },
  },
  accountSelectItem: {
    display: "flex",
    alignItems: "center",
    border: "solid 1px #ECECEC",
    borderBottom: "none",
    "&:last-of-type": {
      borderBottom: "solid 1px #ECECEC",
    },
    padding: "10px 20px 10px 11px",
    "& > img": {
      paddingLeft: 3,
      paddingRight: 12,
    },
  },
  accountSelectItemLabel: {
    "& :first-child": {
      fontWeight: 600,
      fontSize: 12,
      lineHeight: "24px",
      color: "#212533",
    },
    "& :last-child": {
      fontSize: 12,
      lineHeight: "22px",
      color: "#A0A3B5",
    },
  },
  active: {
    border: `solid 1px ${theme.palette.primary.main}`,
  },
  information: {
    display: "grid",
    gridTemplateColumns: "1fr 1fr",
    gridColumnGap: 8,
    justifyContent: "space-between",
    padding: 20,
    paddingBottom: 8,
    color: "#626576",
    fontSize: 12,
    lineHeight: "22px",
    background: "#FBFCFD",
    border: "1px solid #F4F4F4",
    borderRadius: 4,
    "& :nth-child(2n)": {
      textAlign: "right",
    },
    "& > div": {
      paddingBottom: 12,
    },
  },
}));

type Props = {
  amount: string;
  handleBack: () => void;
  handleNext: () => void;
  bank: string;
  setBank: (bank: string) => void;
  account: string;
  setAcccount: (account: string) => void;
};

export default function InvestmentsContributeMethod({
  amount,
  handleBack,
  handleNext,
  bank,
  setBank,
  account,
  setAcccount,
}: Props) {
  const classes = useStyles();

  return (
    <ContributionCard
      title={
        <div className={classes.titleContainer}>
          <div className={classes.title}>Payment total</div>
          <Chip label={`£ ${amount}`} className={classes.chip} />
        </div>
      }
      backButton={{ title: "Back", onClick: handleBack }}
      nextButton={{ title: "Next", onClick: handleNext }}
    >
      <Field label="Select payment method">
        <div className={classes.grid}>
          <div className={classes.card}>
            <img src={bankIcon} alt="bank icon" />
            <span>Bank Transfer</span>
          </div>
          <div className={clsx([classes.card, classes.disabled])}>
            <img src={cardIcon} alt="card icon" />
            <span>Credit / Debit card</span>
          </div>
          <div className={clsx([classes.card, classes.disabled])}>
            <img src={paypalIcon} alt="bank icon" />
            <span>Paypal Transfer</span>
          </div>
        </div>
      </Field>
      <Field
        label="Select Bank"
        helperText="Paying with your bank is completely safe and secure with Open Banking"
      >
        <Select
          value={bank}
          onChange={(v) => setBank(v.target.value as any)}
          style={{ width: "100%" }}
          variant="outlined"
        >
          <MenuItem value="go-bank">Go Bank</MenuItem>
          <MenuItem value="x-bank">X Bank</MenuItem>
          <MenuItem value="y-bank">Y Bank</MenuItem>
          <MenuItem value="z-bank">Z Bank</MenuItem>
        </Select>
      </Field>
      <Field>
        <div className={classes.radioGroup}>
          <div style={{ marginRight: 16 }}>
            <Radio checked color="primary" />
            Select Available Account
          </div>
          <div className={classes.disabled} style={{ border: "none" }}>
            <Radio disabled />
            Add Bank details
          </div>
        </div>
      </Field>

      <Field>
        <div className={classes.accountSelect}>
          <div className={clsx([classes.accountSelectItem, classes.active])}>
            <Radio checked color="primary" />
            <img src={walletIcon} alt="wallet icon" />
            <div className={classes.accountSelectItemLabel}>
              <div>Checking account</div>
              <div>**** ***** **** 31820</div>
            </div>
            <div style={{ flex: 1, textAlign: "right" }}>€ 99,920.20</div>
          </div>
          <div className={classes.accountSelectItem}>
            <Radio disabled />
            <img src={walletIcon} alt="wallet icon" />
            <div className={classes.accountSelectItemLabel}>
              <div>Checking account</div>
              <div>**** ***** **** 31820</div>
            </div>
            <div style={{ flex: 1, textAlign: "right" }}>€ 99,920.20</div>
          </div>
        </div>
      </Field>
      <Field label="Payee Information" style={{ marginBottom: 0 }}>
        <div className={classes.information}>
          <div>Payee Account Name</div>
          <div>Financoo investments</div>
          <div>Sort code</div>
          <div>35-64-89</div>
          <div>Account number</div>
          <div>**** ***** **** 22289</div>
          <div>Payment reference</div>
          <div>Financoo investments Ltd</div>
        </div>
      </Field>
    </ContributionCard>
  );
}
