import React from "react";
import {Card, Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import AnalyticsTable from "./AnalyticsTable";
import AnalyticsBarChart from "./AnalyticsBarChart";
import AnalyticsPieChart from "./AnalyticsPieChart";
import mainClasses from "./main.module.css";

const useStyles = makeStyles((theme: Theme) => ({}));

export default ({}) => {
  const classes = useStyles();

  return (
    <>
      <div style={{display: 'flex', justifyContent: 'space-between'}}>
        <Typography className={mainClasses.sectionTitle}>Analytics</Typography>
        <div style={{textAlign: 'right'}}>
          <Typography>Credit account</Typography>
          <Typography>**** ***** **** 4099</Typography>
        </div>
      </div>
      <Card style={{padding: 16, display: 'flex', alignItems: 'center'}}>
        <div style={{flex: 3}}>
          <AnalyticsBarChart/>
        </div>
        <div style={{flex: 1}}>
          <AnalyticsPieChart/>
        </div>
      </Card>
      <AnalyticsTable style={{marginTop: 24, height: 'calc(100% - 52px - 332px - 24px'}}/>
    </>
  )
};
